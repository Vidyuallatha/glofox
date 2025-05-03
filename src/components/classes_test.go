package components

import (
	"errors"
	"testing"
	"time"

	"github.com/Vidyuallatha/glofox/src/entities"
	"github.com/stretchr/testify/assert"
)

type MockClassRepository struct {
	CheckClassExistsFn func(start, end time.Time) bool
	AddClassFn         func(class *entities.Class) (*entities.Class, error)
}

func (m *MockClassRepository) CheckClassExists(start, end time.Time) bool {
	if m.CheckClassExistsFn != nil {
		return m.CheckClassExistsFn(start, end)
	}
	return false
}

func (m *MockClassRepository) AddClass(class *entities.Class) (*entities.Class, error) {
	if m.AddClassFn != nil {
		return m.AddClassFn(class)
	}
	return nil, errors.New("not implemented")
}

func TestClassComponent_Valid(t *testing.T) {
	cc := &ClassesComponent{}

	now := time.Now()
	tests := []struct {
		name     string
		form     *entities.Class
		expected []string
	}{
		{
			name: "should not have any validation errors if form is valid",
			form: &entities.Class{
				ClassName: "Yoga",
				StartDate: now,
				EndDate:   now.AddDate(0, 0, 7),
				Capacity:  10,
			},
			expected: nil,
		},
		{
			name: "should have validation errors if all the fields in form are missing",
			form: &entities.Class{},
			expected: []string{
				"class name is required",
				"invalid start date format (expected YYYY-MM-DD)",
				"invalid end date format (expected YYYY-MM-DD)",
				"capacity is required",
			},
		},
		{
			name: "should add validation error if class name is missing",
			form: &entities.Class{
				StartDate: now,
				EndDate:   now.AddDate(0, 0, 1),
				Capacity:  5,
			},
			expected: []string{"class name is required"},
		},
		{
			name: "should add validation error if capacity is zero",
			form: &entities.Class{
				ClassName: "Yoga",
				StartDate: now,
				EndDate:   now.AddDate(0, 0, 1),
				Capacity:  0,
			},
			expected: []string{"capacity is required"},
		},
		{
			name: "should add validation error if start and end dates are missing",
			form: &entities.Class{
				ClassName: "Yoga",
				Capacity:  5,
			},
			expected: []string{
				"invalid start date format (expected YYYY-MM-DD)",
				"invalid end date format (expected YYYY-MM-DD)",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := cc.Validate(tt.form)

			if tt.expected == nil {
				assert.Empty(t, errs)
			} else {
				var actual []string
				for _, err := range errs {
					actual = append(actual, err.Error())
				}
				assert.ElementsMatch(t, tt.expected, actual)
			}
		})
	}
}

func TestClassComponent_CreateClass(t *testing.T) {
	now := time.Now()
	nextWeek := now.AddDate(0, 0, 7)

	tests := []struct {
		name        string
		classInput  *entities.Class
		mockRepo    *MockClassRepository
		expected    *entities.Class
		expectErr   bool
		expectedErr string
	}{
		{
			name: "should create class for valid form data",
			classInput: &entities.Class{
				ClassName: "Pilates",
				StartDate: now,
				EndDate:   nextWeek,
				Capacity:  15,
			},
			mockRepo: &MockClassRepository{
				CheckClassExistsFn: func(start, end time.Time) bool {
					return false
				},
				AddClassFn: func(class *entities.Class) (*entities.Class, error) {
					return class, nil
				},
			},
			expected: &entities.Class{
				ClassName: "Pilates",
				StartDate: now,
				EndDate:   nextWeek,
				Capacity:  15,
			},
		},
		{
			name: "should throw error if start date is after end data",
			classInput: &entities.Class{
				ClassName: "HIIT",
				StartDate: nextWeek,
				EndDate:   now,
				Capacity:  20,
			},
			mockRepo:    &MockClassRepository{},
			expectErr:   true,
			expectedErr: "start and end dates are invalid",
		},
		{
			name: "should throw error if there is already a class in given date range",
			classInput: &entities.Class{
				ClassName: "Zumba",
				StartDate: now,
				EndDate:   nextWeek,
				Capacity:  10,
			},
			mockRepo: &MockClassRepository{
				CheckClassExistsFn: func(start, end time.Time) bool {
					return true
				},
			},
			expectErr:   true,
			expectedErr: "another class exists in this date range",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &ClassesComponent{
				tt.mockRepo,
			}
			got, err := cs.CreateClass(tt.classInput)

			if tt.expectErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, got)
			}
		})
	}
}
