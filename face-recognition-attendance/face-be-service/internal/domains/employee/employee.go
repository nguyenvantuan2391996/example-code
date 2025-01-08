package employee

import (
	"context"
	"fmt"

	"github.com/pgvector/pgvector-go"
	"github.com/sirupsen/logrus"

	"face-be-service/common/constants"
	"face-be-service/common/database/entities"
	"face-be-service/common/third_party"
	"face-be-service/internal/domains/employee/models"
	"face-be-service/internal/domains/repository"
)

type Employee struct {
	employeeRepo repository.IEmployeeRepositoryInterface
	minioRepo    repository.IMinioRepositoryInterface
}

func NewEmployeeService(
	employeeRepo repository.IEmployeeRepositoryInterface,
	minioRepo repository.IMinioRepositoryInterface) *Employee {
	return &Employee{
		employeeRepo: employeeRepo,
		minioRepo:    minioRepo,
	}
}

func (es *Employee) Insert(ctx context.Context, input *models.ImageInsertInput) (map[string]interface{}, error) {
	logrus.Info(fmt.Sprintf(constants.FormatBeginTask, "Insert", input))

	// extract the image to vector
	extraction, err := third_party.GetInstance().ExtractImage(&third_party.ImageExtractionRequest{
		Image:    input.ImageFile,
		FileName: input.ImageName,
	})
	if err != nil {
		logrus.Errorf(constants.FormatTaskErr, "ExtractImage", err)
		return nil, err
	}

	// search
	employee, err := es.employeeRepo.GetTopByDistanceType(ctx, constants.EuclideanDistance, extraction.Vector)
	if err != nil {
		logrus.Errorf(constants.FormatTaskErr, "GetTopByDistanceType", err)
		return nil, err
	}

	if employee != nil && employee.EmployeeID != input.EmployeeID {
		return nil, fmt.Errorf("employee id is invalid")
	}

	// seek file
	_, err = input.ImageFile.Seek(0, 0)
	if err != nil {
		logrus.Errorf(constants.FormatTaskErr, "Seek", err)
		return nil, err
	}

	// upload minio
	imagePath, err := es.uploadImageToMinIO(input)
	if err != nil {
		logrus.Errorf(constants.FormatTaskErr, "uploadImageToMinIO", err)
		return nil, err
	}

	// insert database
	err = es.employeeRepo.Create(ctx, &entities.Employee{
		EmployeeID:   input.EmployeeID,
		EmployeeName: input.EmployeeName,
		ImagePath:    imagePath,
		Embedding:    pgvector.NewVector(extraction.Vector),
	})
	if err != nil {
		logrus.Errorf(constants.FormatCreateEntityErr, "employee", err)
		return nil, err
	}

	return map[string]interface{}{
		"employee_id": input.EmployeeID,
	}, nil
}

func (es *Employee) Search(ctx context.Context, input *models.ImageSearchInput) (map[string]interface{}, error) {
	logrus.Info(fmt.Sprintf(constants.FormatBeginTask, "Search", input))

	extraction, err := third_party.GetInstance().ExtractImage(&third_party.ImageExtractionRequest{
		Image:    input.ImageFile,
		FileName: input.ImageName,
	})
	if err != nil {
		logrus.Errorf(constants.FormatTaskErr, "ExtractImage", err)
		return nil, err
	}

	if extraction.Vector == nil {
		return map[string]interface{}{
			"employee_name": "Unknown",
		}, nil
	}

	employee, err := es.employeeRepo.GetTopByDistanceType(ctx, constants.EuclideanDistance, extraction.Vector)
	if err != nil {
		logrus.Errorf(constants.FormatTaskErr, "GetTopByDistanceType", err)
		return nil, err
	}

	if employee == nil {
		return map[string]interface{}{
			"employee_name": "Unknown",
		}, nil
	}

	return map[string]interface{}{
		"employee_id":   employee.EmployeeID,
		"employee_name": employee.EmployeeName,
		"image_path":    es.getMinIOPublicURL(employee.ImagePath),
	}, nil
}
