//assumes you have the following environment variables setup for AWS session creation 
// AWS_SDK_LOAD_CONFIG=1
// AWS_ACCESS_KEY_ID=XXXXXXXXXX
// AWS_SECRET_ACCESS_KEY=XXXXXXXX
// AWS_REGION=us-west-2( or AWS_DEFAULT_REGION=us-east-1 if you are having trouble)

package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func mainc() {

	fmt.Println("creating session")
	sess := session.Must(session.NewSession())
	fmt.Println("session created")

	svc := sns.New(sess)
	fmt.Println("service created")

	params := &sns.PublishInput{
		Message: aws.String("testing 123"), 		
		PhoneNumber: aws.String("+14445556666"),	
	}
	resp, err := svc.Publish(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}