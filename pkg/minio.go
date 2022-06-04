package pkg

import (
	"OverflowBackend/internal/config"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	log "github.com/sirupsen/logrus"
)

func NewMinioClient(config *config.Config) (*minio.Client, error) {
	client, err := minio.New(config.Minio.Url, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Minio.User, config.Minio.Password, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	err = client.MakeBucket(ctx, config.Minio.Bucket, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := client.BucketExists(ctx, config.Minio.Bucket)
		if errBucketExists != nil || !exists {
			return nil, err
		}
		log.Info("We already own ", config.Minio.Bucket)
		return client, nil
	}

	log.Info("Successfully created ", config.Minio.Bucket)
	return client, nil
}
