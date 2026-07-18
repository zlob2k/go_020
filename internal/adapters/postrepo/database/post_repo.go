package database

/*
import (
	"fmt"

	model "example.com/zlob2k/go_020/internal/model"
)

//internal/adapters/postrepo/database

type PostRepo struct {
	Dbconn model.Tdbconn
	//c.DbPsw, c.DbUsr, c.DbAddr, c.DbType
}

func NewPostRepo(dbconn model.Tdbconn) *PostRepo {
	//fmt.Printf("\nNewPostRepo() DbAddr=%s", '-')
	fmt.Printf("\nNewPostRepo() DbAddr=%s", dbconn.DBaddr)
	return &PostRepo{Dbconn: dbconn}
}

func (r *PostRepo) GetPostByPath(path string) error {
	return nil
}
*/
///////////////////////////////////////////////////////////
/*
func (r *PostRepo) GetPostByPath(path string) (blog.Post, error) {
	content, err := ioutil.ReadFile(filepath.Join(r.BasePath, path+".md"))

	if err != nil {
		return blog.Post{}, blog.ErrPostNotFound
	}

	post, err := ParseFileContent(string(content))
	post.Path = path

	return post, err
}

func (r *PostRepo) GetAllPosts() ([]blog.Post, error) {
	posts := []blog.Post{}
	entries, err := os.ReadDir(r.BasePath)

	if err != nil {
		return posts, err
	}

	for _, entry := range entries {
		posts = r.maybeLoadPostFromFile(posts, entry)
	}

	return r.sortPostsByTimeDesc(posts), err
}

func (r *PostRepo) maybeLoadPostFromFile(posts []blog.Post, entry fs.DirEntry) []blog.Post {
	if !strings.HasSuffix(entry.Name(), ".md") {
		return posts
	}

	fileName := strings.TrimSuffix(entry.Name(), ".md")
	post, err := r.GetPostByPath(fileName)

	if err != nil {
		log.Printf("WARNING: error loading post \"%s\": %v", fileName, err)

		return posts
	}

	return append(posts, post)
}

func (r *PostRepo) sortPostsByTimeDesc(posts []blog.Post) []blog.Post {
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Time.After(posts[j].Time)
	})

	return posts
}
*/
