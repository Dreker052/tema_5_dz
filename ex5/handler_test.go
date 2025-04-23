package ex5

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHandler(t *testing.T) {

	tests := []struct {
		name         string
		queryParam   url.Values
		expectedCode int
		expectedBody string
	}{
		{"no params", nil, http.StatusBadRequest, ""},
		{"name", url.Values{"name": []string{"Ilya"}}, http.StatusOK, "Hello, Ilya, your name has 4 letters"},
		{"random letters", url.Values{"name": []string{"qwerty"}}, http.StatusOK, "Hello, qwerty, your name has 6 letters"},
		{"two words", url.Values{"name": []string{"one", "two"}}, http.StatusOK, "Hello, one, your name has 3 letters"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			if test.queryParam != nil {
				req.URL.RawQuery = test.queryParam.Encode()
			}

			recorder := httptest.NewRecorder()
			handler := http.HandlerFunc(Handler)

			handler.ServeHTTP(recorder, req)

			status := recorder.Code

			if status != test.expectedCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
			}

			body := recorder.Body.String()
			if body != test.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v", body, test.expectedBody)
			}

		})

	}

}
