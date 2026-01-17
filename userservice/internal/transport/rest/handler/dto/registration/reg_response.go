package regdto

type RegistrationResponse struct {
	IsRegistered bool `json:"is_registered" binding:"required"`
}
