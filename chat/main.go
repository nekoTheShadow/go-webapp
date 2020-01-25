package main

import (
	"chat/trace"
	"flag"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	addr := flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse()

	gomniauth.SetSecurityKey(SecurityKey)
	gomniauth.WithProviders(
		google.New(GoogleClientId, GoogleClientSecret, "http://localhost:3000/auth/callback/google"),
		facebook.New(FacebookClientId, FacebookClientSecret, "http://localhost:3000/auth/callback/facebook"),
		github.New(GithubClientId, GithubClientSecret, "http://localhost:3000/auth/callback/github"),
	)

	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	go r.run()
	log.Println("Webサーバを開始します。ポート: " + *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}
