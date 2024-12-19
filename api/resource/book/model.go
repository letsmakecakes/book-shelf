package book

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Book struct {
	ID            uuid.UUID `gorm:"primarykey"`
	Title         string
	Author        string
	PublishedDate time.Time
	ImageURL      string
	Description   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type Books []*Book

type DTO struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedDate string `json:"published_date"`
	ImageURL      string `json:"image_url"`
	Description   string `json:"description"`
}

type Form struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedDate string `json:"published_date"`
	ImageURL      string `json:"image_url"`
	Description   string `json:"description"`
}

// ToModel Form to Book model
func (f *Form) ToModel() *Book {
	pubDate, _ := time.Parse("2006-01-02", f.PublishedDate)

	return &Book{
		Title:         f.Title,
		Author:        f.Author,
		PublishedDate: pubDate,
		ImageURL:      f.ImageURL,
		Description:   f.Description,
	}
}

// ToDto Book to DTO to show to the end user
func (b *Book) ToDto() *DTO {
	return &DTO{
		ID:            b.ID.String(),
		Title:         b.Title,
		Author:        b.Author,
		PublishedDate: b.PublishedDate.Format("2006-01-02"),
		ImageURL:      b.ImageURL,
		Description:   b.Description,
	}
}

func (bs Books) ToDto() []*DTO {
	dos := make([]*DTO, len(bs))
	for i, v := range bs {
		dos[i] = v.ToDto()
	}

	return dos
}
