package commands

type CommandHandler interface {
	Handle(command Command) error
}
