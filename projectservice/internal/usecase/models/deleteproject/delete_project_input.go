package deletemodel

type DeleteProjectInput struct {
	OwnerId   uint32
	ProjectId uint32
}

func NewDeleteProjectInput(ownerId uint32, projectId uint32) *DeleteProjectInput {
	return &DeleteProjectInput{
		OwnerId:   ownerId,
		ProjectId: projectId,
	}
}
