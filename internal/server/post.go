package server

import (
	"database/sql"
	"github.com/JairoRiver/personal_blog_backend/internal/util"
	db "github.com/JairoRiver/personal_blog_backend/pkg/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

//create post handler
type createPostRequest struct {
	Title    string `json:"title" binding:"required"`
	Subtitle string `json:"subtitle" binding:"required"`
	Content  string `json:"content" binding:"required"`
	UserID   string `json:"user_id" binding:"required,uuid"`
}
type createPostResponse struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

func (server *Server) createPost(ctx *gin.Context) {
	var req createPostRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, _ := uuid.Parse(req.UserID)
	arg := db.CreatePostParams{
		UserID:   userID,
		Title:    req.Title,
		Subtitle: req.Subtitle,
		Content:  req.Content,
	}

	post, err := server.store.CreatePost(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	res := createPostResponse{
		ID:        post.ID,
		Title:     post.Title,
		CreatedAt: post.CreatedAt,
	}
	ctx.JSON(http.StatusOK, res)
}

//get post handler
type getPostRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type getPostResponse struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (server *Server) getPost(ctx *gin.Context) {
	var req getPostRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	postID, _ := uuid.Parse(req.ID)

	post, err := server.store.GetPost(ctx, postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := getPostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Subtitle:  post.Subtitle,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, res)
}

//list post handler
type listPostRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=100"`
}

func (server *Server) listPosts(ctx *gin.Context) {
	var req listPostRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListPostsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	posts, err := server.store.ListPosts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, posts)
}

//update post handler
type updatePostRequest struct {
	ID       string `json:"id" uri:"id" binding:"required,uuid"`
	Title    string `json:"title" binding:"omitempty"`
	Subtitle string `json:"subtitle" binding:"omitempty"`
	Content  string `json:"content" binding:"omitempty"`
}

func (server *Server) updatePost(ctx *gin.Context) {
	var req updatePostRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	postID, _ := uuid.Parse(req.ID)
	post, err := server.store.GetPost(ctx, postID)
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

	arg := db.UpdatePostParams{
		ID:       postID,
		Title:    util.FilterEmptyString(req.Title, post.Title),
		Subtitle: util.FilterEmptyString(req.Subtitle, post.Subtitle),
		Content:  util.FilterEmptyString(req.Content, post.Content),
	}

	postUpdated, err := server.store.UpdatePost(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, postUpdated)
}

type deletePostRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type deletePostResponse struct {
	Message string
	PostID  uuid.UUID
}

func (server *Server) deletePost(ctx *gin.Context) {
	var req deletePostRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	postID, _ := uuid.Parse(req.ID)
	err := server.store.DeletePost(ctx, postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := deletePostResponse{
		Message: "Post deleted whit id: ",
		PostID:  postID,
	}
	ctx.JSON(http.StatusOK, res)
}