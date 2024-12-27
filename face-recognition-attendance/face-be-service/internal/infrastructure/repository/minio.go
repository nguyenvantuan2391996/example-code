package repository

import (
	"context"
	"io"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
)

type MinioRepository struct {
	client  *minio.Client
	timeout time.Duration
}

func NewMinioRepository(client *minio.Client, timeout time.Duration) *MinioRepository {
	return &MinioRepository{
		client:  client,
		timeout: timeout,
	}
}

func (mc *MinioRepository) BucketExists(bucketName string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mc.timeout)
	defer cancel()

	exist, err := mc.client.BucketExists(ctx, bucketName)
	if err != nil {
		return false, err
	}
	return exist, err
}

func (mc *MinioRepository) MakeBucket(bucketName string, opts minio.MakeBucketOptions) error {
	ctx, cancel := context.WithTimeout(context.Background(), mc.timeout)
	defer cancel()

	err := mc.client.MakeBucket(ctx, bucketName, opts)
	if err != nil {
		return err
	}
	return nil
}

func (mc *MinioRepository) PutObject(bucketName, objectName string, rd io.Reader, objectSize int64, opts minio.PutObjectOptions) (minio.UploadInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mc.timeout)
	defer cancel()

	info, err := mc.client.PutObject(ctx, bucketName, objectName, rd, objectSize, opts)
	if err != nil {
		return minio.UploadInfo{}, err
	}
	return info, nil
}

func (mc *MinioRepository) FPutObject(bucketName, objectName, filePath string, opts minio.PutObjectOptions) (minio.UploadInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mc.timeout)
	defer cancel()

	info, err := mc.client.FPutObject(ctx, bucketName, objectName, filePath, opts)
	if err != nil {
		return minio.UploadInfo{}, err
	}
	return info, nil
}

func (mc *MinioRepository) Presign(method, bucketName, objectName string, expires time.Duration, reqParams url.Values) (*url.URL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mc.timeout)
	defer cancel()

	u, err := mc.client.Presign(ctx, method, bucketName, objectName, expires, reqParams)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (mc *MinioRepository) GetObject(bucketName, objectName string, opts minio.GetObjectOptions) (*minio.Object, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mc.timeout)
	defer cancel()

	reader, err := mc.client.GetObject(ctx, bucketName, objectName, opts)
	if err != nil {
		return nil, err
	}
	return reader, nil
}

func (mc *MinioRepository) GetObjectToByte(bucketName, objectName string, opts minio.GetObjectOptions) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mc.timeout)
	defer cancel()

	reader, err := mc.client.GetObject(ctx, bucketName, objectName, opts)
	if err != nil {
		return nil, err
	}
	defer func(reader *minio.Object) {
		errClose := reader.Close()
		if errClose != nil {
			return
		}
	}(reader)

	content, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (mc *MinioRepository) RemoveObject(bucketName, objectName string, opts minio.RemoveObjectOptions) error {
	ctx, cancel := context.WithTimeout(context.Background(), mc.timeout)
	defer cancel()

	err := mc.client.RemoveObject(ctx, bucketName, objectName, opts)
	if err != nil {
		return err
	}
	return nil
}

func (mc *MinioRepository) InitBucket(bucketName string) error {
	existed, err := mc.BucketExists(bucketName)
	if err != nil {
		return err
	}

	if !existed {
		err = mc.MakeBucket(bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}
