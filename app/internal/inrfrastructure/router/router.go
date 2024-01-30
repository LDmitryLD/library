package router

import (
	"net/http"
	"projects/LDmitryLD/library/app/internal/modules"

	"github.com/go-chi/chi/v5"
)

func NewRouter(controllers modules.Controllers) *chi.Mux {
	r := chi.NewRouter()
	setDefaultRoutes(r)

	r.Post("/library/book", controllers.Library.AddBook)
	r.Post("/library/user", controllers.Library.AddUser)
	r.Get("/library/users", controllers.Library.GetUsers)
	r.Post("/library/author", controllers.Library.AddAuthor)
	r.Put("/library/book/rent", controllers.Library.RentBook)
	r.Put("/library/book/back", controllers.Library.BackBook)
	r.Get("/library/authors/top", controllers.Library.GetTop)

	return r
}

func setDefaultRoutes(r *chi.Mux) {
	r.Get("/swagger", swaggerUI)
	r.Get("/public/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))).ServeHTTP(w, r)
	})
}
