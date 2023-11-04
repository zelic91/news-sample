package device

import (
	"context"

	"news/api/gen"
	"news/common"
)

type CreateDeviceParams struct {
	DeviceToken *string
	Platform    *string
	Status      *string
}

type UpdateDeviceParams struct {
	ID          int64
	DeviceToken *string
	Platform    *string
	Status      *string
}

type service struct {
	repo *Repo
}

func NewService(
	repo *Repo,
) *service {
	return &service{
		repo: repo,
	}
}

/*
	For reference

	func (s *service) CreateDevice(ctx context.Context, params *gen.CreateDevice) (*gen.CreateDeviceResponse, error) {
		res, err := s.repo.Create(
			ctx,
			params.DeviceToken,
			params.Platform,
			params.Status,
		)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		return &gen.CreateDeviceResponseBody{
			// TODO: Implement this
		}, nil

		return nil, nil
	}

	func (s *service) UpdateDevice(ctx context.Context, id int64, params *gen.UpdateDevice) (*gen.UpdateDeviceResponse, error) {
		res, err := s.repo.Update(
			ctx,
			id,
			params.DeviceToken,
			params.Platform,
			params.Status,
		)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		return &gen.UpdateDeviceResponseBody{
			// TODO: Implement this
		}, nil

		return nil, nil
	}
*/
func (s *service) CreateDevice(ctx context.Context, user *common.AuthUser, body *gen.CreateDevice) (*gen.DefaultResponse, error) {
	return nil, nil
}
