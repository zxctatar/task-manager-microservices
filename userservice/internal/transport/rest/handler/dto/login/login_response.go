package logindto

type LoginResponse struct {
	FirstName  string `json:"first_name" binding:"required"`
	MiddleName string `json:"middle_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
}
