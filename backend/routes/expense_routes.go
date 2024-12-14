package routes

import (
	"expense-tracker/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Route for fetching all expenses
	router.HandleFunc("/api/expenses", controllers.GetExpenses).Methods("GET")

	// Route for adding a new expense
	router.HandleFunc("/api/expenses", controllers.AddExpense).Methods("POST")

	return router
}
