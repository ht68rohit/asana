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
	Name      string `json:"name"`
	Workspace string `json:"workspace"`
	Public    bool   `json:"public"`
}

var (
	ACCESS_TOKEN = "0/e39a9a286e8791527776fc5366dfaff9"
)

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
