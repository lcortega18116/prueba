package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/lcortega18116/prueba/internal/models"
)

type itemPayload struct {
	Ticker     string `json:"ticker"`
	TargetFrom string `json:"target_from"`
	TargetTo   string `json:"target_to"`
	Company    string `json:"company"`
	Action     string `json:"action"`
	Brokerage  string `json:"brokerage"`
	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`
	Time       string `json:"time"`
}

func ItemsRoutes(db *sqlx.DB) http.Handler {
	r := chi.NewRouter()
	r.Get("/", listItems(db))
	r.Post("/", createItem(db))
	r.Get("/{id}", getItem(db))
	return r
}

func listItems(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var items []models.Item
		if err := db.SelectContext(ctx, &items, `SELECT * FROM items ORDER BY id DESC`); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(items)
	}
}

func createItem(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		defer r.Body.Close()

		var p itemPayload
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var id int64
		err := db.QueryRowContext(ctx, `INSERT INTO items (ticker, target_from, target_to, company, action, brokerage, rating_from, rating_to, time) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id`, p.Ticker, p.TargetFrom, p.TargetTo, p.Company, p.Action, p.Brokerage, p.RatingFrom, p.RatingTo, p.Time).Scan(&id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(map[string]any{"id": id})
	}
}

func getItem(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := chi.URLParam(r, "id")
		var it models.Item
		err := db.GetContext(ctx, &it, `SELECT * FROM items WHERE id = $1`, id)
		if err == sql.ErrNoRows {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(it)
	}
}
