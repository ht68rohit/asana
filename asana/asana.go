package asana

import (
	"encoding/json"
	result "github.com/heaptracetechnology/microservice-asana/result"
	asana "github.com/odeke-em/asana/v1"
	"net/http"
	"os"
)

type Message struct {
	Success    string `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

//CreateProject asana
func CreateProject(responseWriter http.ResponseWriter, request *http.Request) {

	var accessToken = os.Getenv("ACCESS_TOKEN")

	decoder := json.NewDecoder(request.Body)

	var param *asana.ProjectRequest
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	param.Layout = asana.BoardLayout

	client, err := asana.NewClient(accessToken)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	project, projectErr := client.CreateProject(param)
	if projectErr != nil {
		result.WriteErrorResponseString(responseWriter, projectErr.Error())
		return
	}

	bytes, _ := json.Marshal(project)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//CreateTask asana
func CreateTask(responseWriter http.ResponseWriter, request *http.Request) {

	var accessToken = os.Getenv("ACCESS_TOKEN")

	decoder := json.NewDecoder(request.Body)

	var param *asana.TaskRequest
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	client, err := asana.NewClient(accessToken)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	task, taskErr := client.CreateTask(param)
	if taskErr != nil {
		result.WriteErrorResponse(responseWriter, taskErr)
		return
	}

	bytes, _ := json.Marshal(task)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)

}
