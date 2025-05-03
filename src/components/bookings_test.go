package components

import (
	"errors"
	"testing"
	"time"

	"github.com/Vidyuallatha/glofox/src/entities"
	"github.com/stretchr/testify/assert"
)

// MockBookingRepository implements BookingRepository interface
type MockBookingRepository struct {
	CheckClassExistsOnDateFn func(time.Time) bool
	AddBookingFn             func(*entities.Booking) (*entities.Booking, error)
}

func (m *MockBookingRepository) CheckClassExistsOnDate(t time.Time) bool {
	if m.CheckClassExistsOnDateFn != nil {
		return m.CheckClassExistsOnDateFn(t)
	}
	return false
}

func (m *MockBookingRepository) AddBooking(b *entities.Booking) (*entities.Booking, error) {
	if m.AddBookingFn != nil {
		return m.AddBookingFn(b)
	}
	return nil, errors.New("not implemented")
}

func TestBookingsComponent_Valid(t *testing.T) {
	bc := &BookingsComponent{}

	tests := []struct {
		name     string
		form     *entities.Booking
		expected []string
	}{
		{
			name: "should not have any validation errors if booking is valid",
			form: &entities.Booking{
				Name: "Yoga",
				Date: time.Now(),
			},
			expected: nil,
		},
		{
			name: "should add validation error for name if it is missing",
			form: &entities.Booking{
				Name: "",
				Date: time.Now(),
			},
			expected: []string{"name is required"},
		},
		{
			name: "should add validation error for date if it is missing",
			form: &entities.Booking{
				Name: "Yoga",
				Date: time.Time{},
			},
			expected: []string{"invalid date format (expected YYYY-MM-DD)"},
		},
		{
			name: "should add validation error for name and date if both are missing",
			form: &entities.Booking{
				Name: "",
				Date: time.Time{},
			},
			expected: []string{
				"name is required",
				"invalid date format (expected YYYY-MM-DD)",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := bc.Validate(tt.form)
			if tt.expected == nil {
				assert.Empty(t, errs)
			} else {
				var actual []string
				for _, e := range errs {
					actual = append(actual, e.Error())
				}
				assert.Equal(t, tt.expected, actual)
			}
		})
	}
}

func TestBookingsComponent_CreateBooking(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name      string
		booking   *entities.Booking
		mockRepo  *MockBookingRepository
		want      *entities.Booking
		wantErr   bool
		expectErr string
	}{
		{
			name: "should successfully create a booking",
			booking: &entities.Booking{
				Name: "Test booking",
				Date: now,
			},
			mockRepo: &MockBookingRepository{
				CheckClassExistsOnDateFn: func(t time.Time) bool {
					return true
				},
				AddBookingFn: func(b *entities.Booking) (*entities.Booking, error) {
					return b, nil
				},
			},
			want: &entities.Booking{
				Name: "Test booking",
				Date: now,
			},
		},
		{
			name: "should return an error when class does not exist",
			booking: &entities.Booking{
				Name: "Test booking",
				Date: now,
			},
			mockRepo: &MockBookingRepository{
				CheckClassExistsOnDateFn: func(t time.Time) bool {
					return false
				},
			},
			wantErr:   true,
			expectErr: "no class exists on this date",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &BookingsComponent{
				tt.mockRepo,
			}
			got, err := bc.CreateBooking(tt.booking)

			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
