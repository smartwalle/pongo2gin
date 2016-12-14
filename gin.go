package pongo2gin

import (
	"net/http"
	"github.com/gin-gonic/gin/render"
	"github.com/smartwalle/pongo2render"
)

//	var router = gin.Default()
//	router.HTMLRender = pongo2render.NewGinRender("./templates")
//
//	router.GET("/m", func(c *gin.Context) {
//		c.HTML(200, "index.html", pongo2.Context{"aa": "eee"})
//	})
//	router.Run("localhost:9005")

type GinRender struct {
	pongo2render.Render
}

type GinHTML struct {
	*pongo2render.HTML
	data interface{}
}

func NewGinRender(templateDir string) *GinRender {
	var r = &GinRender{}
	r.TemplateDir = templateDir
	return r
}

func (this GinRender) Instance(name string, data interface{}) render.Render {
	var gHtml = &GinHTML{}
	var h = this.GetHTML(name)
	gHtml.HTML = h
	gHtml.data = data
	return gHtml
}

func (this *GinHTML) Render(w http.ResponseWriter) (err error) {
	return this.HTML.ExecuteWriter(w, this.data)
}


////////////////////////////////////////////////////////////////////////////////
const k_PONGO_TO_GIN_KEY = "pongo2gin"

type Context interface {
	Set(key string, value interface{})

	MustGet(key string) interface{}
}

func FromContext(ctx Context) *GinRender {
	return ctx.MustGet(k_PONGO_TO_GIN_KEY).(*GinRender)
}

func ToContext(ctx Context, r *GinRender) {
	ctx.Set(k_PONGO_TO_GIN_KEY, r)
}