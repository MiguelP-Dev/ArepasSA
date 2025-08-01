package handlers

import (
	"net/http"
	"strconv"

	"ArepasSA/internal/models"
	"ArepasSA/internal/services"

	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	service *services.ClientService
}

func NewClientHandler(service *services.ClientService) *ClientHandler {
	return &ClientHandler{service: service}
}

// CreateClient godoc
// @Summary Crear un nuevo cliente
// @Description Crea un nuevo cliente con los datos proporcionados
// @Tags clients
// @Accept json
// @Produce json
// @Param client body models.Client true "Datos del cliente"
// @Success 201 {object} models.Client
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /clients [post]
func (h *ClientHandler) CreateClient(c *gin.Context) {
	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateClient(&client); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, client)
}

// GetClient godoc
// @Summary Obtener un cliente por ID
// @Description Obtiene los detalles de un cliente específico
// @Tags clients
// @Produce json
// @Param id path int true "ID del cliente"
// @Success 200 {object} models.Client
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /clients/{id} [get]
func (h *ClientHandler) GetClient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de cliente inválido"})
		return
	}

	client, err := h.service.GetClient(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente no encontrado"})
		return
	}

	c.JSON(http.StatusOK, client)
}

// GetAllClients godoc
// @Summary Listar todos los clientes
// @Description Obtiene una lista de todos los clientes, con opción de filtrado
// @Tags clients
// @Produce json
// @Param active query bool false "Filtrar solo clientes activos"
// @Param search query string false "Texto para búsqueda"
// @Success 200 {array} models.Client
// @Failure 500 {object} map[string]string
// @Router /clients [get]
func (h *ClientHandler) GetAllClients(c *gin.Context) {
	activeOnly := c.DefaultQuery("active", "true") == "true"
	search := c.Query("search")

	clients, err := h.service.GetAllClients(activeOnly, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, clients)
}

// UpdateClient godoc
// @Summary Actualizar un cliente
// @Description Actualiza los datos de un cliente existente
// @Tags clients
// @Accept json
// @Produce json
// @Param id path int true "ID del cliente"
// @Param client body models.Client true "Datos actualizados del cliente"
// @Success 200 {object} models.Client
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /clients/{id} [put]
func (h *ClientHandler) UpdateClient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de cliente inválido"})
		return
	}

	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client.ID = uint(id)
	if err := h.service.UpdateClient(&client); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, client)
}

// DeleteClient godoc
// @Summary Eliminar un cliente
// @Description Elimina permanentemente un cliente (no recomendado, usar desactivar)
// @Tags clients
// @Produce json
// @Param id path int true "ID del cliente"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /clients/{id} [delete]
func (h *ClientHandler) DeleteClient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de cliente inválido"})
		return
	}

	if err := h.service.DeleteClient(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// DeactivateClient godoc
// @Summary Desactivar un cliente
// @Description Desactiva un cliente (eliminación suave)
// @Tags clients
// @Produce json
// @Param id path int true "ID del cliente"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /clients/{id}/deactivate [patch]
func (h *ClientHandler) DeactivateClient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de cliente inválido"})
		return
	}

	if err := h.service.SoftDeleteClient(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cliente desactivado correctamente"})
}

// AddClientComment godoc
// @Summary Añadir comentario a cliente
// @Description Añade un comentario o preferencia a un cliente
// @Tags clients
// @Accept json
// @Produce json
// @Param id path int true "ID del cliente"
// @Param comment body models.ClientComment true "Datos del comentario"
// @Success 201 {object} models.ClientComment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /clients/{id}/comments [post]
func (h *ClientHandler) AddClientComment(c *gin.Context) {
	clientID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de cliente inválido"})
		return
	}

	var comment models.ClientComment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Obtener el usuario autenticado (si hay sistema de auth)
	user, _ := c.Get("user")
	comment.Author = user.(string)
	comment.ClientID = uint(clientID)

	if err := h.service.AddComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

// GetClientPreferences godoc
// @Summary Obtener preferencias de cliente
// @Description Obtiene todas las preferencias registradas de un cliente
// @Tags clients
// @Produce json
// @Param id path int true "ID del cliente"
// @Success 200 {array} models.ClientComment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /clients/{id}/preferences [get]
func (h *ClientHandler) GetClientPreferences(c *gin.Context) {
	clientID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de cliente inválido"})
		return
	}

	preferences, err := h.service.GetClientPreferences(uint(clientID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, preferences)
}

// GetClientsWithPreferences godoc
// @Summary Obtener clientes con preferencias
// @Description Obtiene todos los clientes que tienen preferencias registradas
// @Tags clients
// @Produce json
// @Success 200 {array} models.Client
// @Failure 500 {object} map[string]string
// @Router /clients/with-preferences [get]
func (h *ClientHandler) GetClientsWithPreferences(c *gin.Context) {
	clients, err := h.service.GetClientsWithPreferences()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, clients)
}
