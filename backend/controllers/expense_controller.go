// package controllers

// import (
// 	"encoding/json"
// 	"expense-tracker/config"
// 	"expense-tracker/models"
// 	"expense-tracker/utils"
// 	"net/http"
// )

// func GetExpenses(w http.ResponseWriter, r *http.Request) {
// 	rows, err := config.DB.Query("SELECT id, name, amount, category, date FROM expenses")
// 	if err != nil {
// 		utils.RespondWithError(w, http.StatusInternalServerError, "Could not fetch expenses")
// 		return
// 	}
// 	defer rows.Close()

// 	expenses := []models.Expense{}
// 	for rows.Next() {
// 		var expense models.Expense
// 		if err := rows.Scan(&expense.ID, &expense.Name, &expense.Amount, &expense.Category, &expense.Date); err != nil {
// 			utils.RespondWithError(w, http.StatusInternalServerError, "Error scanning expense")
// 			return
// 		}
// 		expenses = append(expenses, expense)
// 	}

// 	utils.RespondWithJSON(w, http.StatusOK, expenses)
// }

// func AddExpense(w http.ResponseWriter, r *http.Request) {
// 	var expense models.Expense
// 	if err := json.NewDecoder(r.Body).Decode(&expense); err != nil {
// 		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}

// 	query := "INSERT INTO expenses (name, amount, category, date) VALUES (?, ?, ?, ?)"
// 	_, err := config.DB.Exec(query, expense.Name, expense.Amount, expense.Category, expense.Date)
// 	if err != nil {
// 		utils.RespondWithError(w, http.StatusInternalServerError, "Could not add expense")
// 		return
// 	}

// 	utils.RespondWithJSON(w, http.StatusCreated, "Expense added successfully")
// }

package controllers

import (
	// "database/sql"
	"encoding/json"
	"expense-tracker/config"
	"expense-tracker/models"
	"net/http"
)

// GetExpenses fetches all expenses
// ///change for commet | below is working fine///
func GetExpenses(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, name, amount, category, date FROM expenses"
	rows, err := config.DB.Query(query)
	if err != nil {
		http.Error(w, `{"error": "Could not fetch expenses"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var expenses []models.Expense

	for rows.Next() {
		var expense models.Expense
		err := rows.Scan(&expense.ID, &expense.Name, &expense.Amount, &expense.Category, &expense.Date)
		if err != nil {
			http.Error(w, `{"error": "Error scanning expenses"}`, http.StatusInternalServerError)
			return
		}
		expenses = append(expenses, expense)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expenses)
}

/////change for commet | above is working fine///

// func GetExpenses(w http.ResponseWriter, r *http.Request) {
// 	query := "SELECT id, name, amount, category, date FROM expenses"
// 	rows, err := config.DB.Query(query)
// 	if err != nil {
// 		log.Println("Error executing query:", err) // Log the error
// 		http.Error(w, `{"error": "Could not fetch expenses"}`, http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()

// 	var expenses []models.Expense

// 	for rows.Next() {
// 		var expense models.Expense
// 		err := rows.Scan(&expense.ID, &expense.Name, &expense.Amount, &expense.Category, &expense.Date)
// 		if err != nil {
// 			log.Println("Error scanning row:", err) // Log scanning error
// 			http.Error(w, `{"error": "Error scanning expenses"}`, http.StatusInternalServerError)
// 			return
// 		}
// 		expenses = append(expenses, expense)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(expenses)
// }

// AddExpense adds a new expense

// ///change for commet | below is working fine///
func AddExpense(w http.ResponseWriter, r *http.Request) {
	var expense models.Expense
	if err := json.NewDecoder(r.Body).Decode(&expense); err != nil {
		http.Error(w, `{"error": "Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	query := "INSERT INTO expenses (name, amount, category, date) VALUES (?, ?, ?, ?)"
	_, err := config.DB.Exec(query, expense.Name, expense.Amount, expense.Category, expense.Date)
	if err != nil {
		http.Error(w, `{"error": "Could not add expense"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Expense added successfully"}`))
}

/////change for commet | above is working fine///
