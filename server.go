package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"log"

	_ "github.com/mattn/go-sqlite3"
)

type post struct {
	ID    int
	Title string
	URL   string
	User  string
	Score int
}

type login struct {
	Username string
	Password string
}

type register struct {
	login
	Confirm string
}

func main() {
	tmpl := make(map[string]*template.Template)
	tmpl["posts.html"] = template.Must(template.ParseFiles("templates/posts.html", "templates/base.html"))
	tmpl["login.html"] = template.Must(template.ParseFiles("templates/login.html", "templates/base.html"))
	tmpl["register.html"] = template.Must(template.ParseFiles("templates/register.html", "templates/base.html"))

	db, err := sql.Open("sqlite3", "./gorfrit.db")

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY NOT NULL,
			username VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY NOT NULL,
			title VARCHAR(255) NOT NULL,
			url TEXT NOT NULL,
			poster_id INTEGER NOT NULL,
			created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (poster_id) REFERENCES users(id)
		);
		CREATE INDEX IF NOT EXISTS idx_post_title ON posts(title);
	`)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS votes (
			id INTEGER PRIMARY KEY NOT NULL,
			post_id INTEGERO NOT NULL,
			voter_id INTEGER NOT NULL,
			vote INTEGER NOT NULL,
			created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (voter_id) REFERENCES users(id),
			CONSTRAINT unq_post_voter UNIQUE (post_id, voter_id),
			CONSTRAINT chk_vote CHECK (vote = -1 OR vote = 1)
		);
	`)

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/posts/", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`
			SELECT
				posts.id,
				title,
				url,
				users.username,
				(SELECT IFNULL(sum(vote), 0) FROM votes WHERE votes.post_id=posts.id) AS score
			FROM posts
			INNER JOIN users on posts.poster_id=users.id
			ORDER BY posts.created DESC
			LIMIT 50
		`)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		posts := make([]post, 0)

		for rows.Next() {
			var p post
			err = rows.Scan(&p.ID, &p.Title, &p.URL, &p.User, &p.Score)
			if err != nil {
				panic(err)
			}
			posts = append(posts, p)
		}

		err = tmpl["posts.html"].ExecuteTemplate(w, "base", posts)
		if err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/login/", func(w http.ResponseWriter, r *http.Request) {
		lgin := login{Username: "", Password: ""}

		if r.Method == "POST" {
			r.ParseForm()
			log.Printf("Form: %v", r.Form)
		}
		err = tmpl["login.html"].ExecuteTemplate(w, "base", lgin)
		if err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/register/", func(w http.ResponseWriter, r *http.Request) {
		rgstr := register{login{Username: "", Password: ""}, ""}

		if r.Method == "POST" {
			r.ParseForm()
		}

		err = tmpl["register.html"].ExecuteTemplate(w, "base", rgstr)
		if err != nil {
			panic(err)
		}
	})

	fmt.Println("Listening on 8080")
	http.ListenAndServe(":8080", nil)
}
