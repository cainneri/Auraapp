package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test the accounts API.
func TestAccountsHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/accounts", nil)
	rec := httptest.NewRecorder()

	accountsHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected HTTP 200, got %d", rec.Code)
	}

	var accounts []Account
	if err := json.Unmarshal(rec.Body.Bytes(), &accounts); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}

	if len(accounts) == 0 {
		t.Fatal("expected accounts, received empty response")
	}
}

// Test cluster generation for Kamal Bank.
func TestClustersHandler(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/api/clusters?accountID=kamal",
		nil,
	)
	rec := httptest.NewRecorder()

	clustersHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected HTTP 200, got %d", rec.Code)
	}

	var clusters []Cluster
	if err := json.Unmarshal(rec.Body.Bytes(), &clusters); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}

	if len(clusters) != 258 {
		t.Fatalf("expected 258 clusters, got %d", len(clusters))
	}
}

// Test that the frontend index page is available locally.
func TestStaticFiles(t *testing.T) {
	handler := http.FileServer(http.Dir("./static"))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected HTTP 200, got %d", rec.Code)
	}

	if !strings.HasPrefix(rec.Header().Get("Content-Type"), "text/html") {
		t.Fatalf("expected HTML, got %s", rec.Header().Get("Content-Type"))
	}

	if rec.Body.Len() == 0 {
		t.Fatal("HTML response is empty")
	}
}
