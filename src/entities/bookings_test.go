package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBookingEntity_AddBooking(t *testing.T) {
	// Clearing global state
	Bookings = nil

	entity := &BookingEntity{}
	booking := &Booking{
		Name: "John Doe",
		Date: time.Now(),
	}

	result, err := entity.AddBooking(booking)

	assert.NoError(t, err)
	assert.Equal(t, booking, result)
	assert.Len(t, Bookings, 1)
	assert.Equal(t, "John Doe", Bookings[0].Name)
}

func TestBookingEntity_CheckClassExistsOnDate(t *testing.T) {
	entity := &BookingEntity{}

	start := time.Date(2025, 5, 3, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 0, 7)

	tests := []struct {
		name      string
		classes   []Class
		checkDate time.Time
		expected  bool
	}{
		{
			name: "should return true if there is a class starting on the given date",
			classes: []Class{{
				ClassName: "Yoga",
				StartDate: start,
				EndDate:   end,
			}},
			checkDate: start,
			expected:  true,
		},
		{
			name: "should return true if there is a class ending on the given date",
			classes: []Class{{
				ClassName: "Yoga",
				StartDate: start,
				EndDate:   end,
			}},
			checkDate: end,
			expected:  true,
		},
		{
			name: "should return true if class exists during the given date range",
			classes: []Class{{
				ClassName: "Yoga",
				StartDate: start,
				EndDate:   end,
			}},
			checkDate: start.AddDate(0, 0, 3),
			expected:  true,
		},
		{
			name: "should return false if no class exists before the given date",
			classes: []Class{{
				ClassName: "Yoga",
				StartDate: start,
				EndDate:   end,
			}},
			checkDate: start.AddDate(0, 0, -1),
			expected:  false,
		},
		{
			name: "should return false if no class exists after the given date",
			classes: []Class{{
				ClassName: "Yoga",
				StartDate: start,
				EndDate:   end,
			}},
			checkDate: end.AddDate(0, 0, 1),
			expected:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset.
			Classes = tt.classes

			result := entity.CheckClassExistsOnDate(tt.checkDate)
			assert.Equal(t, tt.expected, result)
		})
	}
}
