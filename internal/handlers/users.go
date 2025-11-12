package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/lcortega18116/prueba/internal/models"
)

type userPayload struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

func UsersRoutes(db *sqlx.DB) http.Handler {
	r := chi.NewRouter()
	r.Get("/", listUsers(db))
	r.Post("/", createUser(db))
	r.Get("/{id}", getUser(db))
	return r
}

func listUsers(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []models.User
		if err := db.Select(&users, `SELECT id, email, full_name FROM users ORDER BY id DESC`); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(users)
	}
}

func createUser(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p userPayload
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var id int64
		// CONFLICT → CockroachDB permite UPSERT
		err := db.QueryRow(`INSERT INTO users (email, full_name) VALUES ($1, $2) RETURNING id`, p.Email, p.FullName).Scan(&id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(map[string]any{"id": id})
	}
}

func getUser(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		var u models.User
		err := db.Get(&u, `SELECT id, email, full_name FROM users WHERE id = $1`, id)
		if err == sql.ErrNoRows {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(u)
	}
}
