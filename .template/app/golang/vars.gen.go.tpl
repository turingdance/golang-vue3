package model

import "turingdance.com/turing/internal/pkg/utils"

var modelArray []interface{} = make([]interface{}, 0)

func RegisterModel(m interface{}) {
	modelArray = append(modelArray, m)
}
func AllRegistedModel() []interface{} {
	return modelArray
}

func MakePKID() string {
	return utils.PKID()
}