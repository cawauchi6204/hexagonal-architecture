package handlers

import (
	"net/http"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/presentation/dto"
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/usecase"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(userUseCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

// @Summary ユーザー登録
// @Description 新規ユーザーを登録する
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.RegisterUserRequest true "ユーザー登録情報"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} ErrorResponse
// @Router /users/register [post]
func (h *UserHandler) Register(c echo.Context) error {
	var req dto.RegisterUserRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := h.userUseCase.Register(c.Request().Context(), req.Username, req.Email, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, dto.NewUserResponse(user))
}

// @Summary ログイン
// @Description ユーザーログインを行う
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "ログイン情報"
// @Success 200 {object} dto.LoginResponse
// @Failure 401 {object} ErrorResponse
// @Router /users/login [post]
func (h *UserHandler) Login(c echo.Context) error {
	var req dto.LoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := h.userUseCase.Login(c.Request().Context(), req.Email, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	// TODO: JWTトークンの生成
	token := "dummy-token"

	return c.JSON(http.StatusOK, dto.LoginResponse{
		Token: token,
		User:  *dto.NewUserResponse(user),
	})
}

// @Summary プロフィール更新
// @Description ユーザープロフィールを更新する
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "ユーザーID"
// @Param request body dto.UpdateProfileRequest true "更新情報"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} ErrorResponse
// @Security BearerAuth
// @Router /users/{id} [put]
func (h *UserHandler) UpdateProfile(c echo.Context) error {
	userID := c.Param("id")
	var req dto.UpdateProfileRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := h.userUseCase.UpdateProfile(c.Request().Context(), userID, req.Username, req.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, dto.NewUserResponse(user))
}

// @Summary フォロー
// @Description ユーザーをフォローする
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "フォローするユーザーID"
// @Success 200 {object} dto.FollowResponse
// @Failure 400 {object} ErrorResponse
// @Security BearerAuth
// @Router /users/{id}/follow [post]
func (h *UserHandler) Follow(c echo.Context) error {
	followerID := c.Get("user_id").(string) // JWTから取得
	followedID := c.Param("id")

	if err := h.userUseCase.Follow(c.Request().Context(), followerID, followedID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, dto.FollowResponse{Success: true})
}

// @Summary フォロー解除
// @Description ユーザーのフォローを解除する
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "フォロー解除するユーザーID"
// @Success 200 {object} dto.FollowResponse
// @Failure 400 {object} ErrorResponse
// @Security BearerAuth
// @Router /users/{id}/unfollow [post]
func (h *UserHandler) Unfollow(c echo.Context) error {
	followerID := c.Get("user_id").(string) // JWTから取得
	followedID := c.Param("id")

	if err := h.userUseCase.Unfollow(c.Request().Context(), followerID, followedID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, dto.FollowResponse{Success: true})
}

// @Summary プロフィール取得
// @Description ユーザーのプロフィールを取得する
// @Tags users
// @Produce json
// @Param id path string true "ユーザーID"
// @Success 200 {object} dto.UserResponse
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [get]
func (h *UserHandler) GetProfile(c echo.Context) error {
	userID := c.Param("id")
	user, err := h.userUseCase.GetProfile(c.Request().Context(), userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, dto.NewUserResponse(user))
}

// @Summary ユーザー一覧取得
// @Description 全ユーザーの一覧を取得する
// @Tags users
// @Produce json
// @Success 200 {object} dto.UserListResponse
// @Router /users [get]
func (h *UserHandler) List(c echo.Context) error {
	users, err := h.userUseCase.GetAllUsers(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, dto.NewUserListResponse(users))
}
