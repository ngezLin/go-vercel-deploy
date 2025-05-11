package api

import (
	"fmt"
	"net/http"
)

func Greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Greetings!!")
}
