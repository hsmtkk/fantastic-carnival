package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type schema struct {
	URL    string `dynamo:"url"`
	SHA256 string `dynamo:"sha256"`
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s url", os.Args[0])
	}

	url := os.Args[1]

	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{Region: aws.String("ap-northeast-1")})
	table := db.Table("test")

	var result schema
	if err := table.Get("url", url).One(&result); err != nil {
		log.Fatalf("get failed; %s", err)
	}
}
