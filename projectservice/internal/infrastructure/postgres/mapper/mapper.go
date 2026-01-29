package posmapper

import (
	projectdomain "projectservice/internal/domain/project"
	posmodels "projectservice/internal/infrastructure/postgres/models"
)

func DomainToModel(pd *projectdomain.ProjectDomain) *posmodels.ProjectPosModel {
	return posmodels.NewProjectPosModel(pd.OwnerId, pd.Name)
}

func ModelToDomain(pm *posmodels.ProjectPosModel) *projectdomain.ProjectDomain {
	return projectdomain.RestoreProjectDomain(pm.Id, pm.OwnerId, pm.Name, pm.CreatedAt)
}
