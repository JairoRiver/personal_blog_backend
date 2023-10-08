package api

import (
	"database/sql"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	db "github.com/JairoRiver/personal_blog_backend/internal/db/sqlc"
	"github.com/JairoRiver/personal_blog_backend/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
)

const tagBucketPath = "tags"

type createTagRequest struct {
	Name string                `form:"name" binding:"required,alphanum"`
	Logo *multipart.FileHeader `form:"logo" binding:"required"`
}

// createTag godoc
//
//	@Summary					Create a new Tag
//	@Description				Create a new Tag, the image are upload to S3 services
//	@Tags						tag,create
//	@Accept						multipart/form-data
//	@Produce					json
//	@Success					200		{object}	db.Tag
//
//	@Param						name	formData	string	true	"This is the tag name"
//
//	@Param						logo	formData	file	true	"This is a image"
//
//	@securityDefinitions.apiKey	token
//	@in							header
//	@name						Authorization
//	@Security					JWT
//	@Router						/tag [post]
func (server *Server) createTag(ctx *gin.Context) {
	var req createTagRequest
	if err := ctx.ShouldBindWith(&req, binding.FormMultipart); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//tagBucketPath := "tags"
	objectName := req.Name + util.RandomString(4)

	// Save image
	fileContent, err := req.Logo.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	byteContainer, err := io.ReadAll(fileContent)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	tagURL, err := server.assetStore.UploadImage(ctx, byteContainer, tagBucketPath, objectName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateTagParams{
		Name:     req.Name,
		ImageUrl: tagURL,
	}

	tag, err := server.store.CreateTag(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			deleteErr := server.assetStore.DeleteImage(ctx, tagBucketPath, objectName)
			if deleteErr != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}

			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, tag)
}

// get Tag handler
type getTagRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// getTag godoc
//
//	@Summary		Get a Tag
//	@Description	Recive the one tag information from a id
//	@Tags			tag,get
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	db.Tag
//
//	@Param			id	path		string	true	"id"
//	@Router			/tag/{id} [get]
func (server *Server) getTag(ctx *gin.Context) {
	var req getTagRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	tagID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	tag, err := server.store.GetTag(ctx, tagID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, tag)
}

// list tag handler

// listTags godoc
//
//	@Summary		List Tags
//	@Description	Recive all tags
//	@Tags			tag,list
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	db.Tag
//	@Router			/tags [get]
func (server *Server) listTags(ctx *gin.Context) {
	tags, err := server.store.ListTags(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, tags)
}

// delete Tag handler

type deleteTagRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// deleteTag godoc
//
//	@Summary					Delete Tag
//	@Description				Delete the tag register
//	@Tags						tag,delete
//	@Accept						json
//	@Produce					json
//	@Param						id	path		string	true	"id"
//	@Success					200	{object}	uuid.UUID
//	@securityDefinitions.apiKey	token
//	@in							header
//	@name						Authorization
//	@Security					JWT
//	@Router						/tag/{id} [delete]
func (server *Server) deleteTag(ctx *gin.Context) {
	var req deleteTagRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	tagID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// get tag info
	tag, err := server.store.GetTag(ctx, tagID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// delete image
	s := tag.ImageUrl
	last := s[strings.LastIndex(s, "/")+1:]
	name := strings.Split(last, ".")[0]

	deleteErr := server.assetStore.DeleteImage(ctx, tagBucketPath, name)
	if deleteErr != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.DeleteTag(ctx, tagID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, tagID)
}
