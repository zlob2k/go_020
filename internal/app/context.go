package app

//internal/app

import (
	"fmt"
	"net/http"
	"os"

	database "example.com/zlob2k/go_020/internal/adapters/postrepo/database"
	blog "example.com/zlob2k/go_020/internal/core/blog"
	web "example.com/zlob2k/go_020/internal/ui/web"
)

type Context struct {
	DbPsw      string
	DbUsr      string
	DbAddr     string
	DbType     string
	ServerAddr string
}

func NewContext() *Context {
	fmt.Printf("\nNewContext()")
	return &Context{
		DbPsw:      os.Getenv("DB_LINK_PSW"),
		DbUsr:      os.Getenv("DB_LINK_USR"), 
		DbAddr:     "http://localhost:5432",
		DbType:     "postgres",
		ServerAddr: "127.0.0.1:443",
	}
}

func (c *Context) WebServer() *web.Server {
	fmt.Printf("\nWebServer()")
	return web.NewServer(c.ServerAddr, c.Router())
}

func (c *Context) Router() http.Handler {
	fmt.Printf("\nRouter()")
	return web.NewRouter(c.UseCases(), c.DbAddr)
}

func (c *Context) UseCases() *web.UseCases {
	fmt.Printf("\nUseCases()")
	return &web.UseCases{
		PostLinks:  c.PostLinksUseCase(),
		GetLinksId: c.GetLinksIdUseCase(),
	}
}
func (c *Context) PostLinksUseCase() *blog.PostLinksUseCase {
	//return blog.NewPostLinksUseCase(c.PostRepos())
	return nil
}
func (c *Context) GetLinksIdUseCase() *blog.GetLinksIdUseCase {
	//return blog.NewGetLinksIdUseCase(c.PostRepos())
	return nil
}

func (c *Context) PostRepos() *database.PostRepo {
	return database.NewPostRepo(c.DbPsw, c.DbUsr, c.DbAddr, c.DbType)
}
