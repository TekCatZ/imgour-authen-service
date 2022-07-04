package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		exitErrorf("Bucket, region, access key, secret, item's name and output's name required\n"+
			"Usage: %s bucket_name region access_key secret item_name output_name",
			os.Args[0])
	}

	bucket := os.Args[1]
	region := os.Args[2]
	accessKey := os.Args[3]
	secret := os.Args[4]
	item := os.Args[5]
	output := os.Args[6]

	file, err := os.Create(output)
	if err != nil {
		exitErrorf("Unable to open file %q, %v", item, err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			exitErrorf("Unable to close file %q, %v", item, err)
		}
	}(file)

	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secret, ""),
	})
	if err != nil {
		exitErrorf("Unable to create session, %v", err)
	}

	downloader := s3manager.NewDownloader(sess)
	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		})
	if err != nil {
		exitErrorf("Unable to download item %q, %v", item, err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}

func exitErrorf(msg string, args ...interface{}) {
	_, err := fmt.Fprintf(os.Stderr, msg+"\n", args...)
	if err != nil {
		panic(err)
		return
	}
	os.Exit(1)
}
