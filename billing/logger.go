package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var s3Client *s3.S3
var logBucket string

func initLogger() {
	sess := session.Must(session.NewSession())
	s3Client = s3.New(sess)
	logBucket = os.Getenv("LOG_BUCKET")
}

func logInfo(req PaymentRequest) {
	writeLog(fmt.Sprintf("INFO: updated card for tenant %s", req.TenantID))
}

func logError(err error, req PaymentRequest) {
	writeLog(fmt.Sprintf("ERROR: tenant=%s err=%v", req.TenantID, err))
}

func writeLog(msg string) {
	key := fmt.Sprintf("logs/%d.txt", time.Now().UnixNano())
	input := &s3.PutObjectInput{
		Bucket: aws.String(logBucket),
		Key:    aws.String(key),
		Body:   strings.NewReader(msg),
	}
	s3Client.PutObject(input)
}
