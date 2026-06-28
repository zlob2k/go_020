package blog

import (
	"fmt"

	database "example.com/zlob2k/go_020/internal/adapters/postrepo/database"
)

type GetLinksIdUseCase struct {
	postRepo database.PostRepo
}

func NewGetLinksIdUseCase(postRepo database.PostRepo) *GetLinksIdUseCase {
	fmt.Printf("\nNewGetLinksIdUseCase()")
	return &GetLinksIdUseCase{postRepo: postRepo}
}

func (u *GetLinksIdUseCase) Run(db_addr string) error {
	//post, err := u.postRepo.GetPostByPath(path)
	//if err != nil {
	//	return RenderedPost{}, err
	//}
	//return u.renderPost(post)
	return nil
}

//func (u *ViewPostUseCase) renderPost(post Post) (RenderedPost, error) {}
