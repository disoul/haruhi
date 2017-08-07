package manager

import (
	"fmt"
	"haruhi/herror"
	"log"
	"net/http"
)

// HaruhiHTTPHandle handle haruhi server errors
type HaruhiHTTPHandle func(http.ResponseWriter, *http.Request) herror.HaruhiError

func (fn HaruhiHTTPHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)

	if err.Error != nil {
		http.Error(w, err.String(), 500)
	}
}

// CreateManagerServer create a manager http server
func CreateManagerServer(port int) {
	err := http.ListenAndServe(fmt.Sprintf("%v", port), nil)

	if err != nil {
		herr := herror.HaruhiError{
			Error:    err,
			ErrorMsg: "faild to create haruhi http server",
		}

		log.Fatal(herr.String())
	}
}
