package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/JairoRiver/personal_blog_backend/internal/db/sqlc"
	"github.com/JairoRiver/personal_blog_backend/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// createUser handler
type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type userResponse struct {
	ID        uuid.UUID
	Username  string
	Email     string
	CreatedAt time.Time
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

// createUser godoc
//
//	@Summary		Create a new User
//	@Description	Create a new user
//	@Tags			user,create
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	userResponse
//
//	@Param			user	body		createUserRequest	true	"User Data"
//	@Router			/user [post]
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

	arg := db.CreateUserParams{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := newUserResponse(user)
	ctx.JSON(http.StatusOK, response)
}

// get User handler
type getUserRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func getUserResponse(user db.User) userResponse {
	return userResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := getUserResponse(user)

	ctx.JSON(http.StatusOK, response)
}

// list user handler
func (server *Server) listUsers(ctx *gin.Context) {
	users, err := server.store.ListUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// update user handler
type updateUserRequestData struct {
	Username string `json:"username" binding:"omitempty,alphanum"`
	Email    string `json:"email" binding:"omitempty,email"`
	Password string `json:"password" binding:"omitempty"`
}
type updateUserRequestID struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) updateUser(ctx *gin.Context) {
	var userIDReq updateUserRequestID
	if err := ctx.ShouldBindUri(&userIDReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	userID, err := uuid.Parse(userIDReq.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	var reqData updateUserRequestData
	if err := ctx.ShouldBindJSON(&reqData); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateUserParams{
		ID: userID,
	}

	//Validate is the password is valid
	if len(reqData.Password) > 0 {
		hashedPassword, err := util.HashPassword(reqData.Password)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		arg.Password = pgtype.Text{String: hashedPassword, Valid: true}
	}

	//Validate is the username is valid
	if len(reqData.Username) > 0 {
		arg.Username = pgtype.Text{String: reqData.Username, Valid: true}
	}

	//Validate is the email is valid
	if len(reqData.Email) > 0 {
		arg.Email = pgtype.Text{String: reqData.Email, Valid: true}
	}

	user, err := server.store.UpdateUser(ctx, arg)

	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := userResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
	ctx.JSON(http.StatusOK, response)

}

// delete User handler
type deleteUserRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) deleteUser(ctx *gin.Context) {
	var req deleteUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.DeleteUser(ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userID)
}
