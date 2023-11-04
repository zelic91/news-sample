package category

import (
	"context"

	"news/api/gen"
	"news/common"
)

type CreateCategoryParams struct {
	Description *string
	Title       *string
}

type UpdateCategoryParams struct {
	ID          int64
	Description *string
	Title       *string
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

	func (s *service) CreateCategory(ctx context.Context, params *gen.CreateCategory) (*gen.CreateCategoryResponse, error) {
		res, err := s.repo.Create(
			ctx,
			params.Description,
			params.Title,
		)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		return &gen.CreateCategoryResponseBody{
			// TODO: Implement this
		}, nil

		return nil, nil
	}

	func (s *service) UpdateCategory(ctx context.Context, id int64, params *gen.UpdateCategory) (*gen.UpdateCategoryResponse, error) {
		res, err := s.repo.Update(
			ctx,
			id,
			params.Description,
			params.Title,
		)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		return &gen.UpdateCategoryResponseBody{
			// TODO: Implement this
		}, nil

		return nil, nil
	}
*/
func (s *service) CreateCategory(ctx context.Context, user *common.AuthUser, body *gen.CreateCategory) (*gen.DefaultResponse, error) {
	return nil, nil
}
func (s *service) GetCategories(ctx context.Context, user *common.AuthUser) (*gen.GetCategoriesBody, error) {
	return nil, nil
}
