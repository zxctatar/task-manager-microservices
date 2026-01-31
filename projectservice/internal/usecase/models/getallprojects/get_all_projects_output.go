package getallmodel

import projectdomain "projectservice/internal/domain/project"

type GetAllProjectsOutput struct {
	Projects []*projectdomain.ProjectDomain
}

func NewGetAllProjectsOutput(projects []*projectdomain.ProjectDomain) *GetAllProjectsOutput {
	return &GetAllProjectsOutput{
		Projects: projects,
	}
}
