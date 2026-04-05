package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const bucketName = "bombo-s3"
const regionName = "us-east-1"

type S3Client interface {
	ListBuckets(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
	CreateBucket(ctx context.Context, params *s3.CreateBucketInput, optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error)
}

type S3Uploader interface {
	Upload(ctx context.Context, input *s3.PutObjectInput, opts ...func(*manager.Uploader)) (*manager.UploadOutput, error)
}

func main() {
	var (
		s3Client *s3.Client
		err      error
	)
	ctx := context.Background()
	if s3Client, err = initS3Client(ctx); err != nil {
		fmt.Printf("unable to initS3Client: %v\n", err)
		os.Exit(1)
	}

	if err = createS3Bucket(ctx, s3Client); err != nil {
		fmt.Printf("unable to createS3Bucket: %v\n", err)
		os.Exit(1)
	}

	if err = uploadToS3Bucket(ctx, manager.NewUploader(s3Client), "upload-example.txt"); err != nil {
		fmt.Printf("unable to uploadToS3Bucket: %v\n", err)
		os.Exit(1)
	}

	fmt.Print("upload To S3 Bucket complite")

}

func initS3Client(ctx context.Context) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(regionName))
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}
	return s3.NewFromConfig(cfg), nil
}

func createS3Bucket(ctx context.Context, s3Client S3Client) error {
	allBuckets, err := s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return fmt.Errorf("ListBuckets call failed: %w", err)
	}

	found := false
	for _, bucket := range allBuckets.Buckets {
		if *bucket.Name == bucketName {
			found = true
			fmt.Printf("bucket created already - skip creation: %s\n", bucketName)
		}
	}

	if !found {
		_, err = s3Client.CreateBucket(ctx, &s3.CreateBucketInput{
			Bucket: aws.String(bucketName),
		})
		if err != nil {
			return fmt.Errorf("create bucket API call failed: %w", err)
		}
	}

	return nil
}

func uploadToS3Bucket(ctx context.Context, uploader S3Uploader, filename string) error {
	// uploader := manager.NewUploader(s3Client)

	_, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
		Body:   strings.NewReader("this is explame upload"),
	})

	if err != nil {
		return fmt.Errorf("Upload API call failed: %w", err)
	}

	return nil
}
