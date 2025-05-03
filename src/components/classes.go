package components

import (
	"errors"
	"github.com/Vidyuallatha/glofox/src/entities"
)

type ClassesComponent struct {
	entities.ClassRepository
}

func InitClassesComponent() *ClassesComponent {
	return &ClassesComponent{
		&entities.ClassEntity{},
	}
}
func (cc *ClassesComponent) GetClassForm() *entities.Class {
	return new(entities.Class)
}

func (cc *ClassesComponent) Validate(form *entities.Class) []error {
	var errs []error
	if form.ClassName == "" {
		errs = append(errs, errors.New("class name is required"))
	}
	if form.StartDate.IsZero() {
		errs = append(errs, errors.New("invalid start date format (expected YYYY-MM-DD)"))
	}
	if form.EndDate.IsZero() {
		errs = append(errs, errors.New("invalid end date format (expected YYYY-MM-DD)"))
	}
	if form.Capacity <= 0 {
		errs = append(errs, errors.New("capacity is required"))
	}
	return errs
}

func (cc *ClassesComponent) CreateClass(class *entities.Class) (*entities.Class, error) {
	if class.StartDate.After(class.EndDate) || class.EndDate.Before(class.StartDate) {
		return nil, errors.New("start and end dates are invalid")
	}

	if cc.CheckClassExists(class.StartDate, class.EndDate) {
		return nil, errors.New("another class exists in this date range")
	}

	return cc.AddClass(class)
}
