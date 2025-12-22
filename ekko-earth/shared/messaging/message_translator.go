package messaging

type MessageTranslator[TFrom any, TTo any] interface {
	Translate(message TFrom) (TTo, error)
}
