package server

import (
	"database/sql"
	db "github.com/JairoRiver/personal_blog_backend/pkg/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"log"
	"net/http"
)

//create role handler
type createRoleRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) createRole(ctx *gin.Context) {
	var req createRoleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	name := req.Name

	role, err := server.store.CreateRole(ctx, name)
	if err != nil {
		if pqError, ok := err.(*pq.Error); ok {
			log.Println(pqError.Code.Name())
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, role)
}

//get role handler
type getRoleRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) getRole(ctx *gin.Context) {
	var req getRoleRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	roleID, _ := uuid.Parse(req.ID)
	role, err := server.store.GetRole(ctx, roleID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, role)
}

//list roles handler
func (server *Server) listRoles(ctx *gin.Context) {
	roles, err := server.store.ListRoles(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, roles)
}

//update role
type updateRoleRequest struct {
	Name string `json:"name"`
	ID   string `json:"id" uri:"id" binding:"uuid"`
}

func (server *Server) updateRole(ctx *gin.Context) {
	var req updateRoleRequest
	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	roleId, _ := uuid.Parse(req.ID)

	arg := db.UpdateRoleParams{
		ID:   roleId,
		Name: req.Name,
	}

	role, err := server.store.UpdateRole(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, role)
}

//delete role
type deleteRoleRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type deleteRoleResponse struct {
	Message string
	RoleID  uuid.UUID
}

func (server *Server) deleteRole(ctx *gin.Context) {
	var req deleteRoleRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	roleID, _ := uuid.Parse(req.ID)
	err := server.store.DeleteRole(ctx, roleID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := deleteRoleResponse{
		Message: "Role deleted whit id: ",
		RoleID:  roleID,
	}
	ctx.JSON(http.StatusOK, res)
}
