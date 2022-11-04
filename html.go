package pongo2gin

import (
	"github.com/gin-gonic/gin/render"
	"github.com/smartwalle/pongo2render"
	"net/http"
)

//	var router = gin.Default()
//	router.HTMLRender = pongo2gin.NewHTMLRender("./templates")
//
//	router.GET("/m", func(c *gin.Context) {
//		c.HTML(200, "index.html", pongo2.Context{"key": "value"})
//	})
//	router.Run("localhost:9005")

type HTMLRender struct {
	pongo2render.Render
}

func NewHTMLRender(templateDir string) *HTMLRender {
	var r = &HTMLRender{}
	r.TemplateDir = templateDir
	return r
}

func (this *HTMLRender) Instance(name string, data interface{}) render.Render {
	var html = &HTML{}
	html.Template = this.Template(name)
	html.data = data
	return html
}

func (this *HTMLRender) InstanceFromString(tpl string, data interface{}) render.Render {
	var html = &HTML{}
	html.Template = this.TemplateFromString(tpl)
	html.data = data
	return html
}

type HTML struct {
	*pongo2render.Template
	data interface{}
}

func (this *HTML) Render(w http.ResponseWriter) (err error) {
	return this.Template.ExecuteWriter(w, this.data)
}

func (this *HTML) WriteContentType(w http.ResponseWriter) {
	pongo2render.WriteContentType(w, []string{"text/html; charset=utf-8"})
}

const key = "_pongo2gin_"

type Context interface {
	Set(key string, value interface{})

	MustGet(key string) interface{}
}

func FromContext(ctx Context) *HTMLRender {
	return ctx.MustGet(key).(*HTMLRender)
}

func ToContext(ctx Context, r *HTMLRender) {
	ctx.Set(key, r)
}
