package user

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
) ([]*dbgen.User, error) {
	return r.query.FindAllUsers(ctx)
}

func (r Repo) FindByID(
	ctx context.Context,
	id int64,
) (*dbgen.User, error) {
	return r.query.FindUserById(ctx, id)
}

func (r Repo) Create(
	ctx context.Context,
	email *string,
	firstName *string,
	lastName *string,
	passwordHashed *string,
	passwordSalt *string,
	status *string,
	username string,
) (*dbgen.User, error) {
	params := dbgen.CreateUserParams{
		Email:          postgres.ToNullString(email),
		FirstName:      postgres.ToNullString(firstName),
		LastName:       postgres.ToNullString(lastName),
		PasswordHashed: postgres.ToNullString(passwordHashed),
		PasswordSalt:   postgres.ToNullString(passwordSalt),
		Status:         postgres.ToNullString(status),
		Username:       username,
	}

	return r.query.CreateUser(ctx, params)
}

func (r Repo) Update(
	ctx context.Context,
	id int64,
	email *string,
	firstName *string,
	lastName *string,
	passwordHashed *string,
	passwordSalt *string,
	status *string,
	username string,
) (*dbgen.User, error) {
	params := dbgen.UpdateUserParams{
		ID:             id,
		Email:          postgres.ToNullString(email),
		FirstName:      postgres.ToNullString(firstName),
		LastName:       postgres.ToNullString(lastName),
		PasswordHashed: postgres.ToNullString(passwordHashed),
		PasswordSalt:   postgres.ToNullString(passwordSalt),
		Status:         postgres.ToNullString(status),
		Username:       username,
	}

	return r.query.UpdateUser(ctx, params)
}

func (r Repo) Delete(
	ctx context.Context,
	id int64,
) error {
	return r.query.DeleteUser(ctx, id)
}

func (r Repo) FindByUsername(
	ctx context.Context,
	username string,
) (*dbgen.User, error) {
	return r.query.FindUserByUsername(ctx, postgres.ToNullString(&username))
}

func (r Repo) FindByEmail(
	ctx context.Context,
	email string,
) (*dbgen.User, error) {
	return r.query.FindUserByEmail(ctx, postgres.ToNullString(&email))
}
