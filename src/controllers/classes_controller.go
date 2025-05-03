package controllers

import (
	"encoding/json"
	"errors"
	"github.com/Vidyuallatha/glofox/src/components"
	"github.com/Vidyuallatha/glofox/src/utils"
	"net/http"
)

type ClassesController struct {
	Component components.ClassesComponent
}

var classesComponent = components.InitClassesComponent()

func HandleClasses(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		controller := ClassesController{}
		controller.CreateClass(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (cc *ClassesController) CreateClass(w http.ResponseWriter, r *http.Request) {
	classForm := classesComponent.GetClassForm()
	if err := json.NewDecoder(r.Body).Decode(classForm); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, nil, []error{errors.New("invalid request body")})
		return
	}

	if err := classesComponent.Validate(classForm); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, nil, err)
		return
	}

	class, err := classesComponent.CreateClass(classForm)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, nil, []error{err})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, class, nil)
}
