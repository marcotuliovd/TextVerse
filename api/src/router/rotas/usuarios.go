package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI: "/usuarios",
		Metodo: http.MethodPost,
		Funcao: controllers.CreateUser,
		Auth: false,
	},
	{
		URI: "/usuarios",
		Metodo: http.MethodGet,
		Funcao: controllers.FindUser,
		Auth: false,
	},
	{
		URI: "/usuarios/{usuarioId}",
		Metodo: http.MethodGet,
		Funcao: controllers.FindUserbyId,
		Auth: false,
	},
	{
		URI: "/usuarios/{usuarioId}",
		Metodo: http.MethodPut,
		Funcao: controllers.UpdateUser,
		Auth: false,
	},
	{
		URI: "/usuarios/{usuarioId}",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeleteUser,
		Auth: false,
	},
}