package dbfx

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"go.uber.org/fx"

	"github.com/yagehu/reactor-controller/config"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In

	Config config.Config
}

type Result struct {
	fx.Out

	Db *sql.DB
}

func New(p Params) (Result, error) {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		p.Config.Postgres.Host,
		p.Config.Postgres.Port,
		p.Config.Postgres.User,
		p.Config.Postgres.Password,
		p.Config.Postgres.Database,
	))
	if err != nil {
		return Result{}, err
	}

	if err := db.Ping(); err != nil {
		return Result{}, err
	}

	return Result{
		Db: db,
	}, nil
}
