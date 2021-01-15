package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/masu-mi/playground/training-code-design/cheap-monster-hunter/game/domain"
	"github.com/masu-mi/playground/training-code-design/cheap-monster-hunter/game/usecase"
)

type applicationHandler struct {
	*mux.Router
	eng *usecase.Engine
}

// NewHTTPHandler retunrs http.Handler to usecase.
func NewHTTPHandler(eng *usecase.Engine) http.Handler {
	r := &applicationHandler{
		Router: mux.NewRouter().StrictSlash(true),
		eng:    eng,
	}
	r.Router.HandleFunc("/attack/{hunter_id}/{monster_id}", r.attackByIDs)
	return r
}

func (ah *applicationHandler) attackByIDs(w http.ResponseWriter, r *http.Request) {
	defer ioutil.ReadAll(r.Body)

	errNotFound := &domain.ErrNotFound{}
	hunter, monster, err := ah.getAttackInput(w, r)
	if errors.As(err, &errNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// StartTx
	profit, err := ah.eng.AttackByHunter(hunter, monster)
	if err != nil {
		// Abort Tx
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Commit Tx
	// if errTx
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]domain.Material{"item": profit})
}

func (ah *applicationHandler) getAttackInput(w http.ResponseWriter, r *http.Request) (*domain.Hunter, *domain.Monster, error) {
	errNotFound := &domain.ErrNotFound{}
	vars := mux.Vars(r)
	var hunter *domain.Hunter
	{
		hunterID, err := getUUID(vars, "hunter_id")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(newErrFailToParse("hunter_id", vars["hunter_id"]))
			return nil, nil, errors.New("Fail to fetch hunter")
		}
		hunter, err = ah.eng.HunterRepository.FindByID(hunterID)
		if err != nil {
			if errors.As(err, &errNotFound) {
				w.WriteHeader(http.StatusBadRequest)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(newErrNotFound("hunter", hunterID))
				return nil, nil, errors.New("Fail to fetch hunter")
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Server Error")
			return nil, nil, errors.New("error")
		}
	}

	var monster *domain.Monster
	{
		monsterID, err := getUUID(vars, "monster_id")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(newErrFailToParse("monster_id", vars["monster_id"]))
			return nil, nil, errors.New("Fail to fetch monster")
		}
		monster, err = ah.eng.MonsterRepository.FindByID(monsterID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(newErrNotFound("monster", monsterID))
			return nil, nil, errors.New("Fail to fetch monster")
		}
	}
	return hunter, monster, nil
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
