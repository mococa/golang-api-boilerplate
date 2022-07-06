package server

import (
	"fmt"
	"net/http"

	database "github.com/mococa/golang-api-boilerplate/server/database"
	"github.com/mococa/golang-api-boilerplate/server/repositories"
	"github.com/mococa/golang-api-boilerplate/server/services"

	controllers "github.com/mococa/golang-api-boilerplate/server/controllers"

	"github.com/gorilla/mux"
)

var (
	/* -------------- Controllers -------------- */
	users_controller controllers.UserController
	books_controller controllers.BookController

	/* -------------- Services -------------- */
	users_service services.UserService
	books_service services.BookService

	/* -------------- Repositories -------------- */
	users_repository repositories.UserRepository
	books_repository repositories.BookRepository
)

func Init() {
	db := database.Init()

	users_repository = repositories.NewUsersRepo(db)
	books_repository = repositories.NewBooksRepo(db)

	users_service = services.NewUserService(&users_repository)
	books_service = services.NewBookService(&books_repository)

	users_controller = controllers.NewUserController(&users_service)
	books_controller = controllers.NewBookController(&books_service)
}

func SetupRoutes(router *mux.Router) {
	/* -------------- Index -------------- */
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Hello world")
	}).Methods("GET")

	/* -------------- User -------------- */
	router.HandleFunc("/user", users_controller.ListUsers).Methods("GET")
	router.HandleFunc("/user", users_controller.CreateUser).Methods("POST")

	/* -------------- Book -------------- */
	router.HandleFunc("/book", books_controller.CreateBook).Methods("POST")
	router.HandleFunc("/book/{bookId}/borrow", books_controller.BorrowBook).Methods("POST")
	router.HandleFunc("/book/{userId}/borrowed", books_controller.BorrowedBooks).Methods("GET")

}
