package pongo2gin

import (
	"net/http"
	"github.com/gin-gonic/gin/render"
	"github.com/smartwalle/pongo2render"
)

//	var router = gin.Default()
//	router.HTMLRender = pongo2gin.NewHTMLRender("./templates")
//
//	router.GET("/m", func(c *gin.Context) {
//		c.HTML(200, "index.html", pongo2.Context{"key": "value"})
//	})
//	router.Run("localhost:9005")

// --------------------------------------------------------------------------------
type HTMLRender struct {
	pongo2render.Render
}

func NewHTMLRender(templateDir string) *HTMLRender {
	var r = &HTMLRender{}
	r.TemplateDir = templateDir
	return r
}

func (this *HTMLRender) Instance(name string, data interface{}) render.Render {
	var gHtml = &Template{}
	var h = this.Template(name)
	gHtml.Template = h
	gHtml.data = data
	return gHtml
}

func (this *HTMLRender) InstanceFromString(tpl string, data interface{}) render.Render {
	var gHtml = &Template{}
	var h = this.TemplateFromString(tpl)
	gHtml.Template = h
	gHtml.data = data
	return gHtml
}

// --------------------------------------------------------------------------------
type Template struct {
	*pongo2render.Template
	data interface{}
}

func (this *Template) Render(w http.ResponseWriter) (err error) {
	return this.Template.ExecuteWriter(w, this.data)
}

func (this *Template) WriteContentType(w http.ResponseWriter) {
	pongo2render.WriteContentType(w, []string{"text/html; charset=utf-8"})
}

// --------------------------------------------------------------------------------
const k_PONGO_TO_GIN_KEY = "pongo2gin"

type Context interface {
	Set(key string, value interface{})

	MustGet(key string) interface{}
}

func FromContext(ctx Context) *HTMLRender {
	return ctx.MustGet(k_PONGO_TO_GIN_KEY).(*HTMLRender)
}

func ToContext(ctx Context, r *HTMLRender) {
	ctx.Set(k_PONGO_TO_GIN_KEY, r)
}