package device

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
) ([]*dbgen.Device, error) {
	return r.query.FindAllDevices(ctx)
}

func (r Repo) FindByID(
	ctx context.Context,
	id int64,
) (*dbgen.Device, error) {
	return r.query.FindDeviceById(ctx, id)
}

func (r Repo) Create(
	ctx context.Context,
	deviceToken *string,
	platform *string,
	status *string,
) (*dbgen.Device, error) {
	params := dbgen.CreateDeviceParams{
		DeviceToken: postgres.ToNullString(deviceToken),
		Platform:    postgres.ToNullString(platform),
		Status:      postgres.ToNullString(status),
	}

	return r.query.CreateDevice(ctx, params)
}

func (r Repo) Update(
	ctx context.Context,
	id int64,
	deviceToken *string,
	platform *string,
	status *string,
) (*dbgen.Device, error) {
	params := dbgen.UpdateDeviceParams{
		ID:          id,
		DeviceToken: postgres.ToNullString(deviceToken),
		Platform:    postgres.ToNullString(platform),
		Status:      postgres.ToNullString(status),
	}

	return r.query.UpdateDevice(ctx, params)
}

func (r Repo) Delete(
	ctx context.Context,
	id int64,
) error {
	return r.query.DeleteDevice(ctx, id)
}
