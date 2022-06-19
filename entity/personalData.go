package entity

type PersonalData struct {
	Id           string // uuid without dash
	CreatedAt    string
	UpdatedAt    string
	IpAddress    string
	DeviceAccess string
	Name         string
	Email        string
}
type CreatePersonalData struct {
	IpAddress    string `json:"ip_address" binding:"required"`
	DeviceAccess string `json:"device_access" binding:"required"`
}

type CreateDetailPersonalData struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type PersonalDataResponse struct {
	Id           string `json:"id"`
	IpAddress    string `json:"ip_address"`
	DeviceAccess string `json:"device_access"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Token        string `json:"token"`
}

const (
	ValidationValid   = "valid"
	ValidationInvalid = "invalid"
)

type ValidationResponse struct {
	Validation    string `json:"validation"`
	MessageDetail string `json:"message_detail"`
}

func (p *PersonalData) ToPersonalDataResponse(token string) PersonalDataResponse {
	return PersonalDataResponse{
		Id:           p.Id,
		IpAddress:    p.IpAddress,
		DeviceAccess: p.DeviceAccess,
		Name:         p.Name,
		Email:        p.Email,
		Token:        token,
	}
}
