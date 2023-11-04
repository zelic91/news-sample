package object

import (
	"context"

	"news/api/gen"
	"news/common"
)

type CreateObjectParams struct {
	CategoryID int64
	Content    *string
	Title      *string
}

type UpdateObjectParams struct {
	ID         int64
	CategoryID int64
	Content    *string
	Title      *string
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

	func (s *service) CreateObject(ctx context.Context, params *gen.CreateObject) (*gen.CreateObjectResponse, error) {
		res, err := s.repo.Create(
			ctx,
			params.CategoryId,
			params.Content,
			params.Title,
		)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		return &gen.CreateObjectResponseBody{
			// TODO: Implement this
		}, nil

		return nil, nil
	}

	func (s *service) UpdateObject(ctx context.Context, id int64, params *gen.UpdateObject) (*gen.UpdateObjectResponse, error) {
		res, err := s.repo.Update(
			ctx,
			id,
			params.CategoryId,
			params.Content,
			params.Title,
		)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		return &gen.UpdateObjectResponseBody{
			// TODO: Implement this
		}, nil

		return nil, nil
	}
*/
func (s *service) CreateObject(ctx context.Context, user *common.AuthUser, body *gen.CreateObject) (*gen.DefaultResponse, error) {
	return nil, nil
}
func (s *service) GetObjects(ctx context.Context, user *common.AuthUser) (*gen.GetObjectsBody, error) {
	return nil, nil
}
