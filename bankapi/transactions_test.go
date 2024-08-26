package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TEST DEPOSITS
func TestDepositHandler(t *testing.T) {
	// Create a mock HTTP request
	req, err := http.NewRequest("GET", "/deposit?number=0017286376&amount=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the deposit handler function
	deposit(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}
}

func TestDepositHandlerMissingAccountNumber(t *testing.T) {
	// Create a mock HTTP request without the account number query parameter
	req, err := http.NewRequest("GET", "/deposit?amount=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the deposit handler function
	deposit(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := "Account number is missing!"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}

func TestDepositHandlerInvalidAccountNumber(t *testing.T) {
	// Create a mock HTTP request with an invalid account number query parameter
	req, err := http.NewRequest("GET", "/deposit?number=abc&amount=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the deposit handler function
	deposit(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := "Invalid account number!"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}

func TestDepositHandlerInvalidAmount(t *testing.T) {
	// Create a mock HTTP request with an invalid amount query parameter
	req, err := http.NewRequest("GET", "/deposit?number=0017286376&amount=xyz", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the deposit handler function
	deposit(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := "Invalid amount number!"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}

func TestDepositHandlerAccountNotFound(t *testing.T) {
	// Create a mock HTTP request with a non-existent account number query parameter
	req, err := http.NewRequest("GET", "/deposit?number=9999&amount=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the deposit handler function
	deposit(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := "Error getting account: account not found"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}

func TestDepositHandlerNegativeAmount(t *testing.T) {
	// Create a mock HTTP request with a negative amount query parameter
	req, err := http.NewRequest("GET", "/deposit?number=0017286376&amount=-10", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the deposit handler function
	deposit(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := "the amount to deposit should be greater than zero"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}

// TEST STATEMENTS
func TestStatementHandler(t *testing.T) {
	// Create a mock HTTP request
	req, err := http.NewRequest("GET", "/statement?number=0017286376", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the statement handler function
	statement(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}
}

func TestStatementHandlerMissingAccountNumber(t *testing.T) {
	// Create a mock HTTP request without the account number query parameter
	req, err := http.NewRequest("GET", "/statement", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the statement handler function
	statement(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := "Account number is missing!"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}

func TestStatementHandlerInvalidAccountNumber(t *testing.T) {
	// Create a mock HTTP request with an invalid account number query parameter
	req, err := http.NewRequest("GET", "/statement?number=abc", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the statement handler function
	statement(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := "Invalid account number!"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}

func TestStatementHandlerAccountNotFound(t *testing.T) {
	// Create a mock HTTP request with a non-existent account number query parameter
	req, err := http.NewRequest("GET", "/statement?number=9999", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the statement handler function
	statement(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := "Error getting account: account not found"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}

// TEST WITHDRAWALS
func TestWithdrawHandler(t *testing.T) {
	// Create a mock HTTP request
	req, err := http.NewRequest("GET", "/withdraw?number=1001&amount=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the withdraw handler function
	withdraw(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}
}

func TestWithdrawHandlerMissingAccountNumber(t *testing.T) {
	// Create a mock HTTP request without the account number query parameter
	req, err := http.NewRequest("GET", "/withdraw?amount=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the withdraw handler function
	withdraw(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := "Account number is missing!"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}

func TestWithdrawHandlerInvalidAccountNumber(t *testing.T) {
	// Create a mock HTTP request with an invalid account number query parameter
	req, err := http.NewRequest("GET", "/withdraw?number=abc&amount=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the withdraw handler function
	withdraw(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := "Invalid account number!"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}

func TestWithdrawHandlerInvalidAmount(t *testing.T) {
	// Create a mock HTTP request with an invalid amount query parameter
	req, err := http.NewRequest("GET", "/withdraw?number=0017286376&amount=xyz", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the withdraw handler function
	withdraw(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := "Invalid amount number!"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}

func TestWithdrawHandlerAccountNotFound(t *testing.T) {
	// Create a mock HTTP request with a non-existent account number query parameter
	req, err := http.NewRequest("GET", "/withdraw?number=9999&amount=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the withdraw handler function
	withdraw(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := "Error getting account: account not found"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}

// TEST TRANSFERS
func TestTransferHandler(t *testing.T) {
	// Create a mock HTTP request
	req, err := http.NewRequest("GET", "/transfer?from=0018989351&to=0017286376&amount=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the transfer handler function
	transfer(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}
}

func TestTransferHandlerMissingFromAccountNumber(t *testing.T) {
	// Create a mock HTTP request without the from account number query parameter
	req, err := http.NewRequest("GET", "/transfer?to=0017286376&amount=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the transfer handler function
	transfer(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := "You need two account numbers to complete a transfer!"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}

func TestTransferHandlerMissingToAccountNumber(t *testing.T) {
	// Create a mock HTTP request without the to account number query parameter
	req, err := http.NewRequest("GET", "/transfer?from=0018989351&amount=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the transfer handler function
	transfer(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := "You need two account numbers to complete a transfer!"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}

func TestTransferHandlerMissingAmount(t *testing.T) {
	// Create a mock HTTP request without the amount query parameter
	req, err := http.NewRequest("GET", "/transfer?from=0018989351&to=0017286376", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the transfer handler function
	transfer(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := "Amount is invalid!"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}

func TestTransferHandlerInvalidFromAccountNumber(t *testing.T) {
	// Create a mock HTTP request with an invalid from account number query parameter
	req, err := http.NewRequest("GET", "/transfer?from=abc&to=0017286376&amount=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the transfer handler function
	transfer(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := "Invalid debiting account number!"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}

func TestTransferHandlerInvalidToAccountNumber(t *testing.T) {
	// Create a mock HTTP request with an invalid to account number query parameter
	req, err := http.NewRequest("GET", "/transfer?from=0018989351&to=abc&amount=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the transfer handler function
	transfer(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := "Invalid receiving account number!"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}

func TestTransferHandlerInvalidAmount(t *testing.T) {
	// Create a mock HTTP request with an invalid amount query parameter
	req, err := http.NewRequest("GET", "/transfer?from=0018989351&to=0017286376&amount=xyz", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the transfer handler function
	transfer(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := "Amount is invalid!"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}

func TestTransferHandlerFromAccountNotFound(t *testing.T) {
	// Create a mock HTTP request with a non-existent from account number query parameter
	req, err := http.NewRequest("GET", "/transfer?from=9999&to=0017286376&amount=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the transfer handler function
	transfer(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := "Error getting Debit account: account not found"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}

func TestTransferHandlerToAccountNotFound(t *testing.T) {
	// Create a mock HTTP request with a non-existent to account number query parameter
	req, err := http.NewRequest("GET", "/transfer?from=0018989351&to=9999&amount=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the transfer handler function
	transfer(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := "Error getting Receiving account: account not found"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}

func TestTransferHandlerNegativeAmount(t *testing.T) {
	// Create a mock HTTP request with a negative amount query parameter
	req, err := http.NewRequest("GET", "/transfer?from=0018989351&to=0017286376&amount=-10", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the transfer handler function
	transfer(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := "the amount to transfer should be greater than zero"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}

func TestTransferHandlerInsufficientBalance(t *testing.T) {
	// Create a mock HTTP request with an amount greater than the from account balance
	req, err := http.NewRequest("GET", "/transfer?from=0018989351&to=0017286376&amount=1000000", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the transfer handler function
	transfer(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := "insufficient balance to transfer"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}
