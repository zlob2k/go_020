package app

//internal/app

import (
	"fmt"
	"net/http"
	"os"

	database "example.com/zlob2k/go_020/internal/adapters/postrepo/database"
	blog "example.com/zlob2k/go_020/internal/core/blog"
	model "example.com/zlob2k/go_020/internal/model"
	web "example.com/zlob2k/go_020/internal/ui/web"
	//app "example.com/zlob2k/go_020/internal/app"
)

type Context struct {
	Dbparam    model.Tdbparam
	ServerAddr string
}

func NewContext() *Context {
	//ToDo:  config.yaml
	fmt.Printf("\nNewContext()")
	Dbparam := model.Tdbparam{
		DBuser: os.Getenv("DB_LINK_USR"), //env.GetString("TEMPLATE_PATH", filepath.Join("web", "template")),
		DBpasw: os.Getenv("DB_LINK_PSW"),
		DBaddr: "localhost:5432",
		DBtype: "postgres",
		DBname: "links",
	}
	return &Context{
		Dbparam:    Dbparam,
		ServerAddr: "127.0.0.1:443",
	}
}

func (c *Context) WebServer() *web.Server {
	fmt.Printf("\nWebServer()")
	return web.NewServer(c.ServerAddr, c.Router())
}

func (c *Context) Router() http.Handler {
	fmt.Printf("\nRouter()")
	return web.NewRouter(c.UseCases(), c.Dbparam)
}
func (c *Context) UseCases() *web.UseCases {
	fmt.Printf("\nUseCases()")
	Repository := c.NoteRepos()
	return &web.UseCases{
		PostLinks:  c.PostLinksUseCase(Repository),
		GetLinksId: c.GetLinksIdUseCase(Repository),
	}
}
func (c *Context) PostLinksUseCase(repos *database.NoteRepo) *blog.PostLinksUseCase {
	//return blog.NewPostLinksUseCase(c.PostRepos())
	return blog.NewPostLinksUseCase(repos)
}
func (c *Context) GetLinksIdUseCase(repos *database.NoteRepo) *blog.GetLinksIdUseCase {
	//return blog.NewGetLinksIdUseCase(c.PostRepos())
	return blog.NewGetLinksIdUseCase(repos)
}
func (c *Context) NoteRepos() *database.NoteRepo {
	return database.NewNoteRepo(c.Dbparam)
}

func (c *Context) Shutdown() {

}

/*
func (c *Context) UseCases() *web.UseCases {
	fmt.Printf("\nUseCases()")
	return &web.UseCases{
		PostLinks:  c.PostLinksUseCase(),
		GetLinksId: c.GetLinksIdUseCase(),
	}
}
func (c *Context) PostLinksUseCase() *blog.PostLinksUseCase {
	//return blog.NewPostLinksUseCase(c.PostRepos())
	return blog.NewPostLinksUseCase(c.NoteRepos())
}
func (c *Context) GetLinksIdUseCase() *blog.GetLinksIdUseCase {
	//return blog.NewGetLinksIdUseCase(c.PostRepos())
	return blog.NewGetLinksIdUseCase(c.NoteRepos())
}
func (c *Context) NoteRepos() *database.NoteRepo {
	return database.NewNoteRepo(c.Dbconn)
}
*/
