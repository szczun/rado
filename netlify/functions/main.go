package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Counter struct {
	Count int `json:"count"`
}

func main() {
	http.HandleFunc(
		"/.netlify/functions/main",
		func(w http.ResponseWriter, r *http.Request) {
			filePath := "./counter.json"

			// Odczyt pliku JSON
			if r.Method == http.MethodGet {
				data, err := os.ReadFile(filePath)
				if err != nil {
					http.Error(
						w,
						"Error reading file",
						http.StatusInternalServerError,
					)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(data)
			}

			// Zaktualizowanie licznika
			if r.Method == http.MethodPost {
				data, err := os.ReadFile(filePath)
				if err != nil {
					http.Error(
						w,
						"Error reading file",
						http.StatusInternalServerError,
					)
					return
				}

				var counter Counter
				err = json.Unmarshal(data, &counter)
				if err != nil {
					http.Error(
						w,
						"Error unmarshalling data",
						http.StatusInternalServerError,
					)
					return
				}

				// Zwiększenie licznika
				counter.Count++

				// Zapisanie zaktualizowanego pliku
				newData, err := json.MarshalIndent(counter, "", "  ")
				if err != nil {
					http.Error(
						w,
						"Error marshalling data",
						http.StatusInternalServerError,
					)
					return
				}

				err = os.WriteFile(filePath, newData, 0644)
				if err != nil {
					http.Error(
						w,
						"Error writing file",
						http.StatusInternalServerError,
					)
					return
				}

				// Zwrócenie zaktualizowanej wartości
				w.Header().Set("Content-Type", "application/json")
				w.Write(newData)
			}
		},
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	fmt.Println("Starting server on port " + port)
	http.ListenAndServe(":"+port, nil)
}
