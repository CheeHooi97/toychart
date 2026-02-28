package handler

import (
	"encoding/json"
	"toychart/errcode"
	"toychart/model"
	"toychart/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateUser(c echo.Context) error {
	var i struct {
		Username   string            `json:"username"`
		Email      string            `json:"email"`
		Platform   string            `json:"platform"`
		DeviceId   string            `json:"deviceId"`
		DeviceInfo map[string]string `json:"deviceInfo"`
		PNSToken   string            `json:"pnsToken"`
	}

	if msg, err := utils.ValidateRequest(c, &i); err != nil {
		return responseValidationError(c, msg)
	}

	user := new(model.User)
	user.Id = utils.UniqueID()
	user.Username = i.Username
	user.Email = i.Email
	user.Status = true
	user.DateTime()

	if err := h.User.Create(user); err != nil {
		return responseError(c, errcode.InternalServerError)
	}

	deviceInfo, err := json.Marshal(i.DeviceInfo)
	if err != nil {
		return responseError(c, errcode.InternalServerError)
	}

	deviceInfoStr, err := utils.EncryptAES(string(deviceInfo))
	if err != nil {
		return responseError(c, errcode.InternalServerError)
	}

	device := model.NewUserDevice()
	device.UserId = user.Id
	device.Platform = model.UserDevicePlatform(i.Platform)
	device.DeviceId = i.DeviceId
	device.DeviceInfo = deviceInfoStr
	device.PNSToken = i.PNSToken

	if err := h.UserDevice.Create(device); err != nil {
		return responseError(c, errcode.InternalServerError)
	}

	token, err := user.GetAccessToken(device)
	if err != nil {
		return responseError(c, errcode.InternalServerError)
	}

	tk := model.NewToken()
	tk.ReferenceId = user.Id
	tk.DeviceId = i.DeviceId
	tk.AccessToken = token

	if err := h.Token.Create(tk); err != nil {
		return responseError(c, errcode.InternalServerError)
	}

	res := &model.UserWithToken{
		User:  user,
		Token: tk.AccessToken,
	}

	return responseJSON(c, res)
}
