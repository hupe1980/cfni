package cfni

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNodeJSInlineFunction(t *testing.T) {
	input := `async function cfni(event, context) {
	console.log(event);
	return {};
}`
	expected := []string{
		`const response = require('cfn-response');`,
		"async function cfni(e){return console.log(e),{}};",
		"",
		"exports.handler = async function(event, context) {",
		"	let responseData = {};",
		"	try {",
		"		responseData = await cfni(event, context);",
		"	} finally {",
		"		await response.send(event, context, response.SUCCESS, responseData, context.logStreamName, true);",
		"	}",
		"}",
	}
	assert.ElementsMatch(t, expected, createNodeJSInlineFunction(input))
}
