package deletemodel

type DeleteProjectOutput struct {
	IsDeleted bool
}

func NewDeleteProjectOutput(isDeleted bool) *DeleteProjectOutput {
	return &DeleteProjectOutput{
		IsDeleted: isDeleted,
	}
}
