package blog

//internal/core/blog
import (
	"fmt"
	"strconv"

	model "example.com/zlob2k/go_020/internal/model"
)

type GetLinksIdUseCase struct {
	note_repo NoteRepo
}

func NewGetLinksIdUseCase(note_repo NoteRepo) *GetLinksIdUseCase {
	fmt.Printf("\nNewGetLinksIdUseCase()")
	return &GetLinksIdUseCase{note_repo: note_repo}
}

func (u *GetLinksIdUseCase) Run(dbparam model.Tdbparam, path string) (string, error) {
	//func (u *GetLinksIdUseCase) Run(dbparam model.Tdbparam, wr http.ResponseWriter, req *http.Request) (string, error) {
	// 1.Разобрать GET-запрос
	// 2.Прочитать запись из БД
	// 3.UPDATE DATABASE with new Visits
	// 4.Сгенерировать ответ
	// 5.Возвратить ответ

	const op = "GetLinksIdUseCase.Run():"
	fmt.Printf("\n%s", op)

	var err_main error
	var short_code int
	var err01 error
	short_code, err01 = strconv.Atoi(path)
	err_main = err01
	/*
		// 1.Разобрать GET-запрос
		// Перенес во внешний
		short_code, err01 := u.note_repo.ParseRequestGet02(wr, req)
		if err01 != nil {
			fmt.Printf("\n%s %s", op, err01)
		} else {
			fmt.Printf("\n%s short_code=%d", op, short_code)
		}
	*/
	// 2.Прочитать запись из БД
	var note_obj model.TNote
	var err02 error
	if err_main == nil {
		note_obj, err02 = u.note_repo.ReadLinkByCode(short_code)
		if err02 != nil {
			//return "", ErrNoteNotFound
			err_main = ErrNoteNotFound
		}
		fmt.Printf("\nnote.Short_code=%d", note_obj.Short_code)
		fmt.Printf("\nnote.Url=%s", note_obj.Url)
		fmt.Printf("\nnote.Created_at=%s", note_obj.Created_at)
		fmt.Printf("\nnote.Visits=%d", note_obj.Visits)
	}

	//u.note_repo.Closedb()

	// 3 UPDATE DATABASE with new Visits
	var note_obj_02 model.TNote
	if err_main == nil {
		note_obj_02.Short_code = note_obj.Short_code
		note_obj_02.Url = note_obj.Url
		note_obj_02.Created_at = note_obj.Created_at
		note_obj_02.Visits = note_obj.Visits + 1
		err03 := u.note_repo.UpdateLinkByCode(note_obj_02)
		err_main = err03
	}

	// 4.Сгенерировать ответ
	var answer04 string
	var err04 error
	if err_main == nil {
		answer04, err04 = u.note_repo.GenAnswerGet02(note_obj_02)
		err_main = err04
		fmt.Printf("\nAnswer: %s", answer04)
	}

	// 5.Возвратить ответ
	if err_main != nil {
		return "", err_main
	} else {
		return answer04, nil
	}
}

//func (u *ViewPostUseCase) renderPost(post Post) (RenderedPost, error) {}
