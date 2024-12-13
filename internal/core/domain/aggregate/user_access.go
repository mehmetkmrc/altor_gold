package aggregate

import "github.com/mehmetkmrc/ator_gold/internal/core/domain/entity"

type (
	UserAccess struct {
		User 			*entity.User
		AccessToken 	string 			`json:"access_token"`
		AccessPublic 	string 			`json:"access_public"`
		RefreshToken 	string 			`json:"refresh_token"`
		RefreshPublic 	string 			`json:"refresh_public"`
	}
)

func NewUserAccess(user *entity.User, accessToken, accessPublic, refreshToken, refreshPublic string) *UserAccess{
	return &UserAccess{
		User: user,
		AccessToken: accessToken,
		AccessPublic: accessPublic,
		RefreshToken: refreshToken,
		RefreshPublic: refreshPublic,
	}
}