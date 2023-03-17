package cfni

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestS3Client(t *testing.T) {
	t.Run("no-key", func(t *testing.T) {
		expected := `s3_client = boto3.client("s3")`
		assert.Equal(t, expected, s3Client(nil))
	})

	t.Run("with access-key", func(t *testing.T) {
		expected := `s3_client = boto3.session.Session(aws_access_key_id="ABC", aws_secret_access_key="Dummy").client("s3", region_name="us-east-1")`
		assert.Equal(t, expected, s3Client(&S3AccessKey{
			AccessKeyID:     "ABC",
			SecretAccessKey: "Dummy",
		}))
	})

	t.Run("with session-token", func(t *testing.T) {
		expected := `s3_client = boto3.session.Session(aws_access_key_id="ABC", aws_secret_access_key="Dummy", aws_session_token="Token").client("s3", region_name="us-east-1")`
		assert.Equal(t, expected, s3Client(&S3AccessKey{
			AccessKeyID:     "ABC",
			SecretAccessKey: "Dummy",
			SessionToken:    "Token",
		}))
	})
}
