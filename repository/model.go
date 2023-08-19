package repository

import "github.com/mochammadshenna/clean-architecture/entity"

type Book struct {
	ID         string `bson:"id"`
	Title      string `bson:"title"`
	Author     string `bson:"author"`
	CoverImage string `bson:"cover_image"`
}

func (b Book) Entity() *entity.Book {
	return &entity.Book{
		ID:         b.ID,
		Title:      b.Title,
		Author:     b.Author,
		CoverImage: b.CoverImage,
	}
}
