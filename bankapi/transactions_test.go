package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDepositHandler(t *testing.T) {
	// Create a mock HTTP request
	req, err := http.NewRequest("GET", "/deposit?number=1001&amount=10", nil)
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
	expectedBody := "Deposit updated successfully!"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
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
	req, err := http.NewRequest("GET", "/deposit?number=1001&amount=xyz", nil)
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
	expectedBody := "Account with number 9999 can't be found!"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}

func TestDepositHandlerNegativeAmount(t *testing.T) {
	// Create a mock HTTP request with a negative amount query parameter
	req, err := http.NewRequest("GET", "/deposit?number=1001&amount=-10", nil)
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
	expectedBody := "only positive numbers should be allowed to deposit"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}
func TestStatementHandler(t *testing.T) {
	// Create a mock HTTP request
	req, err := http.NewRequest("GET", "/statement?number=1001", nil)
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
	expectedBody := "Statement generated successfully!"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
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
	expectedBody := "Account with number 9999 can't be found!"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}
