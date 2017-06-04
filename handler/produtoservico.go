package handler

import (
	"net/http"

	"github.com/novatrixtech/aguiaapi/lib/context"
	"github.com/novatrixtech/aguiaapi/model"
)

var ProdutosServicos map[int]model.ProdutoServico
var ProdutosServicosInformacoes map[int]model.ProdutoServicoInformacao
var ps model.ProdutoServico

func init() {
	ProdutosServicos = make(map[int]model.ProdutoServico, 0)
	ps = model.ProdutoServico{}
	ps.ID = 1
	ps.Nome = "locacao de veiculos"
	ProdutosServicos[ps.ID] = ps
	ps = model.ProdutoServico{}
	ps.ID = 2
	ps.Nome = "prestacao de servico de limpeza"
	ProdutosServicos[ps.ID] = ps
}

func GetProdutoServico(ctx *context.Context) {

	var retorno []model.ProdutoServico
	for _, ps := range ProdutosServicos {
		retorno = append(retorno, ps)
	}

	ctx.JSON(http.StatusOK, retorno)

}
