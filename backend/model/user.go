package model

import (
	"time"
	"toychart/config"
	"toychart/utils"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	Id       string `gorm:"type:varchar(36);primaryKey" json:"id"`
	Username string `gorm:"type:varchar(255)" json:"username"`
	Email    string `gorm:"type:varchar(255)" json:"email"`
	PhotoURL string `gorm:"type:varchar(255)" json:"photoUrl"`
	Status   bool   `gorm:"type:boolean;default:false" json:"status"`
	BaseModel
}

func NewUser() *User {
	now := time.Now().UTC()

	m := new(User)
	m.Id = utils.UniqueID()
	m.CreatedDateTime = now
	m.UpdatedDateTime = now

	return m
}

func (m *User) DateTime() {
	m.CreatedDateTime = time.Now().UTC()
	m.UpdatedDateTime = time.Now().UTC()
}

func (m *User) UpdateDt() {
	m.UpdatedDateTime = time.Now().UTC()
}

func (u User) GetAccessToken(device *UserDevice) (string, error) {
	claims := jwt.RegisteredClaims{
		ID:      u.Id,
		Subject: device.Id,
		Audience: jwt.ClaimStrings([]string{
			string(device.Platform),
			device.DeviceId,
		}),
		Issuer:    "toychart",
		NotBefore: jwt.NewNumericDate(time.Now().UTC()),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		// ExpiresAt: jwt.NewNumericDate(time.Now().UTC().AddDate(0, 1, 0)),
	}

	return jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(config.AuthenticationPrivateKey)
}
