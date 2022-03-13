package handler

import (
	"log"
	"net/http"

	"github.com/Triticumdico/dashboard/src/app/backend/resource/dummy"

	"github.com/emicklei/go-restful/v3"
)

// APIHandler is a representation of API handler.
type APIHandler struct {
	//
}

// CreateHTTPAPIHandler creates a new HTTP handler that handles all requests to the API of the backend.
func CreateHTTPAPIHandler() (http.Handler, error) {

	apiHandler := APIHandler{}
	wsContainer := restful.NewContainer()

	apiV1Ws := new(restful.WebService)

	apiV1Ws.Path("/api/v1").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	wsContainer.Add(apiV1Ws)

	apiV1Ws.Route(
		apiV1Ws.GET("/test").
			To(apiHandler.handleGetTest).
			Writes(dummy.TableRows{}))

	return wsContainer, nil

}

func (apiHandler *APIHandler) handleGetTest(request *restful.Request, response *restful.Response) {

	result, err := dummy.GetTableRows()
	if err != nil {
		log.Fatal(err)
	}

	response.WriteHeaderAndEntity(http.StatusOK, result)
}
