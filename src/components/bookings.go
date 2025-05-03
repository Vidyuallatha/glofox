package components

import (
	"errors"
	"github.com/Vidyuallatha/glofox/src/entities"
)

type BookingsComponent struct {
	entities.BookingRepository
}

func InitBookingsComponent() *BookingsComponent {
	return &BookingsComponent{
		&entities.BookingEntity{},
	}
}

func (bc *BookingsComponent) GetBookingForm() *entities.Booking {
	return new(entities.Booking)
}

func (bc *BookingsComponent) Validate(form *entities.Booking) []error {
	var errs []error
	if form.Name == "" {
		errs = append(errs, errors.New("name is required"))
	}
	if form.Date.IsZero() {
		errs = append(errs, errors.New("invalid date format (expected YYYY-MM-DD)"))
	}
	return errs
}

func (bc *BookingsComponent) CreateBooking(booking *entities.Booking) (*entities.Booking, error) {
	if !bc.BookingRepository.CheckClassExistsOnDate(booking.Date) {
		return nil, errors.New("no class exists on this date")
	}
	return bc.BookingRepository.AddBooking(booking)
}
