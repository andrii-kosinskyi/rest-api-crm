package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"

	"CRM/internal/model"
	"CRM/internal/service"
)

type Handler struct {
	clientService *service.Client
}

func NewHandler(ctx context.Context) *Handler {
	return &Handler{clientService: service.NewClient(ctx)}
}

// @Summary Create a new client
// @Description Create a new client with the provided information
// @Tags clients
// @Accept json
// @Produce json
// @Param client body model.Client true "Client Information"
// @Success 201 {string} string "client id: {id}"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /clients [post]
func (h *Handler) createClient(c echo.Context) error {
	var newClient model.Client
	err := c.Bind(&newClient)
	if err != nil {
		log.Errorf("endpoint createClient error: %v", err)
		return c.String(http.StatusInternalServerError, wrapApiErr("bad request"))
	}
	err = h.clientService.Create(&newClient)
	if err != nil {
		log.Errorf("endpoint createClient error: %v", err)
		return c.String(http.StatusBadRequest, strconv.FormatInt(newClient.ID, 10))
	}
	return c.String(http.StatusCreated, fmt.Sprintf("client id: %d", newClient.ID))
}

// @Summary Get a client by ID
// @Description Retrieve a client by their ID
// @Tags clients
// @Accept json
// @Produce json
// @Param id path int true "Client ID"
// @Success 200 {object} model.Client "client data"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /clients/{id} [get]
func (h *Handler) getClient(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Errorf("endpoint createClient error: %v", err)
		return c.String(http.StatusBadRequest, wrapApiErr("bad request"))
	}

	client, err := h.clientService.GetByID(id)
	if err != nil {
		log.Errorf("endpoint createClient error: %v", err)
		return c.String(http.StatusBadRequest, wrapApiErr(err.Error()))
	}

	b, err := json.Marshal(client)
	if err != nil {
		log.Errorf("endpoint createClient error: %v", err)
		return c.String(http.StatusInternalServerError, wrapApiErr("bad request"))
	}
	return c.String(http.StatusOK, string(b))
}

// @Summary Get all clients
// @Description Retrieve a list of all clients
// @Tags clients
// @Accept json
// @Produce json
// @Success 200 {array} model.Client "list of clients"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /clients [get]
func (h *Handler) getClients(c echo.Context) error {
	clients, err := h.clientService.GetAll()
	if err != nil {
		log.Errorf("endpoint createClient error: %v", err)
		return c.String(http.StatusBadRequest, wrapApiErr(err.Error()))
	}

	b, err := json.Marshal(clients)
	if err != nil {
		log.Errorf("endpoint createClient error: %v", err)
		return c.String(http.StatusInternalServerError, wrapApiErr("bad request"))
	}

	return c.String(http.StatusOK, string(b))
}

// @Summary Assign a lead to a client
// @Description Assign a new lead to a client with the provided information
// @Tags clients
// @Accept json
// @Produce json
// @Param lead body model.Lead true "Lead Information"
// @Success 200 {object} model.Client "assigned client data"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /clients/assign-lead [post]
func (h *Handler) assignLead(c echo.Context) error {
	var newLead model.Lead
	err := c.Bind(&newLead)
	if err != nil {
		log.Errorf("endpoint createClient error: %v", err)
		return c.String(http.StatusInternalServerError, wrapApiErr("bad request"))
	}
	assignedClient, err := h.clientService.AssignLead(&newLead)
	if err != nil {
		log.Errorf("endpoint assignLead error: %v", err)
		return c.String(http.StatusBadRequest, wrapApiErr(err.Error()))
	}

	b, err := json.Marshal(assignedClient)
	if err != nil {
		log.Errorf("endpoint assignLead error: %v", err)
		return c.String(http.StatusInternalServerError, wrapApiErr("bad request"))
	}

	return c.String(http.StatusOK, string(b))
}
