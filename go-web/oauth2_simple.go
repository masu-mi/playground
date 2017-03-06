package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/satori/go.uuid"

	"golang.org/x/oauth2"
)

type (
	session struct {
		startAt time.Time
		from    string
	}
	user         *oauth2.Token
	loginService struct {
		Config *oauth2.Config

		mSess    sync.RWMutex
		sessions map[string]*session
		mUser    sync.RWMutex
		users    map[string]user

		CookieName string
		Secure     bool
		Domain     string
		Path       string
		MaxAge     int
	}
)

func NewLoginService(a *oauth2.Config) *loginService {
	return &loginService{
		Config:   a,
		mSess:    sync.RWMutex{},
		sessions: map[string]*session{},
		mUser:    sync.RWMutex{},
		users:    map[string]user{},

		CookieName: "_I",
		Secure:     false,
		Domain:     "localhost",
		Path:       "/",
		MaxAge:     600,
	}
}

func (s *loginService) RegisterSession(path string) string {
	state := newState()
	s.mSess.Lock()
	s.sessions[state] = &session{
		startAt: time.Now(),
		from:    path,
	}
	s.mSess.Unlock()
	return state
}
func (s *loginService) GetSession(state string) (*session, error) {
	s.mSess.RLock()
	sess, ok := s.sessions[state]
	s.mSess.RUnlock()
	if !ok {
		return nil, errors.New("no session in sessions")
	} else if time.Now().After(sess.startAt.Add(1 * time.Minute)) {
		return nil, errors.New("time out")
	}
	return sess, nil
}
func (s *loginService) AddUser(id, session string, t *oauth2.Token) {
	s.mSess.Lock()
	delete(s.sessions, session)
	s.mSess.Unlock()

	s.mUser.Lock()
	s.users[id] = user(t)
	s.mUser.Unlock()
}

func (s *loginService) loginHandler(w http.ResponseWriter, r *http.Request) {
	url := s.Config.AuthCodeURL(
		s.RegisterSession(r.URL.Query().Get("from")),
	)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func newState() string {
	return uuid.NewV4().String()
}

func (s *loginService) dispatchCallbackHandler(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")
	sess, err := s.GetSession(state)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	code := r.URL.Query().Get("code")
	token, err := s.Config.Exchange(context.TODO(), code)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := newState()
	s.AddUser(id, state, token)

	http.SetCookie(w, &http.Cookie{
		Name: s.CookieName, Value: id,
		Secure: s.Secure,
		Domain: s.Domain,
		Path:   s.Path,
		MaxAge: s.MaxAge,
	})
	path := sess.from
	if path == "" {
		path = "/"
	}
	http.Redirect(w, r, path, http.StatusTemporaryRedirect)
}

func (s *loginService) loginCheck(prefix string) func(func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(h func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			url := prefix + "?from=" + r.URL.Path
			if c, err := r.Cookie("_I"); err != nil {
				http.Redirect(w, r, url, http.StatusTemporaryRedirect)
				return
			} else if _, ok := s.users[c.Value]; !ok {
				http.Redirect(w, r, url, http.StatusTemporaryRedirect)
				return
			}
			h(w, r)
		}
	}
}

func setLoginMechanism(m *http.ServeMux, prefix string, s *loginService) func(func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	m.HandleFunc(prefix, s.loginHandler)
	m.HandleFunc(prefix+"/callback", s.dispatchCallbackHandler)
	return s.loginCheck(prefix)
}

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
	safen := setLoginMechanism(m, "/auth/github", NewLoginService(conf))

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
