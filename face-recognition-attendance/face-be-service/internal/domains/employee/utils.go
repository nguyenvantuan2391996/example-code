package employee

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"face-be-service/common/constants"
	"face-be-service/internal/domains/employee/models"
)

func (es *Employee) uploadImageToMinIO(input *models.ImageInsertInput) (string, error) {
	now := time.Now()
	folderName := fmt.Sprintf("%v/%v/%v/%v", "employee", now.Year(), int(now.Month()), now.Day())
	fileExtension := filepath.Ext(input.ImageName)
	fileName := fmt.Sprintf("%v%s", now.UnixNano(), fileExtension)

	minioPath := fmt.Sprintf("%s/%s", folderName, fileName)

	_, err := es.minioRepo.PutObject(
		viper.GetString(constants.MinioBucket),
		minioPath,
		input.ImageFile,
		input.ImageSize,
		minio.PutObjectOptions{},
	)

	if err != nil {
		logrus.Errorf(constants.FormatTaskErr, "PutObject", err)
		return "", err
	}

	return minioPath, nil
}

func (es *Employee) getMinIOPublicURL(imagePath string) string {
	return fmt.Sprintf("http://%s/%s/%s", viper.GetString(constants.MinioURI),
		viper.GetString(constants.MinioBucket), imagePath)
}
