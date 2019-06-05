package asana

import (
	"encoding/json"
	"fmt"
	asana "github.com/heaptracetechnology/microservice-asana/pkg/asana/v1"
	result "github.com/heaptracetechnology/microservice-asana/result"
	"log"
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
		fmt.Println("decodeErr :: ", decodeErr)
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	res, _ := json.Marshal(param)
	fmt.Println("res :::::", string(res))

	client, err := asana.NewClient(accessToken)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	task, taskErr := client.CreateTask(param)
	if taskErr != nil {
		fmt.Println("taskErr :: ", taskErr.Error())
		result.WriteErrorResponseString(responseWriter, taskErr.Error())
		return
	}
	fmt.Println("task :: ", task)
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

//DeleteTask asana
func DeleteTask(responseWriter http.ResponseWriter, request *http.Request) {

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

	if err := client.DeleteTask(param.TaskID); err != nil {
		fmt.Println("err ::", err)
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	message := Message{"true", "Task deleted successfully", http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//ListTask asana
func ListTask(responseWriter http.ResponseWriter, request *http.Request) {

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

	taskPagesChan, err := client.ListMyTasks(param)
	if err != nil {
		log.Fatal(err)
	}

	var listTasks []*asana.Task
	pageCount := 0
	for page := range taskPagesChan {
		if err := page.Err; err != nil {
			log.Printf("Page: #%d err: %v", pageCount, err)
			continue
		}

		for _, task := range page.Tasks {
			listTasks = append(listTasks, task)
		}
		pageCount += 1
	}
	bytes, _ := json.Marshal(listTasks)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)

}

//ListWorkspace asana
func ListWorkspace(responseWriter http.ResponseWriter, request *http.Request) {

	var accessToken = os.Getenv("ACCESS_TOKEN")

	client, err := asana.NewClient(accessToken)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}
	workspacesChan, _ := client.ListMyWorkspaces()

	var listWorkspace []*asana.Workspace

	pageCount := 0
	for page := range workspacesChan {
		if err := page.Err; err != nil {
			continue
		}

		for _, workspace := range page.Workspaces {
			listWorkspace = append(listWorkspace, workspace)
		}
		pageCount += 1
	}
	bytes, _ := json.Marshal(listWorkspace)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)

}

//FindTask asana
func FindTask(responseWriter http.ResponseWriter, request *http.Request) {

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

	taskDetails, err := client.FindTaskByID(param.TaskID)
	if err != nil {
		result.WriteErrorResponseString(responseWriter, err.Error())
		return
	}

	bytes, _ := json.Marshal(taskDetails)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//FindProject asana
func FindProject(responseWriter http.ResponseWriter, request *http.Request) {

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

	projectDetails, err := client.FindProjectByID(param.ProjectID)
	if err != nil {
		result.WriteErrorResponseString(responseWriter, err.Error())
		return
	}
	bytes, _ := json.Marshal(projectDetails)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//UpdateProject asana
func UpdateProject(responseWriter http.ResponseWriter, request *http.Request) {
	var accessToken = os.Getenv("ACCESS_TOKEN")

	decoder := json.NewDecoder(request.Body)

	var param *asana.ProjectRequest
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

	update, updateErr := client.UpdateProject(param)
	if updateErr != nil {
		result.WriteErrorResponseString(responseWriter, updateErr.Error())
		return
	}

	bytes, _ := json.Marshal(update)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//ListProjectTasks asana
func ListProjectTasks(responseWriter http.ResponseWriter, request *http.Request) {

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

	taskPagesChan, _, err := client.ListTasksForProject(param)
	if err != nil {
		log.Fatal(err)
	}

	var tasklist []*asana.Task
	pageCount := 0
	for page := range taskPagesChan {
		if err := page.Err; err != nil {
			log.Printf("Page: #%d err: %v", pageCount, err)
			continue
		}

		for _, task := range page.Tasks {
			tasklist = append(tasklist, task)
		}
		pageCount += 1
	}

	bytes, _ := json.Marshal(tasklist)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
