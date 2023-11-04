package postgres

import (
	"database/sql"
	"news/config"

	"github.com/jmoiron/sqlx"
)

func Init(config *config.Config) *sqlx.DB {
	return sqlx.MustOpen("postgres", config.PostgresURL)
}

func ToNullInt64(input *int64) sql.NullInt64 {
	if input == nil {
		return sql.NullInt64{Valid: false}
	}
	return sql.NullInt64{Int64: *input, Valid: true}
}

func ToNullString(input *string) sql.NullString {
	if input == nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: *input, Valid: true}
}

func ToNullableString(input string) *string {
	return &input
}
