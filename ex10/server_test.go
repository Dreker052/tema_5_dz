package ex10

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		path     string
		wantCode int
		wantBody string
	}{
		{"hello without name", "GET", "/hello", 200, "Hello, World!"},
		{"hello with name", "GET", "/hello?name=Ilya", 200, "Hello, Ilya!"},
		{"good  sum", "GET", "/sum?num1=10&num2=5", 200, "15"},
		{"good  multiply", "GET", "/multi?num1=10&num2=5", 200, "50"},
		{"Not Found", "GET", "/", 404, "404 page not found\n"},
	}

	router := setupRouter()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, _ := http.NewRequest(test.method, test.path, nil)
			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)

			if rr.Code != test.wantCode {
				t.Errorf("got code %d, want %d", rr.Code, test.wantCode)
			}

			if test.wantBody != "" && rr.Body.String() != test.wantBody {
				t.Errorf("got body %q, want %q", rr.Body.String(), test.wantBody)
			}
		})
	}
}
