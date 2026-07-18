package blog

import (
	"encoding/json"
	"fmt"
	"strconv"

	model "example.com/zlob2k/go_020/internal/model"
)

type PostLinksUseCase struct {
	note_repo NoteRepo
}

func NewPostLinksUseCase(note_repo NoteRepo) *PostLinksUseCase {
	fmt.Printf("\nNewPostLinksUseCase()")
	return &PostLinksUseCase{note_repo: note_repo}
}

func (u *PostLinksUseCase) Run(dbparam model.Tdbparam, url_obj model.TUrlObj) (string, error) {
	// 1.Generate new SHORT_CODE
	// 2.Update database with new URL
	// 3.Generate JSON answer
	// 4.Return answer

	const op = "PostLinksUseCase.Run():"
	fmt.Printf("\n%s", op)

	var err_main error
	var url_obj01 model.TNote
	//var err02 error
	//////////////////////////////
	// 1.Generate new SHORT_CODE
	url_obj01.Short_code, err_main = u.note_repo.GetNewCode()
	url_obj01.Url = url_obj.Url
	//url_obj01.Created_at = time.Now()
	//url_obj01.Visits = 0

	fmt.Printf("\nurl_obj01.Short_code=%d", url_obj01.Short_code)
	fmt.Printf("\nurl_obj01.Url=%s", url_obj01.Url)
	//fmt.Printf("\nurl_obj01.Created_at=%s", url_obj01.Created_at)
	//fmt.Printf("\nurl_obj01.Visits=%d", url_obj01.Visits)

	////////////////////////////
	// 2.Update database with new URL
	if err_main == nil {
		err_main = u.note_repo.InsertLink(url_obj01)
		if err_main != nil {
			fmt.Printf("\n%s Error: %s", op, err_main)
		}
	}

	//////////////////////////////
	// 3.Generate JSON answer
	var json_answer_str string
	type answer_type_0108 struct {
		Short_code string `json:"short_code"`
	}
	if err_main == nil {
		answer := answer_type_0108{Short_code: strconv.Itoa(url_obj01.Short_code)}
		json_data, err_main := json.Marshal(answer)
		if err_main == nil {
			json_answer_str = string(json_data)
			fmt.Printf("\n%s JSON_answer: %s", op, json_answer_str)
		}
	}
	return json_answer_str, err_main
}

//func (u *ViewPostUseCase) renderPost(post Post) (RenderedPost, error) {
