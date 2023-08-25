package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
)

func main() {
	profile := "training-dev"
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile(profile))
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	output, err := client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{Bucket: aws.String("datasync-dev")})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Bucket name: %s", *output.Name)
	log.Printf("first page results:\n")
	for _, object := range output.Contents {
		log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
	}

}
