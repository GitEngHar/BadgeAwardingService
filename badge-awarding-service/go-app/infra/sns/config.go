package sns

import (
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type Config struct {
	topicArn string
	client   *sns.Client
}
