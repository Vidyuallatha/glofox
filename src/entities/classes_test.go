package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClassEntity_AddClass(t *testing.T) {
	entity := ClassEntity{}

	start := time.Date(2025, 5, 10, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 0, 7)

	tests := []struct {
		name        string
		existing    []Class
		newClass    *Class
		expectErr   bool
		expectedErr string
	}{
		{
			name:     "should add class when no conflict exists",
			existing: []Class{},
			newClass: &Class{
				ClassName: "Yoga",
				StartDate: start,
				EndDate:   end,
				Capacity:  10,
			},
			expectErr: false,
		},
		{
			name: "should return error when new class overlaps with existing class",
			existing: []Class{{
				ClassName: "Zumba",
				StartDate: start,
				EndDate:   end,
			}},
			newClass: &Class{
				ClassName: "Pilates",
				StartDate: start.AddDate(0, 0, 3),
				EndDate:   end.AddDate(0, 0, 3),
				Capacity:  8,
			},
			expectErr:   true,
			expectedErr: "another class already exists in that date range",
		},
		{
			name: "should return error when start date exactly matches existing start date",
			existing: []Class{{
				ClassName: "Zumba",
				StartDate: start,
				EndDate:   end,
			}},
			newClass: &Class{
				ClassName: "HIIT",
				StartDate: start,
				EndDate:   end.AddDate(0, 0, 1),
				Capacity:  12,
			},
			expectErr:   true,
			expectedErr: "another class already exists in that date range",
		},
		{
			name: "should return error when end date exactly matches existing end date",
			existing: []Class{{
				ClassName: "Zumba",
				StartDate: start,
				EndDate:   end,
			}},
			newClass: &Class{
				ClassName: "Stretching",
				StartDate: start.AddDate(0, 0, -3),
				EndDate:   end,
				Capacity:  5,
			},
			expectErr:   true,
			expectedErr: "another class already exists in that date range",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Classes = tt.existing

			result, err := entity.AddClass(tt.newClass)

			if tt.expectErr {
				assert.Nil(t, result)
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.newClass, result)
				assert.Contains(t, Classes, *tt.newClass)
			}
		})
	}
}

func TestClassEntity_CheckClassExists(t *testing.T) {
	entity := ClassEntity{}

	start := time.Date(2025, 5, 10, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 0, 7)

	tests := []struct {
		name                 string
		existing             []Class
		checkStart, checkEnd time.Time
		expected             bool
	}{
		{
			name: "should return true if a class exists in given date range",
			existing: []Class{{
				ClassName: "Yoga",
				StartDate: start,
				EndDate:   end,
			}},
			checkStart: start.AddDate(0, 0, 1),
			checkEnd:   end.AddDate(0, 0, -1),
			expected:   true,
		},
		{
			name: "should return true if start date matches existing class",
			existing: []Class{{
				ClassName: "Zumba",
				StartDate: start,
				EndDate:   end,
			}},
			checkStart: start,
			checkEnd:   end.AddDate(0, 0, 2),
			expected:   true,
		},
		{
			name: "should return true if end date matches existing class",
			existing: []Class{{
				ClassName: "Pilates",
				StartDate: start,
				EndDate:   end,
			}},
			checkStart: start.AddDate(0, 0, -2),
			checkEnd:   end,
			expected:   true,
		},
		{
			name: "should return false if no class overlaps with date range",
			existing: []Class{{
				ClassName: "Spin",
				StartDate: start,
				EndDate:   end,
			}},
			checkStart: end.AddDate(0, 0, 1),
			checkEnd:   end.AddDate(0, 0, 3),
			expected:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Classes = tt.existing
			result := entity.CheckClassExists(tt.checkStart, tt.checkEnd)
			assert.Equal(t, tt.expected, result)
		})
	}
}
