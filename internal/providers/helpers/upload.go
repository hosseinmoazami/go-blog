package helpers

import (
	"errors"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SaveUploadFile(c *gin.Context, des string) (string, error) {
	file, err := c.FormFile("image")

	if err != nil {
		return "", errors.New("no file is received")
	}

	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension

	if err := c.SaveUploadedFile(file, des+newFileName); err != nil {
		return "", errors.New("unable to save the file")
	}

	return newFileName, nil
}
