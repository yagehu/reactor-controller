package errs

import (
	"bytes"
	"fmt"
)

type Error struct {
	Op   Op
	Kind Kind
	Err  error
}

type Op string

func E(args ...interface{}) error {
	if len(args) == 0 {
		panic("errs.E called with 0 argument")
	}

	var e Error

	for _, arg := range args {
		switch arg := arg.(type) {
		case Op:
			e.Op = arg
		case Kind:
			e.Kind = arg
		case *Error:
			// Make a copy
			cpy := *arg
			e.Err = &cpy
		case error:
			e.Err = arg
		default:
			return fmt.Errorf("unknown type %T, value %v in error call", arg, arg)
		}
	}

	prev, ok := e.Err.(*Error)
	if !ok {
		return &e
	}

	if prev.Kind == e.Kind {
		prev.Kind = KindOther
	}

	// If this error has Kind unset or Other, pull up the inner one.
	if e.Kind == KindOther {
		e.Kind = prev.Kind
		prev.Kind = KindOther
	}

	return &e
}

func (e Error) Error() string {
	b := new(bytes.Buffer)

	if e.Op != "" {
		pad(b, ": ")
		b.WriteString(string(e.Op))
	}

	if e.Kind != 0 {
		pad(b, ": ")
		b.WriteString(e.Kind.String())
	}

	if e.Err != nil {
		// Indent on new line if we are cascading non-empty errs.
		if prevErr, ok := e.Err.(*Error); ok {
			if !prevErr.isZero() {
				pad(b, ":\n\t")
				b.WriteString(e.Err.Error())
			}
		} else {
			pad(b, ": ")
			b.WriteString(e.Err.Error())
		}
	}

	if b.Len() == 0 {
		return "no error"
	}

	return b.String()
}

func (e *Error) isZero() bool {
	return e.Op == "" && e.Kind == 0 && e.Err == nil
}

// pad appends str to the buffer if the buffer already has some data.
func pad(b *bytes.Buffer, str string) {
	if b.Len() == 0 {
		return
	}

	b.WriteString(str)
}
