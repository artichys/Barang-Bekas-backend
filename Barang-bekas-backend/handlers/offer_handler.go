package handlers

import (
	"encoding/json"
	"net/http"
	"database/sql"
	"Barang-bekas-backend/database"
	"Barang-bekas-backend/models"
	"github.com/gorilla/mux"
)

func GetOffers(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, item_id, user_id, offered_price, status FROM Offers")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var offers []models.Offer
	for rows.Next() {
		var offer models.Offer
		if err := rows.Scan(&offer.ID, &offer.ItemID, &offer.UserID, &offer.OfferedPrice, &offer.Status); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		offers = append(offers, offer)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(offers)
}

func GetOfferByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var offer models.Offer
	err := database.DB.QueryRow("SELECT id, item_id, user_id, offered_price, status FROM Offers WHERE id = ?", id).
		Scan(&offer.ID, &offer.ItemID, &offer.UserID, &offer.OfferedPrice, &offer.Status)
	if err == sql.ErrNoRows {
		http.Error(w, "Offer not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(offer)
}

func CreateOffer(w http.ResponseWriter, r *http.Request) {
	var offer models.Offer
	if err := json.NewDecoder(r.Body).Decode(&offer); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec("INSERT INTO Offers (item_id, user_id, offered_price, status) VALUES (?, ?, ?, ?)",
		offer.ItemID, offer.UserID, offer.OfferedPrice, offer.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateOffer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var offer models.Offer
	if err := json.NewDecoder(r.Body).Decode(&offer); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec("UPDATE Offers SET item_id = ?, user_id = ?, offered_price = ?, status = ? WHERE id = ?",
		offer.ItemID, offer.UserID, offer.OfferedPrice, offer.Status, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteOffer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	_, err := database.DB.Exec("DELETE FROM Offers WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
