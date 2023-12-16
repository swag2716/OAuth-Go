package main

import (
	"OAuth-Go/controllers"
	"OAuth-Go/helpers"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestOAuthGoogleLogin(t *testing.T) {
	req, err := http.NewRequest("GET", "/auth/google/login", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.OauthGoogleLogin)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusTemporaryRedirect {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusTemporaryRedirect)
	}
}

func TestOAuthGithubLogin(t *testing.T) {
	req, err := http.NewRequest("GET", "/auth/github/login", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.OauthGithubLogin)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusTemporaryRedirect {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusTemporaryRedirect)
	}
}

func TestOAuthGoogleCallback(t *testing.T) {
	state := helpers.GenerateStateOauthCookie(httptest.NewRecorder())

	values := url.Values{}
	values.Add("code", "valid_code")
	values.Add("state", state)

	req, err := http.NewRequest("GET", "/auth/google/callback?"+values.Encode(), nil)
	if err != nil {
		t.Fatal(err)
	}

	req.AddCookie(&http.Cookie{Name: "oauthstate", Value: state})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.OauthGoogleCallback)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusTemporaryRedirect {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusTemporaryRedirect)
	}
}

func TestOAuthGithubCallback(t *testing.T) {
	state := helpers.GenerateStateOauthCookie(httptest.NewRecorder())

	values := url.Values{}
	values.Add("code", "valid_code")
	values.Add("state", state)

	req, err := http.NewRequest("GET", "/auth/github/callback?"+values.Encode(), nil)
	if err != nil {
		t.Fatal(err)
	}

	req.AddCookie(&http.Cookie{Name: "oauthstate", Value: state})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.OauthGithubCallback)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusTemporaryRedirect {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusTemporaryRedirect)
	}
}
