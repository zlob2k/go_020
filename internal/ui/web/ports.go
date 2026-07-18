package web

//internal/ui/web/
import (
	model "example.com/zlob2k/go_020/internal/model"
)

//blog "example.com/zlob2k/go_020/internal/core/blog"

type UseCases struct {
	PostLinks  PostLinksUseCase
	GetLinksId GetLinksIdUseCase
}

type PostLinksUseCase interface {
	Run(dbparam model.Tdbparam, url_obj model.TUrlObj) (string, error)
}

type GetLinksIdUseCase interface {
	Run(dbparam model.Tdbparam, path string) (string, error)
	//Run(dbparam model.Tdbparam, w http.ResponseWriter, r *http.Request) (string, error)
}
