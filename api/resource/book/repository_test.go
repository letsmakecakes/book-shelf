package book

import (
	"book-shelf/mock/db"
	"book-shelf/util/test"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestRepository_List(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := db.NewMockDB()
	test.NoError(t, err)

	repo := NewRepository(mockDB)

	mockRows := sqlmock.NewRows([]string{"id", "title", "author"}).AddRow(uuid.New(), "Book_1", "Author_1").AddRow(uuid.New(), "Book_2", "Author_2")

	mock.ExpectQuery("^SELECT (.+) FROM \"books\"").WillReturnRows(mockRows)

	books, err := repo.List()
	test.NoError(t, err)
	test.Equal(t, 2, len(books))
}

func TestRepository_Create(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := db.NewMockDB()
	test.NoError(t, err)

	repo := NewRepository(mockDB)

	id := uuid.New()
	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO \"books\" ").WithArgs(id, "Title", "Author", db.AnyTime{}, "", "", db.AnyTime{}, db.AnyTime{}, nil).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	book := &Book{ID: id, Title: "Title", Author: "Author", PublishedDate: time.Now()}
	_, err = repo.Create(book)
	test.NoError(t, err)
}

func TestRepository_Read(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := db.NewMockDB()
	test.NoError(t, err)

	repo := NewRepository(mockDB)

	id := uuid.New()
	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO \"books\" ").WithArgs(id, "Title", "Author", db.AnyTime{}, "", "", db.AnyTime{}, db.AnyTime{}, nil).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	book := &Book{ID: id, Title: "Title", Author: "Author", PublishedDate: time.Now()}
	_, err = repo.Create(book)
	test.NoError(t, err)
}

func TestRepository_Update(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := db.NewMockDB()
	test.NoError(t, err)

	repo := NewRepository(mockDB)

	id := uuid.New()
	_ = sqlmock.NewRows([]string{"id", "title", "author"}).AddRow(id, "Book1", "Author1")

	mock.ExpectBegin()
	mock.ExpectExec("^UPDATE \"books\" SET").WithArgs("Title", "Author", db.AnyTime{}, "", "", db.AnyTime{}, id).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	book := &Book{ID: id, Title: "Title", Author: "Author"}
	rows, err := repo.Update(book)
	test.NoError(t, err)
	test.Equal(t, 1, rows)
}

func TestRepository_Delete(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := db.NewMockDB()
	test.NoError(t, err)

	repo := NewRepository(mockDB)

	id := uuid.New()
	_ = sqlmock.NewRows([]string{"id", "title", "author"}).AddRow(id, "Book1", "Author1")

	mock.ExpectBegin()
	mock.ExpectExec("^UPDATE \"books\" SET \"deleted_at\"").WithArgs(db.AnyTime{}, id).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	rows, err := repo.Delete(id)
	test.NoError(t, err)
	test.Equal(t, 1, rows)
}
