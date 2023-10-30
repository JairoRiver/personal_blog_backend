package api

import (
	"database/sql"
	"net/http"

	_ "github.com/JairoRiver/personal_blog_backend/internal/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// get Post By Id Private handler
type getPostByIdPublicRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// getPostByIdPublic godoc
//
//	@Summary		Get a Post by Id Public
//	@Description	Recive the one post public
//	@Tags			post,get
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	db.Post
//
//	@Param			id	path		string	true	"id"
//	@Router			/post/{id} [get]
func (server *Server) getPostByIdPublic(ctx *gin.Context) {
	var req getPostByIdPublicRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	postID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	post, err := server.store.GetPostByIdPublic(ctx, postID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, post)
}

// get Post By Category Public handler
type getPostByCategoryPublicRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// getPostByCategoryPublic godoc
//
//	@Summary		Get a Post by Category Public
//	@Description	Recive the one post by Category
//	@Tags			post,list
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	db.Post
//
//	@Param			id	path		string	true	"id"
//	@Router			/category-post/{id} [get]
func (server *Server) getPostByCategoryPublic(ctx *gin.Context) {
	var req getPostByCategoryPublicRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	categoryID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	post, err := server.store.GetPostByCategoryPublic(ctx, categoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, post)
}

// get Post By Tag Public handler
type getPostByTagPublicRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// getPostByTagPublic godoc
//
//	@Summary		Get a Post by tag Public
//	@Description	Recive the one post by Tag
//	@Tags			post,list
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	db.Post
//
//	@Param			id	path		string	true	"id"
//	@Router			/tag-post/{id} [get]
func (server *Server) getPostByTagPublic(ctx *gin.Context) {
	var req getPostByTagPublicRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	tagID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	post, err := server.store.GetPostByTagPublic(ctx, tagID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, post)
}

// list Post Public handler
// listPostPublic godoc
//
//	@Summary		List all Posts
//	@Description	Recive all posts publics
//	@Tags			post,list
//	@Produce		json
//	@Success		200	{object}	db.ListPostsPublicRow
//	@Router			/posts [get]
func (server *Server) listPostsPublic(ctx *gin.Context) {

	posts, err := server.store.ListPostsPublic(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, posts)
}
