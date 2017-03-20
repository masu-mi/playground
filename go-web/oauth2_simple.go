package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/masu-mi/gimmick.git/login"

	"golang.org/x/oauth2"
)

func main() {
	var addr string
	flag.StringVar(&addr, "addr", ":8080", "addr")
	flag.Parse()

	conf := &oauth2.Config{
		RedirectURL: "http://localhost:8080/auth/github/callback",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
		Scopes: []string{"user", "public_repo", "repo"},
	}
	conf.ClientID = os.Getenv("GITHUB_CLIENT_ID")
	conf.ClientSecret = os.Getenv("GITHUB_CLIENT_SECRET")

	m := http.NewServeMux()
	safen := login.AddService(m, "/auth/github", login.NewService(conf))

	m.HandleFunc("/path-1", safen(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello path-1"))
	}))
	m.HandleFunc("/path-2", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello path-2"))
	})
	// サーバー設定
	s := &http.Server{
		Addr:         addr,
		Handler:      m,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	fmt.Printf("%#v\n", s)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
