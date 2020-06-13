package errs

type Kind int

const (
	KindOther Kind = iota
	KindReactorNotFound
)

func (k Kind) String() string {
	switch k {
	case KindReactorNotFound:
		return "reactor not found"
	default:
		return "unknown error kind"
	}
}
