package policies

import "log/slog"

type Policy interface {
	Apply(value any) error
}

type PolicyHandler struct {
	Policies []Policy
}

func NewPolicyHandler(policies ...Policy) *PolicyHandler {
	slog.Debug("Creating policy handler")

	return &PolicyHandler{
		Policies: policies,
	}
}

func (handler *PolicyHandler) Apply(value any) error {
	for _, policy := range handler.Policies {
		if err := policy.Apply(value); err != nil {
			return err
		}
	}

	return nil
}
