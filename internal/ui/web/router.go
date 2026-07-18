package web

// internal/ui/web

import (
	"fmt"
	"net/http"

	model "example.com/zlob2k/go_020/internal/model"
)

func NewRouter(usecases *UseCases, dbparam model.Tdbparam) http.Handler {

	const op = "NewRouter()"
	fmt.Printf("\n%s", op)

	mux := http.NewServeMux()

	H_PostLinksHandler := NewPostLinksHandler(usecases.PostLinks, dbparam)
	H_GetLinksIdHandler := NewGetLinksIdHandler(usecases.GetLinksId, dbparam)

	mux.Handle("/links", H_PostLinksHandler)
	mux.Handle("/links/{id}", H_GetLinksIdHandler)

	//mux.Handle("/links", H_PostLinksHandler.ServeHTTP)
	//mux.Handle("/links/{id}", H_GetLinksIdHandler.ServeHTTP)

	return mux

}
