package controlers

import (
	"fmt"
	"net/http"

	"github.com/ValdemiRamos/usuarios/models"
	"github.com/labstack/echo"
)

//Home é a pagina principal da aplicação
func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Olá Valdemi...")
}

//Insert é uma função que insere registros
func Insert(c echo.Context) error {
	nome := c.FormValue("nome")
	email := c.FormValue("email")

	var usuario models.Usuarios
	usuario.Nome = nome
	usuario.Email = email

	fmt.Println(nome, email)

	if nome != "" && email != "" {
		if _, err := models.UsuarioModel.Insert(usuario); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"mensagem": "Não foi possível gravar os registros no banco de dados! Tente Novamente..",
			})
		}

		return c.JSON(http.StatusCreated, map[string]string{
			"menssagem": "Registro gravado com Sucesso!!",
		})

	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"menssagem": "Os campos precisam ser passados!",
	})

}
