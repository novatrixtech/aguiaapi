package model

type ProdutoServico struct {
	ID   int
	Nome string
}

type ProdutoServicoInformacao struct {
	ProdutoServico ProdutoServico
	MaiorValor     Contrato
	Media          float64
	MenorValor     Contrato
}
