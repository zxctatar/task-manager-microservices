package deletemodel

type DeleteProjectInput struct {
	Name string
}

func NewDeleteProjectInput(name string) *DeleteProjectInput {
	return &DeleteProjectInput{
		Name: name,
	}
}
