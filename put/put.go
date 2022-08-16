package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
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
	hash, err := getHash(url)
	if err != nil {
		log.Fatalf("getHash failed; %s", err)
	}

	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{Region: aws.String("ap-northeast-1")})
	table := db.Table("test")

	item := schema{URL: url, SHA256: hash}
	if err := table.Put(item).Run(); err != nil {
		log.Fatalf("put failed; %s", err)
	}
}

func getHash(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("http.Get failed; %w", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("io.ReadAll failed; %w", err)
	}
	return getSHA256(data), nil
}

func getSHA256(data []byte) string {
	sum := sha256.Sum256(data)
	return fmt.Sprintf("%x", sum)
}
