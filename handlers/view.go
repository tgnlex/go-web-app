package handlers

import (
	"html/template"
	"net/http"

	"example.com/project/utils"
)

var newTemplate = utils.CreateTemplate

var index = newTemplate("view/index.html")
var login = newTemplate("view/login.html")

type Res = http.ResponseWriter
type Req = *http.Request
type View = *template.Template
