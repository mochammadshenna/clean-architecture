package service

import (
	"context"

	"github.com/mochammadshenna/clean-architecture/entity"
	"github.com/mochammadshenna/clean-architecture/service/api"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
)

type Service interface {
	CreateBook(ctx context.Context, req api.CreateBookRequest) (*api.CreateBookResponse, error)
	GetBook(ctx context.Context, req api.GetBookRequest) (*api.GetBookResponse, error)
}

type service struct {
	adapter Adapter
}

func NewService(adapter Adapter) Service {
	return &service{
		adapter: adapter,
	}
}

func (s *service) GetBook(ctx context.Context, req api.GetBookRequest) (*api.GetBookResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	book, err := s.adapter.BookRepository.FindBookByID(ctx, req.BookID)
	if err != nil {
		log.WithFields(log.Fields{
			"err":     err,
			"payload": req,
		}).Errorln("[GetBook] An error while get book from db")
		return nil, err
	}

	return &api.GetBookResponse{
		Book: book,
	}, nil
}

func (s *service) CreateBook(ctx context.Context, req api.CreateBookRequest) (*api.CreateBookResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	err := s.adapter.BookRepository.InsertBook(ctx, entity.Book{
		ID:         xid.New().String(),
		Title:      req.Title,
		Author:     req.Author,
		CoverImage: req.CoverImage,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"err":     err,
			"payload": req,
		}).Errorln("[CreateBook] An error while create book to db")
		return nil, err
	}

	return &api.CreateBookResponse{
		Message: "Success",
	}, nil
}
