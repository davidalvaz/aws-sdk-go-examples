package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
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

	client := sts.NewFromConfig(cfg)
	identity, err := client.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Account: %s, Arn: %s, UserId: %s\n", aws.ToString(identity.Account), aws.ToString(identity.Arn), aws.ToString(identity.UserId))
}
