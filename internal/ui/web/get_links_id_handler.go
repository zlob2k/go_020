package web

// internal/ui/web

import (
	"fmt"
	"net/http"
)

/*
	2. Получение оригинальной ссылки
	GET /links/{short_code}
	Response
	{
	  "url": "https://example.com/some/very/long/url",
	  "visits": 15
	}
	При каждом запросе необходимо увеличивать счётчик visits.
*/

type GetLinksIdHandler struct {
	Usecase GetLinksIdUseCase
	Dbconn  string
}

func NewGetLinksIdHandler(usecase GetLinksIdUseCase, dbconn string) *GetLinksIdHandler {
	fmt.Printf("\nNewGetLinksIdHandler()")
	return &GetLinksIdHandler{
		Usecase: usecase,
		Dbconn:  dbconn,
	}
}

func (h *GetLinksIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//answer, err := h.Usecase.Run(h.Dbconn)
	//if err != nil {
	//	SendAnswer(error, w)
	//} else {
	//	SendAnswerHTTP(answer, w)
	//}
	fmt.Printf("\nGetLinksIdHandler happened!")
}
