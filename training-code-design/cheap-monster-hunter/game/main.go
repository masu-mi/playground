package main

import (
	"log"
	"net/http"
	"os"

	app "github.com/masu-mi/playground/training-code-design/cheap-monster-hunter/game/adapter/controller/http"
	"github.com/masu-mi/playground/training-code-design/cheap-monster-hunter/game/adapter/gateway/dummy"
	"github.com/masu-mi/playground/training-code-design/cheap-monster-hunter/game/usecase"
)

func init() {
	usecase.Logger = log.New(os.Stdout, "[chep-monster-hunter:http]: ", log.Lshortfile)
}

func main() {
	handler := app.NewHTTPHandler(&dummy.Gateway{
		HunterRepo:  &dummy.HunterRepo{},
		MonsterRepo: &dummy.MonsterRepo{},
		EventLogger: usecase.NewEventBus(&eventLogger{}),
	})
	http.ListenAndServe(":8080", handler)
}

type eventLogger struct{}

func (el *eventLogger) Receive(e usecase.Event) {
	if _, ok := e.(*usecase.EventAttack); ok {
		log.Printf("[Issued Event] Summary: %s", e.Summary())
	}
}
