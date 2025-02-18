package users

type GetUserRequest struct {
	Id       *int    `json:"id"`
	Username *string `json:"username,omitempty"`
}
