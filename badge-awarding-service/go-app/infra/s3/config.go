package s3

import "github.com/aws/aws-sdk-go/service/s3"

type Config struct {
	client     *s3.S3
	bucketName string
	bucketKey  string
}
