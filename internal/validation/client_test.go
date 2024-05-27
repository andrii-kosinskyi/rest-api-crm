package validation

import (
	"context"
	"testing"

	"CRM/internal/model"
)

func TestClientValidator_OnCreate(t *testing.T) {
	tests := []struct {
		name           string
		client         *model.Client
		expectedErrors []string
	}{
		{
			name: "Valid client",
			client: &model.Client{
				Name:              "ValidName",
				Priority:          1,
				WorkingHoursStart: "09:00",
				WorkingHoursEnd:   "17:00",
			},
			expectedErrors: []string{},
		},
		{
			name: "Name too short",
			client: &model.Client{
				Name:              "AB",
				Priority:          1,
				WorkingHoursStart: "09:00",
				WorkingHoursEnd:   "17:00",
			},
			expectedErrors: []string{"Name is too short"},
		},
		{
			name: "Priority too low",
			client: &model.Client{
				Name:              "ValidName",
				Priority:          -1,
				WorkingHoursStart: "09:00",
				WorkingHoursEnd:   "17:00",
			},
			expectedErrors: []string{"Priority is too low"},
		},
		{
			name: "Invalid working hours start format",
			client: &model.Client{
				Name:              "ValidName",
				Priority:          1,
				WorkingHoursStart: "nine o'clock",
				WorkingHoursEnd:   "17:00",
			},
			expectedErrors: []string{"invalid working hours start format, should be HH:MM"},
		},
		{
			name: "Invalid working hours end format",
			client: &model.Client{
				Name:              "ValidName",
				Priority:          1,
				WorkingHoursStart: "09:00",
				WorkingHoursEnd:   "5:00 PM",
			},
			expectedErrors: []string{"invalid working hours end format, should be HH:MM"},
		},
		{
			name: "Working hours start not before end",
			client: &model.Client{
				Name:              "ValidName",
				Priority:          1,
				WorkingHoursStart: "17:00",
				WorkingHoursEnd:   "09:00",
			},
			expectedErrors: []string{"working hours start must be before working hours end"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cv := NewClientValidator(context.Background())
			errors := cv.OnCreate(tt.client)

			if len(errors) != len(tt.expectedErrors) {
				t.Errorf("expected %d errors, got %d", len(tt.expectedErrors), len(errors))
			}

			for i, err := range errors {
				if err != tt.expectedErrors[i] {
					t.Errorf("expected error '%s', got '%s'", tt.expectedErrors[i], err)
				}
			}
		})
	}
}
