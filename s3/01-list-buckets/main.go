package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	profile := "put-my-cli-config-profile-name-here"
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithSharedConfigProfile(profile),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 client
	clientS3 := s3.NewFromConfig(cfg)

	// Get List of buckets
	output, err := clientS3.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		log.Fatal(err)
	}

	for i, b := range output.Buckets {
		log.Printf("Name bucket %d: %s\n", i+1, aws.ToString(b.Name))
	}

	// Show Owner S3 service
	log.Printf("Owner: %s\n", aws.ToString(output.Owner.DisplayName))
}
