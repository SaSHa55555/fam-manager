package cmd

import (
	"github.com/SaSHa55555/fam-manager/internal/api"
	"github.com/SaSHa55555/fam-manager/internal/api/repository"
	"github.com/SaSHa55555/fam-manager/internal/api/service"
	transport "github.com/SaSHa55555/fam-manager/internal/api/transport/http"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"github.com/uptrace/opentelemetry-go-extra/otelsqlx"
)

type Api struct {
	ApiTransport api.IApiTransport
}

func ConfigureApi() (*Api, error) {
	db, err := otelsqlx.Connect(
		"pgx",
		"postgres://postgres:password@localhost:5432/manager",
		otelsql.WithDBName("manager"),
	)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	r := repository.NewRepository(db)
	s := service.NewService(r)
	handler := transport.NewHandler(s)

	return &Api{
		ApiTransport: handler,
	}, nil
}
