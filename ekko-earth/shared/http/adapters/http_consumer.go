package adapters

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type HttpConsume[TIncomingMessage any, TOutgoingMessage any] interface {
	Consume(
		vars map[string]string,
		body TIncomingMessage,
		ctx context.Context) (*TOutgoingMessage, error)
}

type HttpConsumer[TIncomingMessage any, TOutgoingMessage any] struct {
	Route *mux.Route
}

type HttpConsumerConfiguration struct {
	Route   string
	Methods []string
}

func NewHttpConsumer[TIncomingMessage any, TOutgoingMessage any](
	server *HttpServer,
	configuration HttpConsumerConfiguration,
	consume HttpConsume[TIncomingMessage, TOutgoingMessage],
) *HttpConsumer[TIncomingMessage, TOutgoingMessage] {
	route := server.Router.HandleFunc(configuration.Route, func(writer http.ResponseWriter, request *http.Request) {
		Handle(writer, request, consume.Consume)
	}).Methods(configuration.Methods...)

	return &HttpConsumer[TIncomingMessage, TOutgoingMessage]{
		Route: route,
	}
}

func Handle[TIncomingMessage any, TOutgoingMessage any](
	writer http.ResponseWriter,
	request *http.Request,
	consume func(map[string]string, TIncomingMessage, context.Context) (*TOutgoingMessage, error),
) error {
	vars := mux.Vars(request)

	var body TIncomingMessage

	err := json.NewDecoder(request.Body).Decode(&body)

	if request.Body != nil {
		err := json.NewDecoder(request.Body).Decode(&body)

		if err != nil && err != io.EOF {
			writer.WriteHeader(http.StatusBadRequest)
			return err
		}
	}

	result, err := consume(vars, body, request.Context())

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(writer).Encode(map[string]string{
			"error": err.Error(),
		})

		return err
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")

	json.NewEncoder(writer).Encode(result)

	return nil
}
