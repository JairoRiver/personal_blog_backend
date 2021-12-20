package server

import (
	"database/sql"
	"github.com/JairoRiver/personal_blog_backend/internal/util"
	db "github.com/JairoRiver/personal_blog_backend/pkg/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"net/http"
	"time"
)

//createUser handler
type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	RoleID   string `json:"role_id" binding:"required,uuid"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type createUserResponse struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	roleID, _ := uuid.Parse(req.RoleID)

	arg := db.CreateUserParams{
		Username: req.Username,
		RoleID:   roleID,
		Email:    req.Email,
		Password: hashedPassword,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := createUserResponse{
		Username:  user.Username,
		Email:     user.Email,
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
	}
	ctx.JSON(http.StatusCreated, res)
}

//get user handler
type getUserRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type getUserResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	RoleID    uuid.UUID `json:"role_id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, _ := uuid.Parse(req.ID)

	user, err := server.store.GetUser(ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := getUserResponse{
		ID:        user.ID,
		Username:  user.Username,
		RoleID:    user.RoleID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	ctx.JSON(http.StatusOK, res)
}

//list users handler
type listUserRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=100"`
}

func (server *Server) listUsers(ctx *gin.Context) {
	var req listUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	users, err := server.store.ListUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
}

//update user handler
type updateUserRequest struct {
	ID       string `json:"id" uri:"id" binding:"required,uuid"`
	Username string `json:"username" binding:"omitempty"`
	Email    string `json:"email" binding:"omitempty,email"`
	RoleID   string `json:"role_id" binding:"omitempty,uuid"`
}

func (server *Server) updateUser(ctx *gin.Context) {
	var req updateUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Get one user
	userID, _ := uuid.Parse(req.ID)
	user, err := server.store.GetUser(ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	roleID, err := uuid.Parse(req.RoleID)
	if err != nil {
		roleID = user.RoleID
	}

	arg := db.UpdateUserParams{
		ID:       userID,
		Username: util.FilterEmptyString(req.Username, user.Username),
		RoleID:   roleID,
		Email:    util.FilterEmptyString(req.Email, user.Email),
	}

	userUpdated, err := server.store.UpdateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violatio":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userUpdated)
}

//delete user handler
type deleteUserRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type deleteUserResponse struct {
	Message string
	UserID  uuid.UUID
}

func (server *Server) deleteUser(ctx *gin.Context) {
	var req deleteUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, _ := uuid.Parse(req.ID)

	err := server.store.DeleteUser(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := deleteUserResponse{
		Message: "User deleted whit id: ",
		UserID:  userID,
	}
	ctx.JSON(http.StatusOK, res)
}
