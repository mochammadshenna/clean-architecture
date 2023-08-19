package service

import (
	"context"
	"errors"
	"testing"

	"github.com/mochammadshenna/clean-architecture/entity"
	"github.com/mochammadshenna/clean-architecture/mocks"
	"github.com/mochammadshenna/clean-architecture/service/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	adapter            Adapter
	mockBookRepository *mocks.BookRepository
)

func initTest() {
	mockBookRepository = new(mocks.BookRepository)
	adapter = Adapter{
		BookRepository: mockBookRepository,
	}
}

func Test_service_GetBook(t *testing.T) {
	var (
		ctx    = context.Background()
		bookID = "some-id"
		book   = &entity.Book{
			ID:         bookID,
			Title:      "some title",
			Author:     "some author",
			CoverImage: "some_image.jpg",
		}
	)
	type args struct {
		ctx context.Context
		req api.GetBookRequest
	}
	tests := []struct {
		name    string
		prepare func()
		args    args
		want    *api.GetBookResponse
		wantErr bool
	}{
		{
			name: "invalid argument empty book id",
			prepare: func() {

			},
			args: args{
				ctx: ctx,
				req: api.GetBookRequest{
					BookID: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error database when get book from db",
			prepare: func() {
				mockBookRepository.On("FindBookByID", mock.Anything, bookID).
					Return(nil, errors.New("some error"))
			},
			args: args{
				ctx: ctx,
				req: api.GetBookRequest{
					BookID: bookID,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			prepare: func() {
				mockBookRepository.On("FindBookByID", mock.Anything, bookID).
					Return(book, nil)
			},
			args: args{
				ctx: ctx,
				req: api.GetBookRequest{
					BookID: bookID,
				},
			},
			want:    &api.GetBookResponse{Book: book},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initTest()
			if tt.prepare != nil {
				tt.prepare()
			}
			s := &service{
				adapter: adapter,
			}
			got, err := s.GetBook(tt.args.ctx, tt.args.req)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_service_CreateBook(t *testing.T) {
	var (
		ctx        = context.Background()
		title      = "book title"
		author     = "author name"
		coverImage = "cover_url.jpg"
	)
	type args struct {
		ctx context.Context
		req api.CreateBookRequest
	}
	tests := []struct {
		name    string
		prepare func()
		args    args
		want    *api.CreateBookResponse
		wantErr bool
	}{
		{
			name:    "invalid argument empty title",
			prepare: nil,
			args: args{
				ctx: ctx,
				req: api.CreateBookRequest{
					Title:      "",
					Author:     author,
					CoverImage: coverImage,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid argument empty author",
			prepare: nil,
			args: args{
				ctx: ctx,
				req: api.CreateBookRequest{
					Title:      title,
					Author:     "",
					CoverImage: coverImage,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid argument empty cover image",
			prepare: nil,
			args: args{
				ctx: ctx,
				req: api.CreateBookRequest{
					Title:      title,
					Author:     author,
					CoverImage: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid argument cover image contain space",
			prepare: nil,
			args: args{
				ctx: ctx,
				req: api.CreateBookRequest{
					Title:      title,
					Author:     author,
					CoverImage: "tes .png",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error database",
			prepare: func() {
				mockBookRepository.On("InsertBook", mock.Anything, mock.Anything).
					Return(errors.New("some error"))
			},
			args: args{
				ctx: ctx,
				req: api.CreateBookRequest{
					Title:      title,
					Author:     author,
					CoverImage: coverImage,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			prepare: func() {
				mockBookRepository.On("InsertBook", mock.Anything, mock.MatchedBy(func(b entity.Book) bool {
					assert.Equal(t, title, b.Title)
					assert.Equal(t, author, b.Author)
					assert.Equal(t, coverImage, b.CoverImage)
					return true
				})).Return(nil)
			},
			args: args{
				ctx: ctx,
				req: api.CreateBookRequest{
					Title:      title,
					Author:     author,
					CoverImage: coverImage,
				},
			},
			want:    &api.CreateBookResponse{Message: "Success"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initTest()
			if tt.prepare != nil {
				tt.prepare()
			}
			s := &service{
				adapter: adapter,
			}
			got, err := s.CreateBook(tt.args.ctx, tt.args.req)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
