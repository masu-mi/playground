package main

import (
	"log"
	"net/http"
	"os"

	app "github.com/masu-mi/playground/training-code-design/cheap-monster-hunter/game/adapter/controller/http"
	"github.com/masu-mi/playground/training-code-design/cheap-monster-hunter/game/adapter/gateway/dummy"
	"github.com/masu-mi/playground/training-code-design/cheap-monster-hunter/game/domain/service"
)

func init() {
	service.Logger = log.New(os.Stdout, "[chep-monster-hunter:http]: ", log.Lshortfile)
}

func main() {
	handler := app.NewHTTPHandler(&dummy.Gateway{
		HunterRepo:  &dummy.HunterRepo{},
		MonsterRepo: &dummy.MonsterRepo{},
	})
	handler.SetLogger(service.Logger)
	http.ListenAndServe(":8080", handler)
}
