package web

// internal/ui/web

import (
	"fmt"
	"net/http"
)

/*
   1. Создание короткой ссылки
   POST /links
   Request
   {"url": "https://example.com/some/very/long/url"}
   Response
   {"short_code": "abc123"}
   Сервис должен:
   — сгенерировать short_code
   — сохранить ссылку в базу данных
   — вернуть short_code в ответе
*/

type PostLinksHandler struct {
	Usecase PostLinksUseCase
	Dbconn  string
}

func NewPostLinksHandler(usecase PostLinksUseCase, dbconn string) *PostLinksHandler {
	fmt.Printf("\nNewPostLinksHandler()")
	return &PostLinksHandler{
		Usecase: usecase,
		Dbconn:  dbconn,
	}
}

func (h *PostLinksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//answer, err := h.Usecase.Run(h.Dbconn)
	//if err != nil {
	//	SendAnswer(error, w)
	//} else {
	//	SendAnswerHTTP(answer, w)
	//}
	fmt.Printf("\nPostLinksHandler happened!")
}
