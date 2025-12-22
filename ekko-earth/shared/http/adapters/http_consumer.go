package adapters

import (
	"encoding/json"

	"net/http"

	"github.com/ekko-earth/shared/messaging"
)

type HttpConsumer[TIncomingMessage any, TMessage any] struct{}

type HttpConsumerConfiguration struct {
	Route   string
	Methods []string
}

func NewHttpConsumer[TIncomingMessage any, TMessage any](
	server HttpServer,
	messageTranslator messaging.MessageTranslator[TIncomingMessage, TMessage],
	messageHandler messaging.MessageHandler[TMessage],
	configuration HttpConsumerConfiguration,
) *HttpConsumer[TIncomingMessage, TMessage] {
	organisationsRoute := server.Router.PathPrefix(configuration.Route).Subrouter()

	organisationsRoute.HandleFunc("", func(writer http.ResponseWriter, request *http.Request) {
		handle(writer, request, messageHandler, messageTranslator)
	}).Methods(configuration.Methods...)

	return &HttpConsumer[TIncomingMessage, TMessage]{}
}

func handle[TIncomingMessage any, TMessage any](
	writer http.ResponseWriter,
	request *http.Request,
	messageHandler messaging.MessageHandler[TMessage],
	messageTranslator messaging.MessageTranslator[TIncomingMessage, TMessage],
) error {
	var incomingMessage TIncomingMessage

	err := json.NewDecoder(request.Body).Decode(&incomingMessage)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return err
	}

	translatedMessage, err := messageTranslator.Translate(incomingMessage)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return err
	}

	result, err := messageHandler.Handle(translatedMessage, request.Context())

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
