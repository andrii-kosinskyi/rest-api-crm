package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"sort"
	"strings"
	"time"

	"CRM/internal/model"
	"CRM/internal/repository"
	"CRM/internal/validation"
)

type Repository interface {
	Create(c *model.Client) error
	GetByID(id int64) (*model.Client, error)
	GetAll() ([]*model.Client, error)
}

type Validator interface {
	OnCreate(c *model.Client) []string
}

type Client struct {
	ctx       context.Context
	repo      Repository
	validator Validator
}

func NewClient(ctx context.Context) *Client {
	repo := repository.NewInMemory(ctx)
	valid := validation.NewClientValidator(ctx)

	return &Client{ctx, repo, valid}
}

func (srv *Client) Create(newClient *model.Client) error {
	errMessages := srv.validator.OnCreate(newClient)
	if len(errMessages) != 0 {
		return errors.New(strings.Join(errMessages, ";\n"))
	}

	err := srv.repo.Create(newClient)
	if err != nil {
		return err
	}

	return nil
}

func (srv *Client) GetByID(id int64) (*model.Client, error) {
	c, err := srv.repo.GetByID(id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("client not found by id: %d", id))
	}

	return c, nil
}

func (srv *Client) GetAll() ([]*model.Client, error) {
	c, err := srv.repo.GetAll()
	if err != nil {
		return nil, errors.New("invalid request")
	}

	return c, nil
}

func (srv *Client) AssignLead(l *model.Lead) (*model.Client, error) {
	var availableClients []*model.Client
	if l.DateTime.IsZero() {
		l.DateTime = time.Now()
	}

	clients, err := srv.GetAll()
	if err != nil {
		log.Errorf("endpoint assignLead error: %v", err)
		return nil, errors.New("invalid request")
	}

	leadTime, _ := time.Parse("15:04", l.DateTime.Format("15:04"))
	for _, client := range clients {
		startTime, _ := time.Parse("15:04", client.WorkingHoursStart)
		endTime, _ := time.Parse("15:04", client.WorkingHoursEnd)
		fmt.Printf("leadTime %s", l.DateTime)
		fmt.Printf("startTime %s endTime %s", startTime, endTime)

		if leadTime.After(startTime) && leadTime.Before(endTime) {
			availableClients = append(availableClients, client)
		}
	}

	if len(availableClients) == 0 {
		return nil, errors.New("not found available client")
	}

	sort.Slice(availableClients, func(i, j int) bool {
		return availableClients[i].Priority > availableClients[j].Priority
	})

	for _, client := range availableClients {
		if client.LeadCapacity != 0 {
			client.LeadCapacity -= 1
			return availableClients[0], nil
		}
	}

	return nil, errors.New("not found available client")
}
