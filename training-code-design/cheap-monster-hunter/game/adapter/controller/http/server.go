package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/masu-mi/playground/training-code-design/cheap-monster-hunter/game/adapter/gateway"
	"github.com/masu-mi/playground/training-code-design/cheap-monster-hunter/game/domain"
	"github.com/masu-mi/playground/training-code-design/cheap-monster-hunter/game/domain/service"
)

type applicationHandler struct {
	*mux.Router
	gw  gateway.TransactionalGateway
	eng *service.Engine
	// Logger is hook point of internal logs.
	Logger *log.Logger
}

// NewHTTPHandler retunrs http.Handler to usecase.
func NewHTTPHandler(gw gateway.TransactionalGateway) *applicationHandler {
	r := &applicationHandler{
		Router: mux.NewRouter().StrictSlash(true),
		gw:     gw,
		eng: &service.Engine{
			HunterRepository:  gw.HunterRepository(),
			MonsterRepository: gw.MonsterRepository(),
		},
	}
	r.Router.HandleFunc("/attack/{hunter_id}/{monster_id}", r.attackByIDs)
	return r
}

func (ah *applicationHandler) attackByIDs(w http.ResponseWriter, r *http.Request) {
	ah.logf("[START] attackByIDs()")
	defer ah.logf("[END] attackByIDs()")
	ctx, cancel := context.WithTimeout(r.Context(), 1000*time.Millisecond)
	defer cancel()
	// do Commit/Abort
	ctx, commit, abort := ah.gw.ContextWithTx(ctx)
	defer commit()

	vars := mux.Vars(r)
	var hunter *domain.Hunter
	{
		hunterID, err := getUUID(vars, "hunter_id")
		if err != nil {
			reportFailToParse(w, "hunter_id", vars["hunter_id"])
			return
		}
		hunter, err = ah.eng.HunterRepository.FindByID(ctx, hunterID)
		if err != nil {
			errNotFound := &domain.ErrNotFound{}
			if errors.As(err, &errNotFound) {
				reportNotFoundError(w, "hunter", errNotFound)
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
			abort()
			return
		}
	}

	var monster *domain.Monster
	{
		monsterID, err := getUUID(vars, "monster_id")
		if err != nil {
			reportFailToParse(w, "monster_id", vars["monster_id"])
			return
		}
		monster, err = ah.eng.MonsterRepository.FindByID(ctx, monsterID)
		if err != nil {
			errNotFound := &domain.ErrNotFound{}
			if errors.As(err, &errNotFound) {
				reportNotFoundError(w, "monster", errNotFound)
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
			abort()
			return
		}
	}

	profit, err := ah.eng.AttackByHunterWithContext(ctx, hunter, monster)
	if err != nil {
		errNotFound := &domain.ErrNotFound{}
		if errors.As(err, &errNotFound) {
			reportNotFoundError(w, "entity", errNotFound)
		} else {
			reportFailToParse(w, "key", "value")
		}
		abort()
		return
	}
	reportProfit(w, profit)
}

func (ah *applicationHandler) SetLogger(l *log.Logger) {
	ah.Logger = l
}

func (ah *applicationHandler) logf(format string, v ...interface{}) {
	if ah.Logger == nil {
		return
	}
	ah.Logger.Printf(format, v...)
}

func reportProfit(w http.ResponseWriter, profit []domain.Material) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]domain.Material{"item": profit})
}

func reportFailToParse(w http.ResponseWriter, k, v string) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(newErrFailToParse(k, v))
}

func reportNotFoundError(w http.ResponseWriter, tipe string, e *domain.ErrNotFound) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newErrNotFound(tipe, e.ID))
}

func reportCanceled(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
}

func newErrFailToParse(name string, value string) map[string]string {
	return map[string]string{"error": "fail to parse", "name": name, "value": value}
}

func newErrNotFound(tipe string, id fmt.Stringer) map[string]string {
	return map[string]string{"error": "not found", "type": tipe, "id": id.String()}
}

func getUUID(m map[string]string, key string) (uuid.UUID, error) {
	v := m[key]
	return uuid.Parse(v)
}
