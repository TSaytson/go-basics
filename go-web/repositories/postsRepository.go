package repositories

import (
	"database/sql"
	"encoding/json"
	Db "go-web/database"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

var db *sql.DB = Db.ConnectDatabase()

func GetPosts() []byte {

	posts := []Post{}

	rows, err := db.Query("SELECT * FROM posts;")
	checkErr(err)

	for rows.Next() {
		post := Post{}
		rows.Scan(&post.Id, &post.Title, &post.Body)
		posts = append(posts, post)
	}
	rows.Close()
	res, err := json.Marshal(posts)
	checkErr(err)
	return res
}

func CreatePost() {
	statement, err := db.Prepare(`INSERT INTO posts
		(title, body) values (?, ?);`)
	checkErr(err)
	_, err = statement.Exec("This is a Post", "This is the body")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
