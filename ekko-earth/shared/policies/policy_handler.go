package policies

type Policy interface {
	Apply(value any) error
}

type PolicyHandler struct {
	Policies []Policy
}

func NewPolicyHandler(policies ...Policy) *PolicyHandler {
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
