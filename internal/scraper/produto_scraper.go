package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/ygorfelipe/busca-preco-api-go.git/internal/domain/entities"
)

type ProdutoScraper struct{}

func (s ProdutoScraper) Scrape(
	document *goquery.Document,
) []entities.ProdutoEntity {

	var produtos []entities.ProdutoEntity

	document.Find(".produto").Each(
		func(i int, element *goquery.Selection) {

			produto := entities.ProdutoEntity{
				Nome:  element.Find(".nome").Text(),
				Preco: 0,
				Image: "",
			}

			produtos = append(produtos, produto)
		},
	)

	return produtos
}
