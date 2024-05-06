package main

import (
	"mkassymk/forum/internal/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestPostView(t *testing.T) {
	// Create a new instance of our application struct which uses the mocked // dependencies.
	app := newTestApplication(t)
	// Establish a new test server for running end-to-end tests.
	ts := newTestServer(t, app.routes())
	defer ts.Close()
	// Set up some table-driven tests to check the responses sent by our // application for different URLs.
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{
			name:     "Valid ID",
			urlPath:  "/post/view?id=1",
			wantCode: http.StatusOK,
		},
		{
			name:     "Non-existent ID",
			urlPath:  "/post/view?id=40450594",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Negative ID",
			urlPath:  "/post/view?id=-1",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Decimal ID",
			urlPath:  "/post/view?id=1.23",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "String ID",
			urlPath:  "/post/view?id=foo",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Empty ID",
			urlPath:  "/post/view?id=",
			wantCode: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.urlPath)
			assert.Equal(t, code, tt.wantCode)
			if tt.wantBody != "" {
				assert.StringContains(t, body, tt.wantBody)
			}
		})
	}
}

func TestUserSignup(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	const (
		validName     = "Bob"
		validLastname = "Marley"
		validPassword = "validPa$$word"
		validEmail    = "bob@example.com"
	)

	tests := []struct {
		name          string
		userFirstName string
		userLastName  string
		userEmail     string
		userPassword  string
		wantCode      int
	}{
		{
			name:          "Valid submission",
			userFirstName: validName,
			userLastName:  validLastname,
			userEmail:     validEmail,
			userPassword:  validPassword,
			wantCode:      http.StatusSeeOther,
		},
		{
			name:          "Empty name",
			userFirstName: "",
			userLastName:  validLastname,
			userEmail:     validEmail,
			userPassword:  validPassword,
			wantCode:      http.StatusSeeOther,
		},
		{
			name:          "Empty email",
			userFirstName: validName,
			userLastName:  validLastname,
			userEmail:     "",
			userPassword:  validPassword,
			wantCode:      http.StatusSeeOther,
		},
		{
			name:          "Empty password",
			userFirstName: validName,
			userLastName:  validLastname,
			userEmail:     validEmail,
			userPassword:  "",
			wantCode:      http.StatusSeeOther,
		},
		{
			name:          "Invalid email",
			userFirstName: validName,
			userLastName:  validLastname,
			userEmail:     "bob@example",
			userPassword:  validPassword,
			wantCode:      http.StatusSeeOther,
		},
		{
			name:          "Short password",
			userFirstName: validName,
			userLastName:  validLastname,
			userEmail:     validEmail,
			userPassword:  "pa$$",
			wantCode:      http.StatusSeeOther,
		},
		{
			name:          "Duplicate email",
			userFirstName: validName,
			userLastName:  validLastname,
			userEmail:     "dupe@example.com",
			userPassword:  validPassword,
			wantCode:      http.StatusSeeOther,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{}
			form.Add("firstname", tt.userFirstName)
			form.Add("lastname", tt.userLastName)
			form.Add("email", tt.userEmail)
			form.Add("password", tt.userPassword)
			code, _, _ := ts.postForm(t, "/user/signup", form)
			assert.Equal(t, code, tt.wantCode)
		})
	}
}
