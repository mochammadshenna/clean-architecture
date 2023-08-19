package repository

import (
	"context"

	"github.com/aeramu/mongolib"
	"github.com/mochammadshenna/clean-architecture/entity"
)

type repository struct {
	db *mongolib.Database
}

func (r *repository) FindBookByID(ctx context.Context, bookID string) (*entity.Book, error) {
	var book Book

	err := r.db.Coll("book").Query().
		Equal("id", bookID).
		FindOne(ctx).Consume(&book)
	if err != nil {
		return nil, err
	}

	return book.Entity(), nil
}

func (r *repository) InsertBook(ctx context.Context, book entity.Book) error {
	model := Book{
		ID:         book.ID,
		Title:      book.Title,
		Author:     book.Author,
		CoverImage: book.CoverImage,
	}
	err := r.db.Coll("book").Save(ctx, mongolib.NewObjectID(), model)
	if err != nil {
		return err
	}

	return nil
}
