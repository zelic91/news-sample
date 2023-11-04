package object

import (
	"context"
	"news/db/postgres"
	"news/db/postgres/dbgen"

	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db    *sqlx.DB
	query *dbgen.Queries
}

func NewRepo(db *sqlx.DB) Repo {
	return Repo{
		db:    db,
		query: dbgen.New(db),
	}
}

func (r Repo) FindAll(
	ctx context.Context,
) ([]*dbgen.Object, error) {
	return r.query.FindAllObjects(ctx)
}

func (r Repo) FindByID(
	ctx context.Context,
	id int64,
) (*dbgen.Object, error) {
	return r.query.FindObjectById(ctx, id)
}

func (r Repo) Create(
	ctx context.Context,
	categoryId int64,
	content *string,
	title *string,
) (*dbgen.Object, error) {
	params := dbgen.CreateObjectParams{
		CategoryID: categoryId,
		Content:    postgres.ToNullString(content),
		Title:      postgres.ToNullString(title),
	}

	return r.query.CreateObject(ctx, params)
}

func (r Repo) Update(
	ctx context.Context,
	id int64,
	categoryId int64,
	content *string,
	title *string,
) (*dbgen.Object, error) {
	params := dbgen.UpdateObjectParams{
		ID:         id,
		CategoryID: categoryId,
		Content:    postgres.ToNullString(content),
		Title:      postgres.ToNullString(title),
	}

	return r.query.UpdateObject(ctx, params)
}

func (r Repo) Delete(
	ctx context.Context,
	id int64,
) error {
	return r.query.DeleteObject(ctx, id)
}
