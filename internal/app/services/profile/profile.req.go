package profile

type UpdateProfileRequest struct {
	Username string `json:"username"`
	Tag      string `json:"tag"`
	Bio      string `json:"bio"`
}
