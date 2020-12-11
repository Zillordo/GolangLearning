package main

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("sessions/bcrypt/templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signUp)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	user := getUser(w, r)
	_ = tpl.ExecuteTemplate(w, "index.gohtml", user)
}

func bar(w http.ResponseWriter, r *http.Request) {
	user := getUser(w, r)
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	_ = tpl.ExecuteTemplate(w, "bar.gohtml", user)
}

func signUp(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	var usr user

	if r.Method == http.MethodPost {
		userName := r.FormValue("username")
		pass := r.FormValue("password")
		firstName := r.FormValue("firstname")
		lastName := r.FormValue("lastname")

		if _, ok := dbUsers[userName]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		sId, err := uuid.NewV4()
		if err != nil {
			log.Fatal(err)
		}
		cookie := &http.Cookie{
			Name:  "session",
			Value: sId.String(),
		}
		http.SetCookie(w, cookie)
		dbSessions[cookie.Value] = userName

		bs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		usr = user{userName, bs, firstName, lastName}
		dbUsers[userName] = usr

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	_ = tpl.ExecuteTemplate(w, "signup.gohtml", usr)
}
