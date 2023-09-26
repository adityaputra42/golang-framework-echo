package request

type UpdateUser struct {
	Id          int    `json:"id"`
	OldPassword string `json:"old_password"`
	Password    string `json:"password"`
}
