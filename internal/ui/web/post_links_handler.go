package web

// internal/ui/web

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	model "example.com/zlob2k/go_020/internal/model"
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
	Dbparam model.Tdbparam
}

func NewPostLinksHandler(usecase PostLinksUseCase, dbparam model.Tdbparam) *PostLinksHandler {
	fmt.Printf("\nNewPostLinksHandler()")
	return &PostLinksHandler{
		Usecase: usecase,
		Dbparam: dbparam,
	}
}

func (h *PostLinksHandler) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	// 0.Parse POST-Request
	// 1.Generate new SHORT_CODE
	// 2.Update database with new URL
	// 3.Generate JSON answer
	// 4.Return answer

	const op = "PostLinksHandler.ServeHTTP():"
	fmt.Printf("\n%s", op)

	/////////////////////////////////
	// 0.Parse POST-Request
	url_obj, err_0210 := ParseRequestPost01(wr, req)
	if err_0210 != nil {
		fmt.Printf("\n%s Error:%s", op, err_0210)
	} else {
		fmt.Printf("\n%s Url:%s", op, url_obj.Url)
	}

	/////////////////////////////////
	// 1.Generate new SHORT_CODE
	// 2.Update database with new URL
	// 3.Generate JSON answer
	// 4.Return answer
	answer, err := h.Usecase.Run(h.Dbparam, url_obj)
	//example answer = {"short_code": "abc123"}

	/////////////////////////////////
	// 5.Send answer
	switch err {
	case nil:
		wr.WriteHeader(http.StatusOK)
		fmt.Fprintf(wr, "%s", answer) //to browser side
		//	SendAnswerHTTP(answer, w)
	default:
		wr.WriteHeader(http.StatusInternalServerError)
		// SendAnswer(err, w)
	}
}

func ParseRequestPost01(wr http.ResponseWriter, req *http.Request) (url model.TUrlObj, err error) {
	const op = "ParseRequestPost01():"
	fmt.Printf("\n%s", op)

	if req.Method == http.MethodPost {
		// Обработка POST-запроса
		//fmt.Fprintf(wr, "\n%s POST-запрос", op) //to browser side
		//fmt.Printf("\n%s POST-запрос", op)      // on server side
		if req.Body != nil {
			// Тело запроса есть, значит, это POST
			body05, err05 := io.ReadAll(req.Body)
			if err05 != nil {
				//fmt.Fprintf(wr, "\nPOST.Body reading error: %v", err05)
				err = err05
			} else {
				BodyStr05 := string(body05)

				// Здесь можно распарсить данные из тела
				reader05 := strings.NewReader(BodyStr05)
				json_decoder05 := json.NewDecoder(reader05)
				err051 := json_decoder05.Decode(&url)
				if err051 != nil {
					//fmt.Fprintf(wr, "\nError in json_decoder05.Decode: %v", err051)
					err = err051
				}
				/*
					fmt.Fprintf(w, "\nUrl: %s", InputType01.Url)
					//Url: rbc.ru/lk
				*/
			}
		}
	} else if req.Method == http.MethodGet {
		// Обработка GET-запроса  http://127.0.0.1:443/links/29
		//fmt.Fprintf(wr, "\n%s GET-запрос", op) //to browser side
		//fmt.Printf("\n%s GET-запрос", op)      // on server side
	}
	return url, err
}
