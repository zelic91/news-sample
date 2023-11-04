package category

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
) ([]*dbgen.Category, error) {
	return r.query.FindAllCategories(ctx)
}

func (r Repo) FindByID(
	ctx context.Context,
	id int64,
) (*dbgen.Category, error) {
	return r.query.FindCategoryById(ctx, id)
}

func (r Repo) Create(
	ctx context.Context,
	description *string,
	title *string,
) (*dbgen.Category, error) {
	params := dbgen.CreateCategoryParams{
		Description: postgres.ToNullString(description),
		Title:       postgres.ToNullString(title),
	}

	return r.query.CreateCategory(ctx, params)
}

func (r Repo) Update(
	ctx context.Context,
	id int64,
	description *string,
	title *string,
) (*dbgen.Category, error) {
	params := dbgen.UpdateCategoryParams{
		ID:          id,
		Description: postgres.ToNullString(description),
		Title:       postgres.ToNullString(title),
	}

	return r.query.UpdateCategory(ctx, params)
}

func (r Repo) Delete(
	ctx context.Context,
	id int64,
) error {
	return r.query.DeleteCategory(ctx, id)
}
