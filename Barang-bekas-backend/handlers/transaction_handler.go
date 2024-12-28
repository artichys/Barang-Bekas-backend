package handlers

import (
	"encoding/json"
	"net/http"
	"database/sql"
	"Barang-bekas-backend/database"
	"Barang-bekas-backend/models"
	"github.com/gorilla/mux"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, offer_id, total_price, transaction_date FROM Transactions")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		if err := rows.Scan(&transaction.ID, &transaction.OfferID, &transaction.TotalPrice, &transaction.TransactionDate); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		transactions = append(transactions, transaction)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

func GetTransactionByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var transaction models.Transaction
	err := database.DB.QueryRow("SELECT id, offer_id, total_price, transaction_date FROM Transactions WHERE id = ?", id).
		Scan(&transaction.ID, &transaction.OfferID, &transaction.TotalPrice, &transaction.TransactionDate)
	if err == sql.ErrNoRows {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec("INSERT INTO Transactions (offer_id, total_price, transaction_date) VALUES (?, ?, ?)",
		transaction.OfferID, transaction.TotalPrice, transaction.TransactionDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var transaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec("UPDATE Transactions SET offer_id = ?, total_price = ?, transaction_date = ? WHERE id = ?",
		transaction.OfferID, transaction.TotalPrice, transaction.TransactionDate, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	_, err := database.DB.Exec("DELETE FROM Transactions WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
