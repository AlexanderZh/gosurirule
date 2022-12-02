package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"gorm.io/gorm"

	"github.com/AlexanderZh/gosurirule/go/database"
	"github.com/AlexanderZh/gosurirule/go/model"
	"github.com/AlexanderZh/gosurirule/go/services"

	"github.com/gorilla/mux"
)

type PgData struct {
	db *gorm.DB
}

func NewPgData(db *gorm.DB) *PgData {
	return &PgData{db: db}
}
func RulesetGetAll(w http.ResponseWriter, r *http.Request) {
	var Ruleset model.Ruleset
	writeDefaultHeaders(&w, "GET, OPTIONS")
	result := database.DB.Order("id").Find(&Ruleset)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Printf("cannot find any ruleset in database. Error: %s", result.Error.Error())		
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Ruleset)
}

func RulesetGetById(w http.ResponseWriter, r *http.Request) {
	var Ruleset model.Ruleset
	vars := mux.Vars(r)
	id, ok := vars["bol"]
	log.Print(id, ok)
	writeDefaultHeaders(&w, "GET, OPTIONS")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Ruleset)
}

//TODO:
func RuleValidate(w http.ResponseWriter, r *http.Request) {
	writeDefaultHeaders(&w, "GET, OPTIONS")
	w.WriteHeader(http.StatusOK)
}

func RuleGetAll(w http.ResponseWriter, r *http.Request) {
	var Rule model.Rule

	writeDefaultHeaders(&w, "GET, OPTIONS")
	result := database.DB.Order("id").Find(&Rule)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Printf("cannot find any rule in database. Error: %s", result.Error.Error())		
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Rule)
}



func RuleGetById(w http.ResponseWriter, r *http.Request) {
	writeDefaultHeaders(&w, "GET, OPTIONS")
	w.WriteHeader(http.StatusOK)
}


func RuleAdd(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	RawRule := &model.RawRule{}
	err := d.Decode(RawRule)
	if(err != nil){
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	Rule,err := services.Parse(RawRule.Raw)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		result := database.DB.Create(&Rule)
		if result.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)

		}
	}
	writeDefaultHeaders(&w, "GET, OPTIONS")
	w.WriteHeader(http.StatusOK)
}


func RuleUpdate(w http.ResponseWriter, r *http.Request) {
	writeDefaultHeaders(&w, "GET, OPTIONS")
	w.WriteHeader(http.StatusOK)
}


func RuleEnable(w http.ResponseWriter, r *http.Request) {
	writeDefaultHeaders(&w, "GET, OPTIONS")
	w.WriteHeader(http.StatusOK)
}


func RuleDisable(w http.ResponseWriter, r *http.Request) {
	writeDefaultHeaders(&w, "GET, OPTIONS")
	w.WriteHeader(http.StatusOK)
}

