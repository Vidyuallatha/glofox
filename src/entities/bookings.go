package entities

import "time"

type Booking struct {
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

// In-memory storage of bookings
var Bookings []Booking

type BookingRepository interface {
	AddBooking(b *Booking) (*Booking, error)
	CheckClassExistsOnDate(date time.Time) bool
}

type BookingEntity struct {
	BookingRepository
}

func (e *BookingEntity) AddBooking(b *Booking) (*Booking, error) {
	Bookings = append(Bookings, *b)
	return b, nil
}

func (e *BookingEntity) CheckClassExistsOnDate(date time.Time) bool {
	for _, c := range Classes {
		// Check if the date falls within the range of start and end date (inclusive)
		if date.Equal(c.StartDate) || date.Equal(c.EndDate) || (date.After(c.StartDate) && date.Before(c.EndDate)) {
			return true
		}
	}
	return false
}
