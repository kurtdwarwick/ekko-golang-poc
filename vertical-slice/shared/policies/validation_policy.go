package policies

type ValidationPolicy interface {
	Apply(value any) error
}
