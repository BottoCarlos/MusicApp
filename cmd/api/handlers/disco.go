package handlers

import (
	"net/http"
	"strconv"

	"github.com/BottoCarlos/MusicApp/internal/disco"
	"github.com/BottoCarlos/MusicApp/internal/domain"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service disco.ServiceInterface
}

func NewHandler(service disco.ServiceInterface) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAll() gin.HandlerFunc {
	return func(context *gin.Context) {
		//Request

		//Process
		data := h.service.GetAll()
		//Response
		context.JSON(http.StatusOK, data)
	}
}

func (h *Handler) GetByID() gin.HandlerFunc {
	return func(context *gin.Context) {
		//request
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}
		//process
		data := h.service.GetByID(id)
		//response
		context.JSON(http.StatusOK, data)
	}
}

func (h *Handler) Add() gin.HandlerFunc {
	return func(context *gin.Context) {
		//request
		var disco domain.Disco
		if err := context.ShouldBindJSON(&disco); err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}
		//process
		h.service.Add(disco)
		//response
		context.JSON(http.StatusOK, "success")
	}
}

func (h *Handler) Update() gin.HandlerFunc {
	return func(context *gin.Context) {
		//request
		var disco domain.Disco
		if err := context.ShouldBindJSON(&disco); err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}
		//process
		h.service.Update(disco)
		//response
		context.JSON(http.StatusOK, "success")
	}
}

func (h *Handler) Delete() gin.HandlerFunc {
	return func(context *gin.Context) {
		//request
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}
		//process
		h.service.Delete(id)
		//response
		context.JSON(http.StatusOK, "success")
	}
}
