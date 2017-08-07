package manager

import (
	"encoding/json"
	"fmt"
	"haruhi/herror"
	"log"
	"net/http"
)

// HaruhiResponse haruhi server response struct
type HaruhiResponse struct {
	ErrorCode int
	ErrorMsg  string
	data      interface{}
}

func (res HaruhiResponse) Encode() (string, error) {
	bytes, err := json.Marshal(res)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// HaruhiHTTPHandle handle haruhi server errors
type HaruhiHTTPHandle func(http.ResponseWriter, *http.Request) (string, herror.HaruhiError)

func registerTaskHandle(w http.ResponseWriter, r *http.Request) (string, herror.HaruhiError) {
	type registerData struct {
		name     string
		depend   []string
		typename string
		path     string
	}

	var data registerData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return "", herror.HaruhiError{
			err,
			"register: can not decode registerData",
			herror.JSON_DECODE_ERROR,
		}
	}

	for _, dependTaskName := range data.depend {
		if RegisteredTasks[dependTaskName].Name == "" {
			return "", herror.HaruhiError{
				// TODO: init error
				err,
				"depend task is not registed",
				herror.UNEXPECT_REGISTER,
			}
		}
	}

	registerTask(data.name, data.typename, data.path, data.depend)

	return "", herror.HaruhiError{}
}

func (fn HaruhiHTTPHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := fn(w, r)

	if err.Error != nil {
		response := HaruhiResponse{
			ErrorCode: int(err.ErrorCode),
			ErrorMsg:  err.ErrorMsg,
		}
		res, e := response.Encode()
		if e != nil {
			http.Error(w, err.String(), 500)
			return
		}
		http.Error(w, res, 500)

		return
	}

	w.Header().Set("Content-Type", "application/json")
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
