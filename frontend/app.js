
const BACKEND_URL = "http://localhost:8080/api/expenses"; // Replace with your backend URL when deploying to AWS

// Fetch and display expenses when the page loads
window.addEventListener("DOMContentLoaded", loadExpenses);

document.getElementById("expense-form").addEventListener("submit", async function (e) {
  e.preventDefault();

  const description = document.getElementById("description").value;
  const amount = parseFloat(document.getElementById("amount").value);
  const category = document.getElementById("category").value;
  const date = document.getElementById("date").value;

  const newExpense = { name: description, amount, category, date };

  try {
    const response = await fetch(BACKEND_URL, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(newExpense),
    });

    if (!response.ok) {
      throw new Error("Failed to add expense");
    }

    // Clear the form and reload expenses
    document.getElementById("expense-form").reset();
    loadExpenses();
  } catch (error) {
    console.error("Error inserting expense:", error);
    alert("Error inserting expense. Please check your backend.");
  }
});

async function loadExpenses() {
  try {
    const response = await fetch(BACKEND_URL);
    if (!response.ok) {
      throw new Error("Failed to fetch expenses");
    }

    const expenses = await response.json();
    renderTable(expenses);
  } catch (error) {
    console.error("Error loading expenses:", error);
    alert("Error loading expenses. Please check your backend.");
  }
}

function renderTable(expenses) {
  const tableBody = document.getElementById("expense-table");
  tableBody.innerHTML = ""; // Clear existing rows

  expenses.forEach(expense => {
    const row = document.createElement("tr");

    row.innerHTML = `
      <td>${expense.name}</td>
      <td>$${expense.amount.toFixed(2)}</td>
      <td>${expense.category}</td>
      <td>${expense.date}</td>
    `;

    tableBody.appendChild(row);
  });
}
