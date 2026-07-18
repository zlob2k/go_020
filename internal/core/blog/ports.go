package blog

//internal/core/blog

import (
	"errors"
	"net/http"

	model "example.com/zlob2k/go_020/internal/model"
)

type NoteRepo interface {
	ReadLinkByCode(short_code int) (model.TNote, error)
	UpdateLinkByCode(note_obj model.TNote) error
	InsertLink(note_obj model.TNote) error
	GetNewCode() (new_code int, err error)
	//Closedb()
	ParseRequestGet02(wr http.ResponseWriter, req *http.Request) (short_code int, err error)
	GenAnswerGet02(model.TNote) (string, error)
}

var ErrNoteNotFound = errors.New("Item not found")
