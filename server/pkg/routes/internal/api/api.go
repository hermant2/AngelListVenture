package api

import (
	"encoding/json"
	"github.com/go-chi/render"
	"github.com/hermant2/angelventureserver/pkg/apperror"
	"io"
	"net/http"
)

type Response struct {
	Model  interface{}
	Status int
}

type errorResponseWrapper struct {
	Error apperror.Standard `json:"error"`
}

func DecodeRequest(body io.ReadCloser, destination interface{}) error {
	return json.NewDecoder(body).Decode(destination)
}

func RenderResponse(writer http.ResponseWriter, request *http.Request, response Response) {
	render.Status(request, response.Status)
	render.JSON(writer, request, response.Model)
}

func RenderError(writer http.ResponseWriter, request *http.Request, err error) {
	switch err := err.(type) {
	case apperror.Standard:
		renderStatusError(writer, request, err)
	default:
		renderStatusError(writer, request, apperror.InternalServerError(apperror.General))
	}
}
func renderStatusError(writer http.ResponseWriter, request *http.Request, err apperror.Standard) {
	render.Status(request, err.Status)
	render.JSON(writer, request, errorResponseWrapper{err})
}
