package api

import (
	"database/sql"
	"net/http"

	db "github.com/JairoRiver/personal_blog_backend/internal/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// createCategory handler
type createCategoryRequest struct {
	Name string `json:"name" binding:"required,ascii"`
}

// createCategory godoc
//
//	@Summary					Create a new Category
//	@Description				Create a new Category
//	@Tags						category,create
//	@Accept						json
//	@Produce					json
//	@Success					200			{object}	db.Category
//
//	@Param						category	body		createCategoryRequest	true	"Category Data"
//	@securityDefinitions.apiKey	token
//	@in							header
//	@name						Authorization
//	@Security					JWT
//	@Router						/category [post]
func (server *Server) createCategory(ctx *gin.Context) {
	var req createCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	category, err := server.store.CreateCategory(ctx, req.Name)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, category)
}

// getCategory handler
type getCategoryRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// getCategory godoc
//
//	@Summary					Get a Category
//	@Description				Recive the one category information from a id
//	@Tags						category,get
//	@Accept						json
//	@Produce					json
//	@Success					200	{object}	db.Category
//
//	@Param						id	path		string	true	"id"
//
//	@securityDefinitions.apiKey	token
//	@in							header
//	@name						Authorization
//	@Security					JWT
//
//	@Router						/category/{id} [get]
func (server *Server) getCategory(ctx *gin.Context) {
	var req getCategoryRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	categoryID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	category, err := server.store.GetCategory(ctx, categoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)
}

// listCategory handler
//
// listCategories godoc
//
//	@Summary					List Categories
//	@Description				Recive all categories
//	@Tags						category,list
//	@Accept						json
//	@Produce					json
//	@Success					200	{object}	db.Category
//	@securityDefinitions.apiKey	token
//	@in							header
//	@name						Authorization
//	@Security					JWT
//	@Router						/categories [get]
func (server *Server) listCategories(ctx *gin.Context) {
	categories, err := server.store.ListCategories(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

// updateCategory handler
type updateCategoryRequestData struct {
	Name string `json:"name" binding:"omitempty,ascii"`
}
type updateCategoryRequestID struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// updateCategory godoc
//
//	@Summary					Update Category
//	@Description				Update the category information
//	@Tags						category,update
//	@Accept						json
//	@Produce					json
//	@Param						id			path		string						true	"id"
//	@Param						category	body		updateCategoryRequestData	true	"Category Data"
//	@Success					200			{object}	db.Category
//	@securityDefinitions.apiKey	token
//	@in							header
//	@name						Authorization
//	@Security					JWT
//	@Router						/category/{id} [put]
func (server *Server) updateCategory(ctx *gin.Context) {
	var categoryIDReq updateCategoryRequestID
	if err := ctx.ShouldBindUri(&categoryIDReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	categoryID, err := uuid.Parse(categoryIDReq.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	var reqData updateCategoryRequestData
	if err := ctx.ShouldBindJSON(&reqData); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCategoryParams{
		ID: categoryID,
	}

	//Validate is the name is valid
	if len(reqData.Name) > 0 {
		arg.Name = pgtype.Text{String: reqData.Name, Valid: true}
	}

	category, err := server.store.UpdateCategory(ctx, arg)

	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)
}

// delete Category handler
type deleteCategoryRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// deleteCategory godoc
//
//	@Summary					Delete Category
//	@Description				Delete the category register
//	@Tags						category,delete
//	@Accept						json
//	@Produce					json
//	@Param						id	path		string	true	"id"
//	@Success					200	{object}	uuid.UUID
//	@securityDefinitions.apiKey	token
//	@in							header
//	@name						Authorization
//	@Security					JWT
//	@Router						/category/{id} [delete]
func (server *Server) deleteCategory(ctx *gin.Context) {
	var req deleteCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	categoryID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.DeleteCategory(ctx, categoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, categoryID)
}
