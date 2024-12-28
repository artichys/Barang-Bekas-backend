package handlers

import (
	"encoding/json"
	"net/http"
	"database/sql"
	"Barang-bekas-backend/database"
	"Barang-bekas-backend/models"
	"github.com/gorilla/mux"
)

func GetReviews(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, user_id, item_id, rating, comment FROM Reviews")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var review models.Review
		if err := rows.Scan(&review.ID, &review.UserID, &review.ItemID, &review.Rating, &review.Comment); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		reviews = append(reviews, review)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviews)
}

func GetReviewByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var review models.Review
	err := database.DB.QueryRow("SELECT id, user_id, item_id, rating, comment FROM Reviews WHERE id = ?", id).
		Scan(&review.ID, &review.UserID, &review.ItemID, &review.Rating, &review.Comment)
	if err == sql.ErrNoRows {
		http.Error(w, "Review not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(review)
}

func CreateReview(w http.ResponseWriter, r *http.Request) {
	var review models.Review
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec("INSERT INTO Reviews (user_id, item_id, rating, comment) VALUES (?, ?, ?, ?)",
		review.UserID, review.ItemID, review.Rating, review.Comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateReview(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var review models.Review
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec("UPDATE Reviews SET user_id = ?, item_id = ?, rating = ?, comment = ? WHERE id = ?",
		review.UserID, review.ItemID, review.Rating, review.Comment, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteReview(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	_, err := database.DB.Exec("DELETE FROM Reviews WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}