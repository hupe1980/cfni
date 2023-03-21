package cfni

import (
	"fmt"
	"strings"

	"github.com/hupe1980/cfni/pkg/obfuscator/js"
)

type CreateCFNCodeExecutionOptions struct {
	Code            string
	Runtime         string
	LogicalRoleID   string
	LogicalLambdaID string
	LogicalCustomID string
	CustomType      string
	S3AccessKey     *S3AccessKey
}

func (c *CFNI) CreateCFNCodeExecutionHandler(opts *CreateCFNCodeExecutionOptions) ([]byte, error) {
	handler := "index.lambda_handler"
	if strings.HasPrefix(opts.Runtime, "nodejs") {
		handler = "index.handler"
	}

	type data struct {
		Code            string
		Handler         string
		Runtime         string
		LogicalRoleID   string
		LogicalLambdaID string
		LogicalCustomID string
		CustomType      string
	}

	return c.createHandler(&HandlerOptions{
		CFNITemplate: "templates/cfn_code_execution.py",
		CFNIData: &data{
			Code:            toPythonList(createNodeJSInlineFunction(opts.Code)),
			Handler:         handler,
			Runtime:         opts.Runtime,
			LogicalRoleID:   opts.LogicalRoleID,
			LogicalLambdaID: opts.LogicalLambdaID,
			LogicalCustomID: opts.LogicalCustomID,
			CustomType:      opts.CustomType,
		},
		S3Client: s3Client(opts.S3AccessKey),
	})
}

func createNodeJSInlineFunction(code string) []string {
	obfuscator := js.New()
	code, _ = obfuscator.Obfuscate(code)

	inlineCode := fmt.Sprintf(`const response = require('cfn-response');
%s

exports.handler = async function(event, context) {
	let responseData = {};
	try {
		responseData = await cfni(event, context);
	} finally {
		await response.send(event, context, response.SUCCESS, responseData, context.logStreamName, true);
	}
}`, code)

	return strings.Split(strings.Replace(inlineCode, `"`, `\"`, -1), "\n")
}
