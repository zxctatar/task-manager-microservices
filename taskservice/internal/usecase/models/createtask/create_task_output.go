package createmodel

type CreateTaskOutput struct {
	TaskId uint32
}

func NewCreateOutput(taskId uint32) *CreateTaskOutput {
	return &CreateTaskOutput{
		TaskId: taskId,
	}
}
