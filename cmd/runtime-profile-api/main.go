package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"runtime-profile-poc/internal/profile"
)

func main() {
	profilePath := os.Getenv("RUNTIME_PROFILE_PATH")
	if profilePath == "" {
		profilePath = "profiles/runtime-profile.yaml"
	}

	runtimeProfile, err := profile.Load(profilePath)
	if err != nil {
		log.Fatalf("failed to load runtime profile: %v", err)
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{
			"status": "ok",
		})
	})

	http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, runtimeProfile)
	})

	http.HandleFunc("/intent/verdict", func(w http.ResponseWriter, r *http.Request) {
		violations := profile.Validate(runtimeProfile)

		status := "valid_profile"
		if len(violations) > 0 {
			status = "invalid_profile"
		}

		writeJSON(w, map[string]any{
			"application": runtimeProfile.Application.Name,
			"environment": runtimeProfile.Application.Environment,
			"status":      status,
			"violations":  violations,
		})
	})

	log.Println("runtime profile api listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func writeJSON(w http.ResponseWriter, value any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(value); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
