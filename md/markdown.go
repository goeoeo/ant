package md

import (
	"fmt"
	"github.com/russross/blackfriday/v2"
)

func MdToHtmlWithCss(input []byte) (out string) {
	var (
		head string
		body []byte
		foot string
	)
	head = `
<html>
	<head>
		<meta charset="utf-8">
		<link href="https://cdn.bootcdn.net/ajax/libs/github-markdown-css/4.0.0/github-markdown.css" rel="stylesheet">
	</head>
	<style>
		html {
			padding: 0 0 0 300px;
		}
		.markdown-body {
			box-sizing: border-box;
			min-width: 500px;
			max-width: 1200px;
			padding: 45px;
		}
	
		@media (max-width: 767px) {
			.markdown-body {
				padding: 15px;
			}
		}
		nav {
			width: 275px;
			position: fixed;
			left: 20px;
			top: 10px;
			bottom: 20px;
			overflow-y: auto;
			background: #fff;
		}
		nav::-webkit-scrollbar {/*滚动条整体样式*/
		
			width: 4px;
			height: 1px;
			background-color: gainsboro;
		
        }
	</style>
	<body class="markdown-body">
`
	foot = `

	<body>
</html>	
`

	r := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
		Flags: blackfriday.CommonHTMLFlags | blackfriday.TOC,
	})

	body = blackfriday.Run(input, blackfriday.WithRenderer(r))

	return head + string(body) + foot
}

func MdToHtmlWithCssNoDir(input []byte) (out string) {
	outTmp := []byte(`
<html>
	<head>
		<meta charset="utf-8">
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
