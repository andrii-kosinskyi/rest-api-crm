package repository

import (
	"context"
	"errors"

	"CRM/internal/model"
)

type InMemoryRepo struct {
	ctx   context.Context
	store []*model.Client
}

const defaultLeadCap = 100

var store = []*model.Client{
	{
		ID:                1,
		Name:              "MXP",
		WorkingHoursStart: "08:00",
		WorkingHoursEnd:   "17:00",
		Priority:          2,
		LeadCapacity:      300,
	},
	{
		ID:                2,
		Name:              "NOVA POSHTA",
		WorkingHoursStart: "06:00",
		WorkingHoursEnd:   "18:45",
		Priority:          1,
		LeadCapacity:      534,
	},
	{
		ID:                3,
		Name:              "BIOSPHERE",
		WorkingHoursStart: "12:00",
		WorkingHoursEnd:   "21:00",
		Priority:          3,
		LeadCapacity:      122,
	}}

func NewInMemory(ctx context.Context) *InMemoryRepo {
	return &InMemoryRepo{ctx, store}
}

// Create - create client
func (r *InMemoryRepo) Create(c *model.Client) error {
	c.ID = int64(len(store) + 1)
	c.LeadCapacity = defaultLeadCap
	r.store = append(r.store, c)
	return nil
}

// GetByID - get client by ID
func (r *InMemoryRepo) GetByID(id int64) (*model.Client, error) {
	for _, v := range r.store {
		if v.ID == id {
			return v, nil
		}
	}
	return nil, errors.New("Not Found")
}

// GetAll - return all clients
func (r *InMemoryRepo) GetAll() ([]*model.Client, error) {
	cSlice := make([]*model.Client, 0, len(r.store))
	for _, v := range r.store {
		cSlice = append(cSlice, v)
	}
	return cSlice, nil
}
