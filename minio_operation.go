package main

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"net/url"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MC *minio.Client

func init() {
	endpoint := "localhost:9000"
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("MINIO_SECRET_ACCESS_KEY")

	// Initialize minio client object.
	var err error
	MC, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", MC) // minioClient is now setup
}

func CheckBuckets(ctx context.Context) error {
	buckets, err := MC.ListBuckets(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}
	return nil
}
func Putfile(ctx context.Context, src multipart.File, file_size int64, file_name string) error {
	_, err := MC.PutObject(ctx, "files", file_name, src, file_size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func GeneratePresignedURL(ctx context.Context, file_name string) (string, error) {
	reqParams := make(url.Values)
	presignedURL, err := MC.PresignedGetObject(ctx, "files", file_name, time.Second*24*60*60, reqParams)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return presignedURL.String(), nil
}
