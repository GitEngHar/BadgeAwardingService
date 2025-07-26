package main

import (
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// 変数を定義
	bucket := "my-test-bucket-badge-award-valid"
	key := "image/test.png" // S3内の画像パス
	output := "test.png"    // ローカルに保存するファイル名

	// セッション作成（デフォルトの環境変数・認証情報を利用）
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"), // 東京リージョンなどに変更してください
	})
	if err != nil {
		fmt.Println("Failed to create session,", err)
		return
	}

	// S3クライアント作成
	svc := s3.New(sess)

	// GetObjectを呼び出して画像を取得
	result, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		fmt.Println("Failed to download image,", err)
		return
	}
	defer result.Body.Close()

	// ローカルファイルに保存
	file, err := os.Create(output)
	if err != nil {
		fmt.Println("Failed to create file,", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, result.Body)
	if err != nil {
		fmt.Println("Failed to write file,", err)
		return
	}

	fmt.Println("Image downloaded successfully:", output)
}
