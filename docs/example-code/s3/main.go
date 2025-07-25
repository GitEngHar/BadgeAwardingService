package main

import (
	"bytes"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	filePath   = "./test.jpg"
	bucketName = "my-test-bucket"
	s3ImgPath  = "image/test.jpg"
	awsRegion  = "ap-northeast-1"
)

func main() {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// session create
	newSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// s3 client create
	svc := s3.New(newSession, &aws.Config{Region: aws.String(awsRegion)})

	downloadKey := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(s3ImgPath),
	}

	image, err := svc.GetObject(downloadKey)
	if err != nil {
		log.Fatal(err)
	}
	//image -> bytes.Buffer
	buf := new(bytes.Buffer)
	buf.ReadFrom(image.Body)

	// ファイルに書き込みします。
	_, err = file.Write(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Finished Download from S3 ")
}
