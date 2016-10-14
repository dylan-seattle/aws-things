package main 
import (
	"log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
    )
func main () {

svc := s3.New(session.New(&aws.Config{Region: aws.String("us-east-1")}))
result, err := svc.ListBuckets(&s3.ListBucketsInput{})
if err != nil {
    log.Println("Failed to list buckets" , err)
    return 
}

log.Println("Buckets:")
for _, bucket := range result.Buckets {
   log.Println("%s : %s\n", aws.StringValue(bucket.Name), bucket.CreationDate) 
}
}