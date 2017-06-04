package model

import "gopkg.in/mgo.v2/bson"

type Contrato struct {
	ID                     bson.ObjectId `bson:"_id"`
	Orgao                  string        `json:"orgao" bson:"orgao"`
	Fornecedor             string        `json:"fornecedor" bson:"Fornecedor"`
	CNPJ                   string        `json:"cNPJ" bson:"CNPJ"`
	Objeto                 string        `json:"objeto" bson:"Objeto"`
	Valor                  float64       `json:"valor" bson:"Valor"`
	Modalidade             string        `json:"modalidade" bson:"Modalidade"`
	Contrato               string        `json:"contrato" bson:"Contrato"`
	DataDeAssinatura       string        `json:"dataDeAssinatura" bson:"dataDeAssinatura"`
	Vigencia               int           `json:"vigencia" bson:"vigencia"`
	ProcessoAdministrativo string        `json:"processoAdministrativo" bson:"ProcessoAdministrativo"`
	Licitacao              string        `json:"licitacao" bson:"licitacao"`
	Evento                 string        `json:"evento" bson:"Evento"`
	DataDaPublicacao       string        `json:"dataDaPublicacao" bson:"DataDaPublicacao"`
}
