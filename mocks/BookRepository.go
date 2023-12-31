// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import entity "github.com/mochammadshenna/clean-architecture/entity"
import mock "github.com/stretchr/testify/mock"

// BookRepository is an autogenerated mock type for the BookRepository type
type BookRepository struct {
	mock.Mock
}

// FindBookByID provides a mock function with given fields: ctx, bookID
func (_m *BookRepository) FindBookByID(ctx context.Context, bookID string) (*entity.Book, error) {
	ret := _m.Called(ctx, bookID)

	var r0 *entity.Book
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Book); ok {
		r0 = rf(ctx, bookID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, bookID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertBook provides a mock function with given fields: ctx, book
func (_m *BookRepository) InsertBook(ctx context.Context, book entity.Book) error {
	ret := _m.Called(ctx, book)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Book) error); ok {
		r0 = rf(ctx, book)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
