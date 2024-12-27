package repository

import (
	"io"

	"github.com/minio/minio-go/v7"
)

//go:generate mockgen -package=repository -destination=iminio_mock.go -source=iminio.go

type IMinioRepositoryInterface interface {
	BucketExists(bucketName string) (bool, error)
	MakeBucket(bucketName string, opts minio.MakeBucketOptions) error
	PutObject(bucketName, objectName string, rd io.Reader, objectSize int64,
		opts minio.PutObjectOptions) (minio.UploadInfo, error)
	InitBucket(bucketName string) error
}
