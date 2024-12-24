package router

import (
	"book-shelf/api/resource/book"
	"book-shelf/api/resource/health"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func New(db *gorm.DB, v *validator.Validate) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", health.Read)

	r.Route("/v1", func(r chi.Router) {
		bookAPI := book.New(db, v)
		r.Get("/books", bookAPI.List)
		r.Post("/books", bookAPI.Create)
		r.Get("/books/{id}", bookAPI.Read)
		r.Put("/books/{id}", bookAPI.Update)
		r.Delete("/books/{id}", bookAPI.Delete)
	})

	return r
}
