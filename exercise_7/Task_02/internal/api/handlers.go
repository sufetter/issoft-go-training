package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"webcrm/internal/models"
)

var userRepository *models.UserRepository

func SetMux(port string) error {
	ur, err := models.NewUserRepository(filepath.Join("storage", "users.db"))
	defer ur.Close()

	if err != nil {
		log.Printf("Dbutil error:\n%v", err)
		return err
	}

	userRepository = ur
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", getUsers)
	mux.HandleFunc("GET /users/{id}", getUser)
	mux.HandleFunc("DELETE /users/{id}", deleteUser)

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
	if err != nil {
		log.Printf("Error starting server:\n%v", err)
		return err
	}
	return nil
}

func getUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := userRepository.GetAllUsers()
	if handleError(w, err, "Error fetching users") {
		return
	}
	respondWithJSON(w, http.StatusOK, users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Id provided", http.StatusBadRequest)
		return
	}

	user, err := userRepository.GetUserById(id)
	if handleError(w, err, "Error fetching user") {
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Id provided", http.StatusBadRequest)
		return
	}

	err = userRepository.DeleteUser(id)
	if handleError(w, err, "Error deleting user") {
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "User deleted successfully"})
}

func handleError(w http.ResponseWriter, err error, logMsg string) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("%s: %v", logMsg, err)
		http.Error(w, "Not found", http.StatusNotFound)
	} else {
		log.Printf("%s: %v", logMsg, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return true
}

func respondWithJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
