package api

import (
	"database/sql"
	"net/http"

	db "github.com/JairoRiver/personal_blog_backend/internal/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// createPost handler
type createPostRequest struct {
	Title      string `json:"title" binding:"required"`
	Subtitle   string `json:"subtitle" binding:"required"`
	Content    string `json:"content" binding:"required"`
	Publicated bool   `json:"publicated"`
	CategoryId string `json:"category_id" binding:"required,uuid"`
}

// createPost godoc
//
//	@Summary					Create a new Post
//	@Description				Create a new Post
//	@Tags						post,create
//	@Accept						json
//	@Produce					json
//	@Success					200		{object}	db.Post
//
//	@Param						post	body		createPostRequest	true	"post Data"
//	@securityDefinitions.apiKey	token
//	@in							header
//	@name						Authorization
//	@Security					JWT
//	@Router						/admin/post [post]
func (server *Server) createPost(ctx *gin.Context) {
	var req createPostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	category_id, err := uuid.Parse(req.CategoryId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreatePostParams{
		CategoryID: category_id,
		Title:      req.Title,
		Subtitle:   req.Subtitle,
		Content:    req.Content,
		Publicated: req.Publicated,
	}

	post, err := server.store.CreatePost(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, post)
}

// get Post By Id Private handler
type getPostByIdPrivateRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// getPostByIdPrivate godoc
//
//	@Summary					Get a Post by Id Private
//	@Description				Recive the one post on the admin panel
//	@Tags						post,get
//	@Accept						json
//	@Produce					json
//	@Success					200	{object}	db.Post
//
//	@Param						id	path		string	true	"id"
//
//	@securityDefinitions.apiKey	token
//	@in							header
//	@name						Authorization
//	@Security					JWT
//
//	@Router						/admin/post/{id} [get]
func (server *Server) getPostByIdPrivate(ctx *gin.Context) {
	var req getPostByIdPrivateRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	postID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	post, err := server.store.GetPostByIdPrivate(ctx, postID)
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

// get Post By Category Private handler
type getPostByCategoryPrivateRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// getPostByCategoryPrivate godoc
//
//	@Summary					Get a Post by Category Private
//	@Description				Recive the one post on the admin panel by Category
//	@Tags						post,list
//	@Accept						json
//	@Produce					json
//	@Success					200	{object}	db.Post
//
//	@Param						id	path		string	true	"id"
//
//	@securityDefinitions.apiKey	token
//	@in							header
//	@name						Authorization
//	@Security					JWT
//
//	@Router						/admin/category-post/{id} [get]
func (server *Server) getPostByCategoryPrivate(ctx *gin.Context) {
	var req getPostByCategoryPrivateRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	categoryID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	post, err := server.store.GetPostByCategoryPrivate(ctx, categoryID)
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

// get Post By Tag Private handler
type getPostByTagPrivateRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// getPostByTagPrivate godoc
//
//	@Summary					Get a Post by tag Private
//	@Description				Recive the one post on the admin panel by Tag
//	@Tags						post,list
//	@Accept						json
//	@Produce					json
//	@Success					200	{object}	db.Post
//
//	@Param						id	path		string	true	"id"
//
//	@securityDefinitions.apiKey	token
//	@in							header
//	@name						Authorization
//	@Security					JWT
//
//	@Router						/admin/tag-post/{id} [get]
func (server *Server) getPostByTagPrivate(ctx *gin.Context) {
	var req getPostByTagPrivateRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	tagID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	post, err := server.store.GetPostByTagPrivate(ctx, tagID)
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

// list Post Private handler
// listPostPrivate godoc
//
//	@Summary					List all Posts Private
//	@Description				Recive all posts on the admin panel
//	@Tags						post,list
//	@Produce					json
//	@Success					200	{object}	db.ListPostsPrivateRow
//
//	@securityDefinitions.apiKey	token
//	@in							header
//	@name						Authorization
//	@Security					JWT
//
//	@Router						/admin/posts [get]
func (server *Server) listPostsPrivate(ctx *gin.Context) {

	posts, err := server.store.ListPostsPrivate(ctx)
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

// updatePost handler
type updatePostRequest struct {
	Title      pgtype.Text `json:"title"`
	Subtitle   pgtype.Text `json:"subtitle"`
	Content    pgtype.Text `json:"content"`
	Publicated pgtype.Bool `json:"publicated"`
	CategoryId pgtype.UUID `json:"category_id"`
}

// updatePost godoc
//
//	@Summary					Update a Post
//	@Description				Update a new Post
//	@Tags						post,update
//	@Accept						json
//	@Produce					json
//	@Success					200		{object}	db.Post
//
//	@Param						post	body		updatePostRequest	true	"post Data"
//	@Param						id		path		string				true	"id"
//	@securityDefinitions.apiKey	token
//	@in							header
//	@name						Authorization
//	@Security					JWT
//	@Router						/admin/post/{id} [put]
func (server *Server) updatePost(ctx *gin.Context) {
	var req updatePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var reqId getPostByIdPrivateRequest
	if err := ctx.ShouldBindUri(&reqId); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	post_id, err := uuid.Parse(reqId.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Validate the input parameters
	arg := db.UpdatePostParams{
		ID: post_id,
	}
	// Title
	if req.Title.Valid {
		arg.Title = req.Title
	}

	// Subtitle
	if req.Subtitle.Valid {
		arg.Subtitle = req.Subtitle
	}

	// Content
	if req.Content.Valid {
		arg.Content = req.Content
	}

	// Publicated
	if req.Publicated.Valid {
		arg.Publicated = req.Publicated
	}

	// CategoryId
	if req.CategoryId.Valid {
		arg.CategoryID = req.CategoryId
	}

	post, err := server.store.UpdatePost(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, post)
}

// delete Post By Id handler
type deletePostRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// deletePost godoc
//
//	@Summary					Delete a Post by Id
//	@Description				Delete one post on the admin panel
//	@Tags						post,delete
//	@Accept						json
//	@Produce					json
//	@Success					200	{object}	db.Post
//
//	@Param						id	path		string	true	"id"
//
//	@securityDefinitions.apiKey	token
//	@in							header
//	@name						Authorization
//	@Security					JWT
//
//	@Router						/admin/post/{id} [delete]
func (server *Server) deletePost(ctx *gin.Context) {
	var req deletePostRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	postID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.DeletePost(ctx, postID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, postID)
}
