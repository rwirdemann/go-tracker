// Code generated by mockery v1.0.0
package mocks

import domain "github.com/rwirdemann/3skills.time/domain"
import mock "github.com/stretchr/testify/mock"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Add provides a mock function with given fields: p
func (_m *Repository) AddProject(p domain.Project) {
	_m.Called(p)
}

// AddBooking provides a mock function with given fields: b
func (_m *Repository) AddBooking(b domain.Booking) {
	_m.Called(b)
}

// AllBookings provides a mock function with given fields: projectId
func (_m *Repository) AllBookings(projectId int) []domain.Booking {
	ret := _m.Called(projectId)

	var r0 []domain.Booking
	if rf, ok := ret.Get(0).(func(int) []domain.Booking); ok {
		r0 = rf(projectId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Booking)
		}
	}

	return r0
}

// AllProjects provides a mock function with given fields: filter
func (_m *Repository) AllProjects(filter string) []domain.Project {
	ret := _m.Called(filter)

	var r0 []domain.Project
	if rf, ok := ret.Get(0).(func(string) []domain.Project); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Project)
		}
	}

	return r0
}
