package tpl

import (
	"../serverErrHandler"
	"../util"
	"html/template"
	"path"
)

const pageTplStr = `
{{$subItemPrefix := .SubItemPrefix}}
<!DOCTYPE html>
<html>
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
	<meta http-equiv="X-UA-Compatible" content="IE=edge"/>
	<meta name="viewport" content="initial-scale=1"/>
	<meta name="format-detection" content="telephone=no"/>
	<meta name="renderer" content="webkit"/>
	<meta name="wap-font-scale" content="no"/>
	<title>{{.Path}}</title>

	<style type="text/css">
		html, body {
			margin: 0;
			padding: 0;
			background: #fff;
		}

		html {
			font-family: "roboto_condensedbold", "Helvetica Neue", Helvetica, Arial, sans-serif;
		}

		body {
			color: #333;
			font-size: 0.625em;
			font-family: Consolas, Monaco, "Andale Mono", "DejaVu Sans Mono", monospace;
		}

		a {
			display: block;
			padding: 0.4em 0.5em;
			color: #000;
			text-decoration: none;
		}

		a:hover {
			color: #000;
			background: #f5f5f5;
		}

		input, button {
			margin: 0;
			padding: 0.25em 0;
		}

		.path-list {
			font-size: 1.5em;
			overflow: hidden;
			border-bottom: 1px #999 solid;
			zoom: 1;
		}

		.path-list a {
			position: relative;
			float: left;
			padding-right: 1.2em;
			text-align: center;
			white-space: nowrap;
			min-width: 1em;
		}

		.path-list a:after {
			content: '';
			position: absolute;
			top: 50%;
			right: 0.4em;
			width: 0.4em;
			height: 0.4em;
			border: 1px solid;
			border-color: #ccc #ccc transparent transparent;
			-webkit-transform: rotate(45deg) translateY(-50%);
			transform: rotate(45deg) translateY(-50%);
		}

		.path-list a:last-child {
			padding-right: 0.5em;
		}

		.path-list a:last-child:after {
			display: none;
		}

		.upload {
			position: relative;
			margin: 1em;
			padding: 1em;
			background: #f7f7f7;
		}

		.upload::before {
			display: none;
			content: '';
			position: absolute;
			left: 0;
			top: 0;
			right: 0;
			bottom: 0;
			opacity: 0.7;
			background: #c9c;
		}

		.upload.dragging::before {
			display: block;
		}

		.upload form {
			margin: 0;
			padding: 0;
		}

		.upload input {
			display: block;
			width: 100%;
			box-sizing: border-box;
		}

		.upload input + input {
			margin-top: 0.5em;
		}

		.archive {
			margin: 0 1em;
			overflow: hidden;
			zoom: 1;
		}

		.archive a {
			position: relative;
			float: left;
			margin: 1em 0.5em 0 0.5em;
			padding: 1em 1em 1em 3em;
			border: 2px #f5f5f5 solid;
		}

		.archive a:hover {
			border-color: #ddd;
		}

		.archive a:before {
			content: '';
			position: absolute;
			left: 1.1em;
			top: 1em;
			height: 1em;
			width: 3px;
			background: #aaa;
		}

		.archive a:after {
			content: '';
			position: absolute;
			left: 0.6em;
			top: 1.1em;
			width: 0.5em;
			height: 0.5em;
			margin-left: 1px;
			border: 3px #aaa solid;
			border-top-color: transparent;
			border-left-color: transparent;
			-webkit-transform: rotate(45deg);
			transform: rotate(45deg);
		}

		.item-list {
			margin: 1em;
		}

		.item-list a {
			display: flex;
			flex-flow: row nowrap;
			align-items: center;
			border-bottom: 1px #f5f5f5 solid;
		}

		.item-list span {
			margin-left: 1em;
			flex-shrink: 0;
		}

		.item-list .name {
			flex-grow: 1;
			flex-shrink: 1;
			flex-basis: 0;
			margin-left: 0;
			font-size: 1.5em;
			word-break: break-all;
		}

		.item-list .size {
			white-space: nowrap;
			text-align: right;
			color: #666;
		}

		.item-list .time {
			width: 10em;
			color: #999;
			text-align: right;
			white-space: nowrap;
			overflow: hidden;
			float: right;
		}

		.error {
			margin: 1em;
			padding: 1em;
			background: #ffc;
		}

		@media only screen and (max-width: 350px) {
			.item-list .time {
				display: none;
			}
		}
	</style>
</head>
<body>

<div class="path-list">
    {{range .Paths}}
		<a href="{{.Path}}">{{html .Name}}</a>
    {{end}}
</div>

{{if .CanUpload}}
	<div class="upload">
		<form method="POST" enctype="multipart/form-data">
			<input type="file" name="files" class="files" multiple="multiple" accept="*/*"/>
			<input type="submit" value="Upload"/>
		</form>
	</div>
{{end}}

{{if .CanArchive}}
	<div class="archive">
		<a href="{{$subItemPrefix}}?tar" download="{{.ItemName}}.tar">.tar</a>
		<a href="{{$subItemPrefix}}?tgz" download="{{.ItemName}}.tar.gz">.tar.gz</a>
		<a href="{{$subItemPrefix}}?zip" download="{{.ItemName}}.zip">.zip</a>
	</div>
{{end}}

<div class="item-list">
	<a href="{{if .IsRoot}}./{{else}}../{{end}}">
		<span class="name">../</span>
		<span class="size"></span>
		<span class="time"></span>
	</a>
    {{range .SubItems}}
        {{$isDir := .IsDir}}
		<a href="{{$subItemPrefix}}{{.Name}}{{if $isDir}}/{{end}}"
		   class="item {{if $isDir}}item-dir{{else}}item-file{{end}}">
			<span class="name">{{html .Name}}{{if $isDir}}/{{end}}</span>
			<span class="size">{{if not $isDir}}{{fmtSize .Size}}{{end}}</span>
			<span class="time">{{fmtTime .ModTime}}</span>
		</a>
    {{end}}
</div>

{{range .Errors}}
	<div class="error">{{.}}</div>
{{end}}

<script type="text/javascript">
	(function () {
		if (!document.querySelector) {
			return;
		}

		var upload = document.querySelector('.upload');
		if (!upload || !upload.addEventListener) {
			return;
		}
		var fileInput = upload.querySelector('.files');

		var addClass = function (ele, className) {
			ele && ele.classList && ele.classList.add(className);
		};

		var removeClass = function (ele, className) {
			ele && ele.classList && ele.classList.remove(className);
		};

		var onDragEnterOver = function (e) {
			e.stopPropagation();
			e.preventDefault();
			addClass(e.currentTarget, 'dragging');
		};

		var onDragLeave = function (e) {
			if (e.target === e.currentTarget) {
				removeClass(e.currentTarget, 'dragging');
			}
		};

		var onDrop = function (e) {
			e.stopPropagation();
			e.preventDefault();
			removeClass(e.currentTarget, 'dragging');

			if (!e.dataTransfer.files) {
				return;
			}
			fileInput.files = e.dataTransfer.files;
		};

		upload.addEventListener('dragenter', onDragEnterOver);
		upload.addEventListener('dragover', onDragEnterOver);
		upload.addEventListener('dragleave', onDragLeave);
		upload.addEventListener('drop', onDrop);
	})();
</script>
</body>
</html>
`

var defaultPage *template.Template

func init() {
	tplObj := template.New("page")
	tplObj = addFuncMap(tplObj)

	var err error
	defaultPage, err = tplObj.Parse(pageTplStr)
	if serverErrHandler.CheckError(err) {
		defaultPage = template.Must(tplObj.Parse("Builtin Template Error"))
	}
}

func LoadPage(tplPath string) (*template.Template, error) {
	var tplObj *template.Template
	var err error

	if len(tplPath) > 0 {
		tplObj = template.New(path.Base(tplPath))
		tplObj = addFuncMap(tplObj)
		tplObj, err = tplObj.ParseFiles(tplPath)
	}
	if err != nil || len(tplPath) == 0 {
		tplObj = defaultPage
	}

	return tplObj, err
}

func addFuncMap(tpl *template.Template) *template.Template {
	return tpl.Funcs(template.FuncMap{
		"fmtSize": util.FormatSize,
		"fmtTime": util.FormatTimeMinute,
	})
}
