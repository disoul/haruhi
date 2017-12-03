package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// HaruhiResponse haruhi server response struct
type HaruhiResponse struct {
	ErrorCode int         `json:"errCode"`
	ErrorMsg  string      `json:"errMsg"`
	Data      interface{} `json:"data"`
}

func (res HaruhiResponse) Encode() (string, error) {
	bytes, err := json.Marshal(res)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// HaruhiHTTPHandle handle haruhi server errors
type HaruhiHTTPHandle func(http.ResponseWriter, *http.Request) (string, HaruhiError)

type RegisterData struct {
	Name     string
	Depend   []string
	Typename string
	Path     string
}

type finishTaskData struct {
	queryId  string
	taskName string
	output   HaruhiOutput
}

func registerTaskHandle(w http.ResponseWriter, r *http.Request) (string, HaruhiError) {
	var data RegisterData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return "", HaruhiError{
			err,
			"register: can not decode registerData",
			JSON_DECODE_ERROR,
		}
	}

	for _, dependTaskName := range data.Depend {
		if RegisteredTasks[dependTaskName].Name == "" {
			return "", HaruhiError{
				// TODO: init error
				err,
				"depend task is not registed",
				UNEXPECT_REGISTER,
			}
		}
	}

	fmt.Print("register Start")
	err = registerTask(data)
	fmt.Print("registerOk")
	if err != nil {
		return "", HaruhiError{
			err,
			"registerTask Error",
			JSON_ENCODE_ERROR,
		}
	}

	res, err := HaruhiResponse{
		Data: "ok",
	}.Encode()

	if err != nil {
		return "", HaruhiError{
			err,
			"can not encode response data",
			JSON_ENCODE_ERROR,
		}
	}

	return res, HaruhiError{}
}

func finishTask(w http.ResponseWriter, r *http.Response) (string, HaruhiError) {
	var data finishTaskData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return "", HaruhiError{
			err,
			"register: can not decode finishTaskData",
			JSON_DECODE_ERROR,
		}
	}

	taskQuery, ok := HaruhiTaskQuery[data.queryId]
	_ = taskQuery
	if !ok {
		return "", HaruhiError{
			Error: fmt.Errorf("can not find queryid in querymap, id: %v", data.queryId),
		}
	}

	// taskQuery.finish()

	return "", HaruhiError{}
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
	fmt.Fprintf(w, res)
}

// CreateManagerServer create a manager http server
func CreateManagerServer(port int) {
	http.HandleFunc("/register", HaruhiHTTPHandle(registerTaskHandle).ServeHTTP)

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)

	if err != nil {
		herr := HaruhiError{
			Error:    err,
			ErrorMsg: "faild to create haruhi http server",
		}

		log.Fatal(herr.String())
	}
}
