package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type SuccessType struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

type JobType struct {
	Title       string `json:"title"`
	Salary      string `json:"salary"`
	Type        string `json:"type"`
	Location    string `json:"location"`
	Description string `json:"description"`
}

func SendSuccessRes(w http.ResponseWriter, statusCode int, data interface{}) {
	resp := SuccessType{
		Status: true,
		Data:   data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(resp)
}

func jobService(w http.ResponseWriter, r *http.Request) {

	jobs := []JobType{
		{
			Title:       "Full-Stack Developer (5 Years Experience)",
			Salary:      "60,000 - 70,000 BDT (per month)",
			Type:        "Full Time - (On-site)",
			Description: "We are looking for a highly skilled and experienced Full-Stack Developer...",
		},
		{
			Title:       "Frontend Developer (3 Years Experience)",
			Salary:      "40,000 - 50,000 BDT (per month)",
			Type:        "Part Time - (Remote)",
			Description: "Seeking a Frontend Developer with ReactJS experience...",
		},
	}
	SendSuccessRes(w, http.StatusOK, jobs)
}

func main() {

	router := http.NewServeMux()

	router.HandleFunc("GET /v1/api/jobs", jobService)

	log.Println("SERVER START SUCCESS:8080")

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatalln("SERVER START ERROR:", err)
	}

}
