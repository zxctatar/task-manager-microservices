package handlmapper

import (
	createdto "projectservice/internal/transport/rest/handler/dto/create"
	deletedto "projectservice/internal/transport/rest/handler/dto/delete"
	createmodel "projectservice/internal/usecase/models/createproject"
	deletemodel "projectservice/internal/usecase/models/deleteproject"
)

func CreateRequestToInput(cr *createdto.CreateRequest, userId uint32) *createmodel.CreateProjectInput {
	return createmodel.NewCreateProjectInput(userId, cr.Name)
}

func CreateOutputToResponse(co *createmodel.CreateProjectOutput) *createdto.CreateResponse {
	return &createdto.CreateResponse{
		IsCreated: co.IsCreated,
	}
}

func DeleteRequestToInput(dr *deletedto.DeleteRequest, userId uint32) *deletemodel.DeleteProjectInput {
	return deletemodel.NewDeleteProjectInput(userId, dr.Name)
}

func DeleteOutputToResponse(do *deletemodel.DeleteProjectOutput) *deletedto.DeleteResponse {
	return &deletedto.DeleteResponse{IsDeleted: do.IsDeleted}
}
