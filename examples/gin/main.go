// Copyright 2017 Matt Ho
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myml/swag"
	"github.com/myml/swag/endpoint"

	"github.com/myml/swag/swagger/ui"
)

func handle(c *gin.Context) {
	io.WriteString(c.Writer, "Insert your code here")
}

// Category example from the swagger pet store
type Category struct {
	ID   int64  `json:"category"`
	Name string `json:"name" enum:"[\"cat\",\"dog\"]"`
}

// Pet example from the swagger pet store
type Pet struct {
	ID        int64    `json:"id"`
	Category  Category `json:"category" required:"true"`
	Name      string   `json:"name" description:"Cute name"`
	PhotoUrls []string `json:"photoUrls"`
	Tags      []string `json:"tags"`
}

func main() {
	api := swag.New()
	router := gin.Default()

	router.GET("/pet/:id", handle)
	api.AddEndpoint(endpoint.New("get", "/pet/{petId}", "Find pet by ID",
		endpoint.Handler(handle),
		endpoint.RequestHeader("X-API-Version", "string", "api version", true),
		endpoint.QueryEnum("order", "string", "order by", false, []string{"asc", "desc"}),
		endpoint.Path("petId", "integer", "ID of pet to return", true),
		endpoint.Response(http.StatusOK, Pet{}, "successful operation"),
	))

	router.POST("/pet", handle)
	api.AddEndpoint(endpoint.New("post", "/pet", "Add a new pet to the store",
		endpoint.Handler(handle),
		endpoint.Description("Additional information on adding a pet to the store"),
		endpoint.Body(Pet{}, "Pet object that needs to be added to the store", true),
		endpoint.Response(http.StatusOK, Pet{}, "Successfully added pet"),
	))

	enableCors := true
	router.GET("/swagger", gin.WrapH(api.Handler(enableCors)))
	router.GET("/swagger_ui/*filepath", gin.WrapH(ui.Handler("/swagger_ui/", api)))
	http.ListenAndServe(":8080", router)
}
