package database

//internal/adapters/postrepo/database
import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	blog "example.com/zlob2k/go_020/internal/core/blog"
	model "example.com/zlob2k/go_020/internal/model"
	pgx "github.com/jackc/pgx/v5"
)

type NoteRepo struct {
	Dbparam model.Tdbparam
	DBref   *pgx.Conn
}

var ctx = context.Background()

func NewNoteRepo(dbparam model.Tdbparam) *NoteRepo {
	const op = "NewNoteRepo():"
	//fmt.Printf("\n%s DbAddr=%s", op, dbparam.DBaddr)
	//fmt.Printf("\nbparam.DBtype=%s", dbparam.DBtype)
	//fmt.Printf("\nbparam.DBuser=%s", dbparam.DBuser)
	//fmt.Printf("\nbparam.DBpasw=%s", dbparam.DBpasw)
	//fmt.Printf("\nbparam.DBaddr=%s", dbparam.DBaddr)
	//fmt.Printf("\nbparam.DBname=%s", dbparam.DBname)

	//Create DB connection
	//Строка подключения  	//"postgres://user:pasw@localhost:5432/links"
	connStr := dbparam.DBtype + "://" + dbparam.DBuser + ":" + dbparam.DBpasw +
		"@" + dbparam.DBaddr + "/" + dbparam.DBname
	fmt.Printf("\n%s connStr=%s", op, connStr)

	// Set connection
	conn, conn_err := pgx.Connect(ctx, connStr)
	if conn_err != nil {
		fmt.Printf("\n%s ", op)
		log.Fatalf("\nНе удалось подключиться к БД: %v", conn_err)
	} else {
		fmt.Printf("\n%s Connected to DB 'links' ok", op)
	}
	//defer conn.Close(ctx)
	return &NoteRepo{Dbparam: dbparam, DBref: conn}
}

func (r *NoteRepo) Shutdown() {
	r.DBref.Close(ctx)
	fmt.Println("Database connection closed")
}

func (r *NoteRepo) ReadLinkByCode(short_code_in int) (model.TNote, error) {
	const op = "ReadLinkByCode():"
	var note model.TNote
	fmt.Printf("\n%s", op)
	rows := r.DBref.QueryRow(ctx, "SELECT short_code, original_url, created_at, visits FROM t_links2 WHERE short_code = $1", short_code_in)
	err := rows.Scan(&note.Short_code, &note.Url, &note.Created_at, &note.Visits)
	if err != nil {
		//log.Fatalf("Error Db_get_row_by_shcode: %v", err)
		fmt.Printf("\n%s Error:%s", op, err)
		//return model.TNote{}, err
		return model.TNote{}, blog.ErrNoteNotFound
	} else {
		return note, nil
	}
}
func (r *NoteRepo) UpdateLinkByCode(note_obj model.TNote) error {
	const op = "UpdateLink():"
	fmt.Printf("\n%s", op)
	Commandtag, err := r.DBref.Exec(ctx, "UPDATE t_links2 SET original_url = $2, created_at = $3, visits = $4 WHERE short_code = $1", note_obj.Short_code, note_obj.Url, note_obj.Created_at, note_obj.Visits)
	if err != nil {
		fmt.Printf("\n%s Error:%s", op, err)
	} else {
		if Commandtag.RowsAffected() != 0 {
			err = nil
		} else {
			err = fmt.Errorf("\n%s 0 lines changed", op)
			fmt.Printf("\n%s Error:%s", op, err)
		}
	}
	return err
}
func (r *NoteRepo) InsertLink(note_obj model.TNote) error {
	const op = "InsertLink():"
	var code_err error
	fmt.Printf("\n%s", op)
	if note_obj.Short_code == 0 {
		note_obj.Short_code, code_err = r.GetNewCode()
		if code_err != nil {
			fmt.Printf("\n%s Error with short_code ", op)
		}
	}
	note_obj.Created_at = time.Now()
	note_obj.Visits = 0
	Commandtag, err := r.DBref.Exec(ctx, "INSERT INTO t_links2 (short_code, original_url, created_at, visits) VALUES ($1, $2, $3, $4)", note_obj.Short_code, note_obj.Url, note_obj.Created_at, note_obj.Visits)
	if err != nil {
		fmt.Printf("\n%s: Error:%s", op, err)
	} else {
		if Commandtag.RowsAffected() != 0 {
			err = nil
		} else {
			err = fmt.Errorf("\n%s 0 lines changed", op)
		}
	}
	return err
}
func (r *NoteRepo) GetNewCode() (new_code int, err error) {
	//read DB - find MAX short_code from DB
	const op = "GetNewCode():"
	var short_code_max4 int
	fmt.Printf("\n%s", op)
	err = r.DBref.QueryRow(ctx, "SELECT COALESCE(MAX(short_code), 0) AS short_code FROM t_links2").Scan(&short_code_max4)
	if err != nil {
		fmt.Printf("\n%s: Error:%s", op, err)
	}
	new_code = short_code_max4 + 1
	return new_code, err
}

func (r *NoteRepo) ParseRequestGet02(wr http.ResponseWriter, req *http.Request) (short_code int, err error) {
	/*
	   const op = "ParseRequestGet02():"
	   fmt.Printf("\n%s", op)

	   	if req.Method == http.MethodPost {
	   		// Сюда не заходит вообще
	   		// Обработка POST-запроса
	   		//fmt.Fprintf(wr, "\n%s POST-запрос", op) //to browser side
	   		//fmt.Printf("\n%s POST-запрос", op)      // on server side
	   		//err = fmt.Errorf("\n%s POST-запрос", op)
	   		//short_code = -1
	   	} else if req.Method == http.MethodGet {

	   		// Обработка GET-запроса  http://127.0.0.1:443/links/29
	   		id := req.PathValue("id")
	   		short_code, err = strconv.Atoi(id)
	   		//fmt.Fprintf(wr, "\n%s GET-запрос", op) //to browser side
	   		//fmt.Printf("\n%s GET-запрос, id=%s", op, id) // on server side
	   	}

	   return short_code, err
	*/
	return 0, nil
}

func (r *NoteRepo) GenAnswerGet02(note_obj model.TNote) (answer string, err error) {
	// Generate JSON answer
	//	{
	//	  "url": "https://example.com/some/very/long/url",
	//	  "visits": 15
	//	}
	const op = "GenAnswerGet02():"
	fmt.Printf("\n%s", op)

	type answer_type0208 struct {
		Url    string `json:"url"`
		Visits int    `json:"visits"`
	}
	answer0208 := answer_type0208{Url: note_obj.Url, Visits: note_obj.Visits}
	json_data0208, err := json.Marshal(answer0208)
	if err != nil {
		///fmt.Printf("\n%s Ошибка при сериализации 0208: %s", op, err)
		fmt.Printf("\n%s Error:%s", op, err)
	} else {
		//fmt.Fprintf(w, "\n%s", string(json_data0208))
		answer = string(json_data0208)
	}
	//Example:		{"url":"https://dzen.ru/pictures","visits":3}
	return answer, err
}
