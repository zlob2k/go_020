package web

// internal/ui/web

import (
	"fmt"
	"net/http"

	blog "example.com/zlob2k/go_020/internal/core/blog"
	model "example.com/zlob2k/go_020/internal/model"
)

/*
2. Получение оригинальной ссылки
GET /links/{short_code}
Response

	{ "url": "https://example.com/some/very/long/url",
	  "visits": 15	}

При каждом запросе необходимо увеличивать счётчик visits.
*/
type GetLinksIdHandler struct {
	Usecase GetLinksIdUseCase
	Dbparam model.Tdbparam
}

func NewGetLinksIdHandler(usecase GetLinksIdUseCase, dbparam model.Tdbparam) *GetLinksIdHandler {
	fmt.Printf("\nNewGetLinksIdHandler()")
	return &GetLinksIdHandler{
		Usecase: usecase,
		Dbparam: dbparam,
	}
}

func (h *GetLinksIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 1.Разобрать GET-запрос
	// 2.Прочитать запись из БД
	// 3.UPDATE DATABASE with new Visits
	// 4.Сгенерировать ответ
	// 5.Возвратить ответ

	const op = "GetLinksIdHandler.ServeHTTP():"
	fmt.Printf("\n%s", op)

	//short_code_0209, err_0209 := ParseRequestGet02(w, r) // было внутри Run()
	path, err_0209 := ParseRequestGet022(w, r)
	if err_0209 != nil {
		fmt.Printf("\n%s Error:%s", op, err_0209)
	} else {
		//fmt.Printf("\n%s short_code_0209=%d", op, short_code_0209)
		fmt.Printf("\n%s path=%s", op, path)
	}

	answer, err := h.Usecase.Run(h.Dbparam, path)
	switch err {
	case nil:
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", answer) //to browser side
		//	SendAnswerHTTP(answer, w)
	case blog.ErrNoteNotFound:
		w.WriteHeader(http.StatusNotFound)
		// SendAnswer(err, w)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		// SendAnswer(err, w)
	}

	/*
		        // habr.com version
				path := path.Base(req.URL.Path)
				renderedPost, err := h.usecase.Run(path)
				switch err {
				case nil:
					res.WriteHeader(http.StatusOK)
					h.template.Render(res, "view_post.html", h.toViewModel(renderedPost))
				case blog.ErrPostNotFound:
					res.WriteHeader(http.StatusNotFound)
					h.template.Render(res, "404.html", nil)
				default:
					res.WriteHeader(http.StatusInternalServerError)
					h.template.Render(res, "500.html", nil)
				}
	*/
}

func ParseRequestGet022(wr http.ResponseWriter, req *http.Request) (path string, err error) {
	//(path string, short_code int, err error)
	const op = "ParseRequestGet022():"
	fmt.Printf("\n%s", op)
	if req.Method == http.MethodPost {
		// Сюда не заходит вообще
		// Обработка POST-запроса
		//fmt.Fprintf(wr, "\n%s POST-запрос", op) //to browser side
		//fmt.Printf("\n%s POST-запрос", op)      // on server side

	} else if req.Method == http.MethodGet {
		// Обработка GET-запроса  http://127.0.0.1:443/links/29
		path = req.PathValue("id")
		//short_code, err = strconv.Atoi(path)
		////fmt.Fprintf(wr, "\n%s GET-запрос", op) //to browser side
		////fmt.Printf("\n%s GET-запрос, id=%s", op, id) // on server side
		if path == "" {
			err = fmt.Errorf("\n%s empty path in GET request", op)
		}
	}
	//return path, short_code, err
	return path, err
}
