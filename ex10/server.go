package ex10

import (
	"fmt"
	"net/http"
	"strconv"
)

func SumHandler(w http.ResponseWriter, r *http.Request) {
	num1, err := strconv.Atoi(r.URL.Query().Get("num1"))
	if err != nil {
		fmt.Fprint(w, http.StatusBadRequest)
		return
	}
	num2, err := strconv.Atoi(r.URL.Query().Get("num2"))
	if err != nil {
		fmt.Fprint(w, http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%d", num1+num2)
}

func MultiHandler(w http.ResponseWriter, r *http.Request) {

	num1, err := strconv.Atoi(r.URL.Query().Get("num1"))
	if err != nil {
		fmt.Fprint(w, http.StatusBadRequest)
		return
	}
	num2, err := strconv.Atoi(r.URL.Query().Get("num2"))
	if err != nil {
		fmt.Fprint(w, http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%d", num1*num2)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func setupRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/sum", SumHandler)
	router.HandleFunc("/multi", MultiHandler)
	router.HandleFunc("/hello", HelloHandler)
	return router
}

func main() {
	router := setupRouter()

	http.ListenAndServe(":8080", router)
}
