package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"toychart/config"
	"toychart/errcode"
	"toychart/kit/oss"

	"github.com/labstack/echo/v4"
)

func (h *Handler) UpdateUser(c echo.Context) error {
	if err := c.Request().ParseMultipartForm(32 << 20); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid multipart form")
	}

	id := c.Param("id")
	username := c.FormValue("username")

	file, fileHeader, err := c.Request().FormFile("photo")
	if err != nil && err != http.ErrMissingFile {
		log.Printf("FormFile error: %v", err)
	}

	user, err := h.User.GetById(id)
	if err != nil {
		return responseError(c, errcode.InternalServerError)
	} else if user == nil {
		return responseError(c, errcode.FailedGetUser)
	}

	if username != "" {
		check, err := h.User.GetByUsername(username)
		if err != nil {
			return responseError(c, errcode.InternalServerError)
		}

		if !check {
			user.Username = username
		}
	}

	if file == nil {
		return responseError(c, errcode.FileError)
	} else {
		defer file.Close()
		var fileName string

		fmt.Println("fileHeader.Filename: ", fileHeader.Filename)

		file, err := fileHeader.Open()
		if err != nil {
			return responseError(c, errcode.InternalServerError)
		}

		defer file.Close()

		fileByte, err := io.ReadAll(file)
		if err != nil {
			return responseError(c, errcode.InternalServerError)
		}

		ext := filepath.Ext(fileHeader.Filename)

		fileName = "user_profile_" + user.Id + ext

		if err := oss.Upload(config.OSSBucket, fileName, fileByte); err != nil {
			return responseError(c, errcode.InternalServerError)
		}

		user.PhotoURL = fileName
	}

	user.UpdateDt()
	if err := h.User.Update(user); err != nil {
		return responseError(c, errcode.InternalServerError)
	}

	return responseJSON(c, true)
}
