package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero"`
	CPF  string `json:"cpf" validate:"len=9"`
	RG   string `json:"rg" validate:"len=11"`
}

func ValidaDadosDeAluno(aluno *Aluno) error {
	if err := validator.validate(aluno); err != nil {
		return err
	}
	return nil
}
