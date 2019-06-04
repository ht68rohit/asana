package asana

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

type AsanaArgs struct {
	Name      string `json:"name,omitempty"`
	Workspace string `json:"workspace,omitempty"`
	Public    bool   `json:"public,omitempty"`
	Assignee  string `json:"assignee,omitempty"`
	TaskID    string `json:"task_id,omitempty"`
	ProjectID string `json:"project_id,omitempty"`
	Project   string `json:"id,omitempty"`
	ProjID    string `json:"project,omitempty"`
}

var (
	ACCESS_TOKEN = "0/e39a9a286e8791527776fc5366dfaff9"
)

var _ = Describe("Create project in Asana without access token", func() {

	asana := AsanaArgs{Name: "Test Project", Workspace: "1125282043940580", Public: true}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/createproject", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateProject)
	handler.ServeHTTP(recorder, request)

	Describe("Create project in Asana", func() {
		Context("create project", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create project in Asana with invalid param", func() {

	os.Setenv("ACCESS_TOKEN", ACCESS_TOKEN)

	asana := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/createproject", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateProject)
	handler.ServeHTTP(recorder, request)

	Describe("Create project in Asana", func() {
		Context("create project", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create project in Asana without workspace ID", func() {

	os.Setenv("ACCESS_TOKEN", ACCESS_TOKEN)

	asana := AsanaArgs{Name: "Test Project", Public: true}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/createproject", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateProject)
	handler.ServeHTTP(recorder, request)

	Describe("Create project in Asana", func() {
		Context("create project", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create project in Asana", func() {

	os.Setenv("ACCESS_TOKEN", ACCESS_TOKEN)

	asana := AsanaArgs{Name: "Test Project", Workspace: "1125282043940580", Public: true}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/createproject", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateProject)
	handler.ServeHTTP(recorder, request)

	Describe("Create project in Asana", func() {
		Context("create project", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create task in project Asana", func() {

	os.Setenv("ACCESS_TOKEN", ACCESS_TOKEN)

	asana := AsanaArgs{Name: "Test Task", Workspace: "1125282043940580", Assignee: "demot636@gmail.com"}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/createtask", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTask)
	handler.ServeHTTP(recorder, request)

	Describe("Create project in Asana", func() {
		Context("create project", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("List workspaces in Asana", func() {

	os.Setenv("ACCESS_TOKEN", ACCESS_TOKEN)

	asana := AsanaArgs{}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/listworkspace", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListWorkspace)
	handler.ServeHTTP(recorder, request)

	Describe("List workspaces in Asana", func() {
		Context("list workspaces", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Find task from Asana", func() {

	os.Setenv("ACCESS_TOKEN", ACCESS_TOKEN)

	asana := AsanaArgs{TaskID: "1125513435136524"}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/findtask", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(FindTask)
	handler.ServeHTTP(recorder, request)

	Describe("Find task in Asana", func() {
		Context("find task", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Find project from Asana", func() {

	os.Setenv("ACCESS_TOKEN", ACCESS_TOKEN)

	asana := AsanaArgs{ProjectID: "1125306086215096"}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/findproject", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(FindProject)
	handler.ServeHTTP(recorder, request)

	Describe("Find project in Asana", func() {
		Context("find project", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Update project in Asana", func() {

	os.Setenv("ACCESS_TOKEN", ACCESS_TOKEN)

	asana := AsanaArgs{Project: "1125306086215096", Name: "New Updated Project"}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/updateproject", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateProject)
	handler.ServeHTTP(recorder, request)

	Describe("Update project in Asana", func() {
		Context("update project", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("List tasks from project in Asana", func() {

	os.Setenv("ACCESS_TOKEN", ACCESS_TOKEN)

	asana := AsanaArgs{ProjID: "1125306086215096"}
	requestBody := new(bytes.Buffer)
	jsonErr := json.NewEncoder(requestBody).Encode(asana)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	request, err := http.NewRequest("POST", "/listprojecttasks", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListProjectTasks)
	handler.ServeHTTP(recorder, request)

	Describe("List tasks from project in Asana", func() {
		Context("list tasks", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})
