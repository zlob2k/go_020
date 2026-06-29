package database

//internal/adapters/postrepo/database

import "fmt"

type PostRepo struct {
	DbAddr string
	//c.DbPsw, c.DbUsr, c.DbAddr, c.DbType
}

func NewPostRepo(DbPsw, DbUsr, DbAddr, DbType string) *PostRepo {
	fmt.Printf("\nNewPostRepo() DbAddr=%s", DbAddr)
	return &PostRepo{DbAddr: DbAddr}
}

