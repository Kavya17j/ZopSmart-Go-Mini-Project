package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetMovies(t *testing.T) {
	resetMovies()
	req, err := http.NewRequest("GET", "/movies", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getMovies)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `[{"id":"1","isbn":"438227","title":"3 Idiots","director":{"firstname":"RajKumar","lastname":"Hirani"}},{"id":"2","isbn":"454555","title":"Student Of The Year","director":{"firstname":"Karan","lastname":"Johar"}}]`

	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestGetMovie(t *testing.T) {
	resetMovies()
	req, err := http.NewRequest("GET", "/movies/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getMovie)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"id":"1","isbn":"438227","title":"3 Idiots","director":{"firstname":"RajKumar","lastname":"Hirani"}}`

	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// Add similar tests for createMovie, updateMovie, and deleteMovie functions

func TestCreateMovie(t *testing.T) {
	resetMovies()
	movie := Movie{
		Isbn:     "123456",
		Title:    "New Movie",
		Director: &Director{Firstname: "New", Lastname: "Director"},
	}
	movieJSON, _ := json.Marshal(movie)

	req, err := http.NewRequest("POST", "/movies", bytes.NewBuffer(movieJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createMovie)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var createdMovie Movie
	json.Unmarshal(rr.Body.Bytes(), &createdMovie)

	// Add assertions for the created movie
}

func TestCreateMovie(t *testing.T) {
	resetMovies()
	movie := Movie{
		Isbn:     "123456",
		Title:    "New Movie",
		Director: &Director{Firstname: "New", Lastname: "Director"},
	}
	movieJSON, _ := json.Marshal(movie)

	req, err := http.NewRequest("POST", "/movies", bytes.NewBuffer(movieJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createMovie)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var createdMovie Movie
	json.Unmarshal(rr.Body.Bytes(), &createdMovie)

	if createdMovie.ID == "" {
		t.Errorf("Expected non-empty ID for created movie, got empty ID")
	}

	// Add more assertions based on your requirements
}

func TestUpdateMovie(t *testing.T) {
	resetMovies()
	updatedMovie := Movie{
		ID:       "1",
		Isbn:     "789012",
		Title:    "Updated Movie",
		Director: &Director{Firstname: "Updated", Lastname: "Director"},
	}
	updatedMovieJSON, _ := json.Marshal(updatedMovie)

	req, err := http.NewRequest("PUT", "/movies/1", bytes.NewBuffer(updatedMovieJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(updateMovie)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var updatedResponse Movie
	json.Unmarshal(rr.Body.Bytes(), &updatedResponse)

	if updatedResponse.Title != updatedMovie.Title {
		t.Errorf("Expected title %s, got %s", updatedMovie.Title, updatedResponse.Title)
	}

	// Add more assertions based on your requirements
}

func TestDeleteMovie(t *testing.T) {
	resetMovies()
	req, err := http.NewRequest("DELETE", "/movies/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteMovie)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var updatedMovies []Movie
	json.Unmarshal(rr.Body.Bytes(), &updatedMovies)

	// Add assertions based on the updatedMovies slice after deletion
}


func resetMovies() {
	movies = []Movie{
		{ID: "1", Isbn: "438227", Title: "3 Idiots", Director: &Director{Firstname: "RajKumar", Lastname: "Hirani"}},
		{ID: "2", Isbn: "454555", Title: "Student Of The Year", Director: &Director{Firstname: "Karan", Lastname: "Johar"}},
	}
}
