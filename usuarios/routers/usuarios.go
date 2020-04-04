package routers

import (
	"github.com/ValdemiRamos/usuarios/controlers"
	"github.com/labstack/echo"
)

//App é uma instancia de Echo
var App *echo.Echo

func init() {

	App = echo.New()
	// Rota pagina incicial

	App.GET("/", controlers.Home)

	api := App.Group("/v1")

	api.POST("/usuarios", controlers.Insert)

}
