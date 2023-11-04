package api

import (
	"context"
	"net/http"
	"news/api/gen"
	"news/common"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type AuthService interface {
	SignUp(ctx context.Context, body *gen.SignUpRequest) (*gen.AuthResponse, error)
	SignIn(ctx context.Context, body *gen.SignInRequest) (*gen.AuthResponse, error)
}
type CategoryService interface {
	CreateCategory(ctx context.Context, user *common.AuthUser, body *gen.CreateCategory) (*gen.DefaultResponse, error)
	GetCategories(ctx context.Context, user *common.AuthUser) (*gen.GetCategoriesBody, error)
}
type DeviceService interface {
	CreateDevice(ctx context.Context, user *common.AuthUser, body *gen.CreateDevice) (*gen.DefaultResponse, error)
}
type ObjectService interface {
	CreateObject(ctx context.Context, user *common.AuthUser, body *gen.CreateObject) (*gen.DefaultResponse, error)
	GetObjects(ctx context.Context, user *common.AuthUser) (*gen.GetObjectsBody, error)
}

type ApiService struct {
	AuthService
	CategoryService
	DeviceService
	ObjectService
}

type ServerImpl struct {
	service ApiService
}

func NewServerImpl(
	authService AuthService,
	categoryService CategoryService,
	deviceService DeviceService,
	objectService ObjectService,
) ServerImpl {
	return ServerImpl{
		service: ApiService{
			AuthService:     authService,
			CategoryService: categoryService,
			DeviceService:   deviceService,
			ObjectService:   objectService,
		},
	}
}
func (s *ServerImpl) SignUp(ctx echo.Context) error {
	var body gen.SignUpRequest
	err := ctx.Bind(&body)

	if err != nil {
		log.Error(err)
		return sendError(ctx, http.StatusBadRequest, 10000, "Invalid request body")
	}

	resp, err := s.service.SignUp(ctx.Request().Context(), &body)

	if err != nil {
		log.Error(err)
		return sendError(ctx, http.StatusInternalServerError, 10001, "Failed to SignUp")
	}

	return ctx.JSON(201, resp)
}
func (s *ServerImpl) SignIn(ctx echo.Context) error {
	var body gen.SignInRequest
	err := ctx.Bind(&body)

	if err != nil {
		log.Error(err)
		return sendError(ctx, http.StatusBadRequest, 10000, "Invalid request body")
	}

	resp, err := s.service.SignIn(ctx.Request().Context(), &body)

	if err != nil {
		log.Error(err)
		return sendError(ctx, http.StatusInternalServerError, 10001, "Failed to SignIn")
	}

	return ctx.JSON(200, resp)
}
func (s *ServerImpl) CreateCategory(ctx echo.Context) error {
	user := GetAuthUserFromContext(ctx)
	var body gen.CreateCategory
	err := ctx.Bind(&body)

	if err != nil {
		log.Error(err)
		return sendError(ctx, http.StatusBadRequest, 10000, "Invalid request body")
	}

	resp, err := s.service.CreateCategory(ctx.Request().Context(), user, &body)

	if err != nil {
		log.Error(err)
		return sendError(ctx, http.StatusInternalServerError, 10001, "Failed to CreateCategory")
	}

	return ctx.JSON(201, resp)
}
func (s *ServerImpl) GetCategories(ctx echo.Context) error {
	user := GetAuthUserFromContext(ctx)

	resp, err := s.service.GetCategories(ctx.Request().Context(), user)

	if err != nil {
		log.Error(err)
		return sendError(ctx, http.StatusInternalServerError, 10001, "Failed to GetCategories")
	}

	return ctx.JSON(200, resp)
}
func (s *ServerImpl) CreateDevice(ctx echo.Context) error {
	user := GetAuthUserFromContext(ctx)
	var body gen.CreateDevice
	err := ctx.Bind(&body)

	if err != nil {
		log.Error(err)
		return sendError(ctx, http.StatusBadRequest, 10000, "Invalid request body")
	}

	resp, err := s.service.CreateDevice(ctx.Request().Context(), user, &body)

	if err != nil {
		log.Error(err)
		return sendError(ctx, http.StatusInternalServerError, 10001, "Failed to CreateDevice")
	}

	return ctx.JSON(201, resp)
}
func (s *ServerImpl) CreateObject(ctx echo.Context) error {
	user := GetAuthUserFromContext(ctx)
	var body gen.CreateObject
	err := ctx.Bind(&body)

	if err != nil {
		log.Error(err)
		return sendError(ctx, http.StatusBadRequest, 10000, "Invalid request body")
	}

	resp, err := s.service.CreateObject(ctx.Request().Context(), user, &body)

	if err != nil {
		log.Error(err)
		return sendError(ctx, http.StatusInternalServerError, 10001, "Failed to CreateObject")
	}

	return ctx.JSON(201, resp)
}
func (s *ServerImpl) GetObjects(ctx echo.Context) error {
	user := GetAuthUserFromContext(ctx)

	resp, err := s.service.GetObjects(ctx.Request().Context(), user)

	if err != nil {
		log.Error(err)
		return sendError(ctx, http.StatusInternalServerError, 10001, "Failed to GetObjects")
	}

	return ctx.JSON(200, resp)
}

func sendError(ctx echo.Context, statusCode int, code int, message string) error {
	resp := gen.Error{
		Code:    code,
		Message: message,
	}

	err := ctx.JSON(statusCode, resp)

	return err
}
