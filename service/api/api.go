package api

import (
	"errors"
	"strings"

	"github.com/mochammadshenna/clean-architecture/entity"
)

type GetBookRequest struct {
	BookID string
}

type GetBookResponse struct {
	Book *entity.Book
}

type CreateBookRequest struct {
	Title      string
	Author     string
	CoverImage string
}

type CreateBookResponse struct {
	Message string
}

func (req GetBookRequest) Validate() error {
	if req.BookID == "" {
		return errors.New("Book ID is required")
	}
	return nil
}

func (req CreateBookRequest) Validate() error {
	if req.Title == "" {
		return errors.New("Title is required")
	}
	if req.Author == "" {
		return errors.New("Author is required")
	}
	if req.CoverImage == "" {
		return errors.New("Cover Image is required")
	}
	if strings.Contains(req.CoverImage, " ") {
		return errors.New("Cover Image cannot contain spaces")
	}

	return nil
}
