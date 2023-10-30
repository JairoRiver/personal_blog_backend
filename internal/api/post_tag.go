package api

import (
	"database/sql"
	"net/http"

	db "github.com/JairoRiver/personal_blog_backend/internal/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// createPost handler
type createPostTag struct {
	PostId string `json:"post_id" binding:"required,uuid"`
	TagId  string `json:"tag_id" binding:"required,uuid"`
}

// createPostTag godoc
//
//	@Summary					Create a new PostTag
//	@Description				Create a new PostTag
//	@Tags						post_tag,create
//	@Accept						json
//	@Produce					json
//	@Success					200		{object}	db.PostsTag
//
//	@Param						post	body		createPostTag	true	"post_tag Data"
//	@securityDefinitions.apiKey	token
//	@in							header
//	@name						Authorization
//	@Security					JWT
//	@Router						/admin/post-tag [post]
func (server *Server) createPostTag(ctx *gin.Context) {
	var req createPostTag
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	post_id, err := uuid.Parse(req.PostId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	tag_id, err := uuid.Parse(req.TagId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreatePostTagParams{
		PostID: post_id,
		TagID:  tag_id,
	}

	postTag, err := server.store.CreatePostTag(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, postTag)
}

// delete PostTag By Id handler
type deletePostTagRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// deletePostTag godoc
//
//	@Summary					Delete a PostTag by Id
//	@Description				Delete one post_tag on the admin panel
//	@Tags						post_tag,delete
//	@Accept						json
//	@Produce					json
//	@Success					200	{object}	db.PostsTag
//
//	@Param						id	path		string	true	"id"
//
//	@securityDefinitions.apiKey	token
//	@in							header
//	@name						Authorization
//	@Security					JWT
//
//	@Router						/admin/post-tag/{id} [delete]
func (server *Server) deletePostTag(ctx *gin.Context) {
	var req deletePostTagRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	postTagID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.DeletePostTag(ctx, postTagID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, postTagID)
}
