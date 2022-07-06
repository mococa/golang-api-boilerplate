package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mococa/golang-api-boilerplate/server/dtos"
	"github.com/mococa/golang-api-boilerplate/server/services"
	uuid "github.com/satori/go.uuid"
)

/* -------------- Types -------------- */
type BookController struct {
	services.BookService
}

/* -------------- Initialization -------------- */
func NewBookController(s *services.BookService) BookController {
	return BookController{*s}
}

/* -------------- Controllers -------------- */
func (controller *BookController) CreateBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var book dtos.CreateBookDTO

	err := json.NewDecoder(req.Body).Decode(&book)
	if err != nil {
		fmt.Println(err)

		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(`{"error": "Error unmarshalling"}`))
		return
	}

	created_book, err := controller.CreateBookService(&book)

	if err != nil {
		fmt.Println(err)

		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	response, err := json.Marshal(created_book)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling"}`))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write([]byte(response))

}

func (controller *BookController) BorrowBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var book dtos.BorrowBookDTO

	err := json.NewDecoder(req.Body).Decode(&book)
	if err != nil {
		fmt.Println(err)

		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(`{"error": "Error unmarshalling"}`))
		return
	}

	params := mux.Vars(req)

	book_id := params["bookId"]

	book.BookID = uuid.FromStringOrNil(book_id)

	borrowed_book, err := controller.BorrowBookService(&book)

	if err != nil {
		fmt.Println(err)

		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	response, err := json.Marshal(borrowed_book)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling"}`))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write([]byte(response))

}

func (controller *BookController) BorrowedBooks(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	book_id := uuid.FromStringOrNil(params["userId"])

	borrowed_books, err := controller.GetAllBorrowedBooksService(&dtos.BorrowedBooksDTO{
		UserID: book_id,
	})

	if err != nil {
		fmt.Println(err)

		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	response, err := json.Marshal(borrowed_books)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling"}`))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write([]byte(response))

}
