package getallmodel

type GetAllProjectsInput struct {
	OwnerId uint32
}

func NewGetAllProjectsInput(ownerId uint32) *GetAllProjectsInput {
	return &GetAllProjectsInput{
		OwnerId: ownerId,
	}
}
