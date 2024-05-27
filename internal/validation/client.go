package validation

import (
	"context"
	"time"

	"CRM/internal/model"
)

type ClientValidator struct {
	ctx context.Context
}

func NewClientValidator(ctx context.Context) *ClientValidator {
	return &ClientValidator{ctx}
}

const (
	nameLength  = 3
	minPriority = 0
)

func (cv *ClientValidator) OnCreate(c *model.Client) []string {
	var errMessages []string
	if len(c.Name) < nameLength {
		errMessages = append(errMessages, "Name is too short")
	}

	if c.Priority < minPriority {
		errMessages = append(errMessages, "Priority is too low")
	}

	layout := "15:04"
	startTime, err := time.Parse(layout, c.WorkingHoursStart)
	if err != nil {
		errMessages = append(errMessages, "invalid working hours start format, should be HH:MM")
	}

	endTime, err := time.Parse(layout, c.WorkingHoursEnd)
	if err != nil {
		errMessages = append(errMessages, "invalid working hours end format, should be HH:MM")
	}

	if !startTime.Before(endTime) {
		errMessages = append(errMessages, "working hours start must be before working hours end")
	}

	return errMessages
}
