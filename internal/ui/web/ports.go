package web

//import (
//blog "example.com/zlob2k/go_020/internal/core/blog"
//)

type UseCases struct {
	PostLinks  PostLinksUseCase
	GetLinksId GetLinksIdUseCase
}

type PostLinksUseCase interface {
	Run(db_addr string) error
}

type GetLinksIdUseCase interface {
	Run(db_addr string) error
}
