package cfni

import "fmt"

type S3AccessKey struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
}

func (s *S3AccessKey) IsValid() bool {
	return s.AccessKeyID != "" && s.SecretAccessKey != ""
}

func s3Client(key *S3AccessKey) string {
	if key != nil && key.AccessKeyID != "" && key.SecretAccessKey != "" && key.SessionToken != "" {
		return fmt.Sprintf(`s3_client = boto3.session.Session(aws_access_key_id="%s", aws_secret_access_key="%s", aws_session_token="%s").client("s3", region_name="us-east-1")`, key.AccessKeyID, key.SecretAccessKey, key.SessionToken)
	}

	if key != nil && key.AccessKeyID != "" && key.SecretAccessKey != "" {
		return fmt.Sprintf(`s3_client = boto3.session.Session(aws_access_key_id="%s", aws_secret_access_key="%s").client("s3", region_name="us-east-1")`, key.AccessKeyID, key.SecretAccessKey)
	}

	return `s3_client = boto3.client("s3")`
}
