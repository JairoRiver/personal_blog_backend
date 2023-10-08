package assets

import (
	"bytes"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Config struct {
	AwsAccessKey  string
	AwsSecret     string
	AwsRegion     string
	AWSBucketName string
}

type S3Store struct {
	s3Client *s3.Client
	region   string
	bucket   string
}

func NewS3AssetStore(cnf S3Config) (ImageStorer, error) {
	creds := credentials.NewStaticCredentialsProvider(cnf.AwsAccessKey, cnf.AwsSecret, "")

	sdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(cnf.AwsRegion), config.WithCredentialsProvider(creds))
	if err != nil {
		log.Println("Couldn't load default configuration.")
		return nil, err
	}

	s3Client := s3.NewFromConfig(sdkConfig)
	maker := &S3Store{
		s3Client: s3Client,
		region:   cnf.AwsRegion,
		bucket:   cnf.AWSBucketName,
	}

	return maker, nil
}

func (ms *S3Store) UploadImage(ctx context.Context, file []byte, path string, name string) (string, error) {
	// Save the file to specific dst.
	r := bytes.NewReader(file)
	key := path + "/" + name + ".jpg"

	_, err := ms.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(ms.bucket),
		Key:    aws.String(key),
		Body:   r,
	})
	if err != nil {
		log.Printf("Couldn't upload file %v to %v. Here's why: %v\n", name, path, err)
	}

	imageURL := "https://" + ms.bucket + ".s3-" + ms.region + ".amazonaws.com/" + key

	return imageURL, nil
}

func (ms *S3Store) DeleteImage(ctx context.Context, path string, name string) error {
	key := path + "/" + name + ".jpg"

	_, err := ms.s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(ms.bucket),
		Key:    aws.String(key),
	})

	return err
}
