package asana

import (
	"encoding/json"
	result "github.com/heaptracetechnology/microservice-asana/result"
	asana "github.com/odeke-em/asana/v1"
	"net/http"
	"os"
)

type AsanaArgument struct {
	TaskID    string `json:"task_id,omitempty"`
	ProjectID string `json:"project_id,omitempty"`
}

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

	param.Layout = asana.ListLayout

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
		result.WriteErrorResponseString(responseWriter, taskErr.Error())
		return
	}

	bytes, _ := json.Marshal(task)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)

}

//DeleteProject asana
func DeleteProject(responseWriter http.ResponseWriter, request *http.Request) {

	var accessToken = os.Getenv("ACCESS_TOKEN")

	decoder := json.NewDecoder(request.Body)

	var param AsanaArgument
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

	if err := client.DeleteProjectByID(param.ProjectID); err != nil {
		result.WriteErrorResponseString(responseWriter, err.Error())
		return
	}

	message := Message{"true", "Project deleted successfully", http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
