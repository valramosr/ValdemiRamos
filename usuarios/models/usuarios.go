package models

import (
	"github.com/ValdemiRamos/usuarios/lib"
)

// representa a tabela de usuários na base de dados
type Usuarios struct {
	ID    int    `db:"id" json:"id"`
	Nome  string `db:"nome" json:"nome"`
	Email string `db:"email" json:"email"`
}

var UsuarioModel = lib.Sess.Collection("usuarios")

