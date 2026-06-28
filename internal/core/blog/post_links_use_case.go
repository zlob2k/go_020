package blog

import (
	"fmt"

	database "example.com/zlob2k/go_020/internal/adapters/postrepo/database"
)

type PostLinksUseCase struct {
	postRepo database.PostRepo
}

func NewPostLinksUseCase(postRepo database.PostRepo) *PostLinksUseCase {
	fmt.Printf("\nNewPostLinksUseCase()")
	return &PostLinksUseCase{postRepo: postRepo}
}

func (u *PostLinksUseCase) Run(db_addr string) error {
	//post, err := u.postRepo.GetPostByPath(path)

	//if err != nil {
	//	return RenderedPost{}, err
	//}

	//return u.renderPost(post)
	return nil
}

//func (u *ViewPostUseCase) renderPost(post Post) (RenderedPost, error) {
