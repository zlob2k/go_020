package web

import (
	"fmt"
	"net/http"
)

// internal/ui/web

func NewRouter(usecases *UseCases, db_addr string) http.Handler {
	mux := http.NewServeMux()

	H_PostLinksHandler := NewPostLinksHandler(usecases.PostLinks, db_addr)
	H_GetLinksIdHandler := NewGetLinksIdHandler(usecases.GetLinksId, db_addr)

	http.HandleFunc("/links", H_PostLinksHandler.ServeHTTP)
	http.HandleFunc("/links/{id}", H_GetLinksIdHandler.ServeHTTP)

	fmt.Printf("\nNewRouter()")
	return mux
}
