package docs

import (
	"github.com/labstack/echo/v4"
	"html/template"
)

type graphiql struct {
	port int32
}

func NewGraphiql(port int32) *graphiql {
	return &graphiql{
		port: port,
	}
}

func (g *graphiql) Start() echo.HandlerFunc {
	return func(c echo.Context) error {
		t := template.Must(template.New("graphql").Parse(`
  <!DOCTYPE html>
  <html>
       <body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
<div id="sandbox" style="position:absolute;top:0;right:0;bottom:0;left:0"></div>
<script src="https://embeddable-sandbox.cdn.apollographql.com/_latest/embeddable-sandbox.umd.production.min.js"></script>
<script>
 new window.EmbeddedSandbox({
   target: "#sandbox",
   // Pass through your server href if you are embedding on an endpoint.
   // Otherwise, you can pass whatever endpoint you want Sandbox to start up with here.
   initialEndpoint: window.location.href,
 });
 // advanced options: https://www.apollographql.com/docs/studio/explorer/sandbox#embedding-sandbox
</script>

       </body>
  </html>
  `))
		t.Execute(c.Response().Writer, &g.port)
		return nil
	}
}
