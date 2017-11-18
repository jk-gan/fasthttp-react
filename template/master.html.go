// Code generated by hero.
// source: C:\Go\dev\src\fasthttp\template\master.html
// DO NOT EDIT!
package template

import (
	"bytes"
	"encoding/json"
	"fasthttp/model"

	"github.com/shiyanhui/hero"
)

func Main(merchant *model.Merchant, buffer *bytes.Buffer) {
	buffer.WriteString(`
<!DOCTYPE html>
<html lang="en">

<head>
    <title>`)
	hero.EscapeHTML(merchant.Name, buffer)
	buffer.WriteString(`</title>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1, user-scalable=0">
    <meta name="description" content="">
    <meta name="author" content="">
    <meta property="og:title" content="`)
	hero.EscapeHTML(merchant.Name, buffer)
	buffer.WriteString(`" />
    <meta property="og:url" content="`)
	hero.EscapeHTML(merchant.LogoURL, buffer)
	buffer.WriteString(`" />
    <meta property="og:description" content="`)
	hero.EscapeHTML(merchant.Name, buffer)
	buffer.WriteString(`" />
    <meta property="og:image" content="`)
	hero.EscapeHTML(merchant.LogoURL, buffer)
	buffer.WriteString(`" />
</head>

`)

	data, _ := json.Marshal(merchant)

	buffer.WriteString(`
<body>
    <script type="text/javascript">
        window.App = {
            merchant: `)
	buffer.WriteString(string(data))
	buffer.WriteString(`,
        };
    </script>
    <script type="text/javascript">
        (function (d, s, id) {
            var js, fjs = d.getElementsByTagName(s)[0];
            if (d.getElementById(id)) { return; }
            js = d.createElement(s); js.id = id;
            js.src = "//connect.facebook.com/en_US/messenger.Extensions.js";
            fjs.parentNode.insertBefore(js, fjs);
        }(document, 'script', 'Messenger'));
    </script>
    <script src="https://res.wx.qq.com/open/js/jweixin-1.0.0.js"></script>
    <div id="app"></div>
</body>

</html>`)

}