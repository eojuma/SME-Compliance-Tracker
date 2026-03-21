package handlers

import (
	"encoding/json"
	"net/http"
)


func HealthCheck(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")

	response:=map[string]string{
		"status":"ok",
		"service":"SME Compliance Tracker",
	}
	json.NewEncoder(w).Encode(response)
}

func APIRoot(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	
	response:=map[string]string{
		"messege":"Welcome to SME Compliance Tracker",
	}
	json.NewEncoder(w).Encode(response)
}