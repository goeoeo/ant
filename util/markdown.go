package util

import (
	"fmt"
	"github.com/russross/blackfriday/v2"
)

func MdToHtmlWithCss(input []byte) (out string) {
	outTmp := []byte(`
<html>
	<head>
		<link href="https://cdn.bootcdn.net/ajax/libs/github-markdown-css/4.0.0/github-markdown.css" rel="stylesheet">
	</head>
	<style>
		.markdown-body {
			box-sizing: border-box;
			min-width: 200px;
			max-width: 980px;
			margin: 0 auto;
			padding: 45px;
		}
	
		@media (max-width: 767px) {
			.markdown-body {
				padding: 15px;
			}
		}
	</style>
	<body class="markdown-body">
	%s
	<body>
</html>	
`)

	return fmt.Sprintf(string(outTmp), blackfriday.Run(input))
}
