package domain

type UserAuth struct {
	ID           int    `json:"id"`
	IdentityType int    `json:"identity_type"`
	Identifier   string `json:"identifier"` // 标示
	Credential   string `json:"credential"`
	UserId       int    `json:"user_id"`
}

type IUserAuthUsecase interface {

}