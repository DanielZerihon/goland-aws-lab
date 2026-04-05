package main

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type MockS3Client struct {
	ListBucketsOutput  *s3.ListBucketsOutput
	CreateBucketOutput *s3.CreateBucketOutput
}

type MockS3Uploader struct {
	UploadOutput *manager.UploadOutput
}

func (m MockS3Client) ListBuckets(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error) {
	return m.ListBucketsOutput, nil
}
func (m MockS3Client) CreateBucket(ctx context.Context, params *s3.CreateBucketInput, optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error) {
	return m.CreateBucketOutput, nil
}

func (m MockS3Uploader) Upload(ctx context.Context, input *s3.PutObjectInput, opts ...func(*manager.Uploader)) (*manager.UploadOutput, error) {
	return m.UploadOutput, nil
}

func TestCreateS3Bucket(t *testing.T) {
	ctx := context.Background()
	err := createS3Bucket(ctx, MockS3Client{
		ListBucketsOutput: &s3.ListBucketsOutput{
			Buckets: []types.Bucket{
				{
					Name: aws.String("test-bucket"),
				},
				{
					Name: aws.String("test-bucket-2"),
				},
			},
		},
		CreateBucketOutput: &s3.CreateBucketOutput{},
	})

	if err != nil {
		t.Fatalf("createS3Bucket error: %v", err)

	}
}


// func uploadToS3Bucket(ctx context.Context, uploader S3Uploader) error {

func TestUploadToS3(t *testing.T) {

	MockUploader := MockS3Uploader{
		UploadOutput: &manager.UploadOutput{},
	}
	err := uploadToS3Bucket(context.Background(), MockUploader, "testdata/test.txt")
	if err != nil {
		t.Fatalf("TestUploadToS3 error: %v", err)
	}
}
