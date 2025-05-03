package entities

import (
	"errors"
	"time"
)

type Class struct {
	ClassName string    `json:"class_name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Capacity  int       `json:"capacity"`
}

type ClassRepository interface {
	AddClass(c *Class) (*Class, error)
	CheckClassExists(start, end time.Time) bool
}

type ClassEntity struct {
	ClassRepository
}

// In-memory storage of classes
var Classes []Class

func (e ClassEntity) AddClass(c *Class) (*Class, error) {
	for _, existing := range Classes {
		if (c.StartDate.Before(existing.EndDate) && c.EndDate.After(existing.StartDate)) ||
			c.StartDate.Equal(existing.StartDate) || c.EndDate.Equal(existing.EndDate) {
			return nil, errors.New("another class already exists in that date range")
		}
	}

	Classes = append(Classes, *c)
	return c, nil
}

func (e ClassEntity) CheckClassExists(start, end time.Time) bool {
	for _, c := range Classes {
		if (start.Before(c.EndDate) && end.After(c.StartDate)) ||
			start.Equal(c.StartDate) || end.Equal(c.EndDate) {
			return true
		}
	}
	return false
}
