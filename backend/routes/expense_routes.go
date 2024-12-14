// package routes

// import (
// 	"expense-tracker/controllers"
// 	"github.com/gorilla/mux"
// )

// func SetupRoutes() *mux.Router {
// 	router := mux.NewRouter()

// 	router.HandleFunc("/api/expenses", controllers.GetExpenses).Methods("GET")
// 	router.HandleFunc("/api/expenses", controllers.AddExpense).Methods("POST")


// 	return router
// }


package routes

import (
	"expense-tracker/controllers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// GET route for fetching expenses
	router.HandleFunc("/api/expenses", controllers.GetExpenses).Methods("GET")

	// POST route for adding a new expense
	router.HandleFunc("/api/expenses", controllers.AddExpense).Methods("POST")

	return router
}
