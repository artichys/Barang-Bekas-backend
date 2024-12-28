package routes

import (
	"Barang-bekas-backend/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// User routes
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	// Item routes
	router.HandleFunc("/item", handlers.GetItems).Methods("GET")
	router.HandleFunc("/item/{id}", handlers.GetItemByID).Methods("GET")
	router.HandleFunc("/item", handlers.CreateItem).Methods("POST")
	router.HandleFunc("/item/{id}", handlers.UpdateItem).Methods("PUT")
	router.HandleFunc("/item/{id}", handlers.DeleteItem).Methods("DELETE")

	// Offer routes
	router.HandleFunc("/offers", handlers.GetOffers).Methods("GET")
	router.HandleFunc("/offers/{id}", handlers.GetOfferByID).Methods("GET")
	router.HandleFunc("/offers", handlers.CreateOffer).Methods("POST")
	router.HandleFunc("/offers/{id}", handlers.UpdateOffer).Methods("PUT")
	router.HandleFunc("/offers/{id}", handlers.DeleteOffer).Methods("DELETE")

	// Transaction routes
	router.HandleFunc("/transactions", handlers.GetTransactions).Methods("GET")
	router.HandleFunc("/transactions/{id}", handlers.GetTransactionByID).Methods("GET")
	router.HandleFunc("/transactions", handlers.CreateTransaction).Methods("POST")
	router.HandleFunc("/transactions/{id}", handlers.UpdateTransaction).Methods("PUT")
	router.HandleFunc("/transactions/{id}", handlers.DeleteTransaction).Methods("DELETE")

	// Review routes
	router.HandleFunc("/reviews", handlers.GetReviews).Methods("GET")
	router.HandleFunc("/reviews/{id}", handlers.GetReviewByID).Methods("GET")
	router.HandleFunc("/reviews", handlers.CreateReview).Methods("POST")
	router.HandleFunc("/reviews/{id}", handlers.UpdateReview).Methods("PUT")
	router.HandleFunc("/reviews/{id}", handlers.DeleteReview).Methods("DELETE")

	return router
}
