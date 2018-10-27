package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.baidu.com/")
	if err != nil {
		return
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

/*

return;
<html>
<head>
	<script>
		location.replace(location.href.replace("https://","http://"));
	</script>
</head>
<body>
	<noscript><meta http-equiv="refresh" content="0;url=http://www.baidu.com/"></noscript>
</body>
</html>
*/
