package controllers

import (
	"encoding/json"
	"errors"
	"github.com/Vidyuallatha/glofox/src/components"
	"github.com/Vidyuallatha/glofox/src/utils"
	"log"
	"net/http"
)

type BookingsController struct {
	Component components.BookingsComponent
}

var bookingsComponent = components.InitBookingsComponent()

func HandleBookings(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		controller := BookingsController{}
		controller.CreateBooking(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (bc *BookingsController) CreateBooking(w http.ResponseWriter, r *http.Request) {
	bookingForm := bookingsComponent.GetBookingForm()
	if err := json.NewDecoder(r.Body).Decode(bookingForm); err != nil {
		log.Println("Error occurred while decoding json ", err)
		utils.WriteJSON(w, http.StatusBadRequest, nil, []error{errors.New("invalid request body")})
		return
	}

	if err := bookingsComponent.Validate(bookingForm); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, nil, err)
		return
	}

	booking, err := bookingsComponent.CreateBooking(bookingForm)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, nil, []error{err})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, booking, nil)
}
