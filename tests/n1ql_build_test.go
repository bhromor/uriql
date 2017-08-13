package tests

import (
	decoder "github.com/kite-social/uriql"
	"github.com/kite-social/uriql/dictionary"
	"github.com/kite-social/uriql/models"
	search "github.com/kite-social/uriql/qbuilder"
	"testing"
)

func printResult(t *testing.T, p string, qp interface{}, q string) {
	t.Logf("Decoding : %s", p)
	t.Logf("Decoded to : %+v", qp)
	t.Logf("Query %s", q)
}

func TestN1QLBuild(t *testing.T) {

	dict := &models.Dictionary{Model: dictionary.N1QLDictionary()}

	decode := decoder.GetQueryDecoder(dict)
	builder := search.GetN1QLQueryBuilder("kite")

	t.Log("Testing Restaurant Parameter : ")

	p := "restaurant?title=mr.burger"
	qp := decode.DecodeQueryString(models.RequestInfo{UserID: "1234567890", Type: "restaurant", Query: p})
	q := builder.Build(qp)
	printResult(t, p, qp, q)

	p = "restaurant?address=dhaka"
	qp = decode.DecodeQueryString(models.RequestInfo{UserID: "1234567890", Type: "restaurant", Query: p})
	q = builder.Build(qp)
	printResult(t, p, qp, q)

	/*decode = search.GetQueryDecoder(dictionary.RestaurantItemsDictionary())
	builder = builder.GetN1QLQueryBuilder()

	t.Log("Testing Restaurant Items Parameter : ")


	p = "restaurant-items?foods=burger"
	qp = decode.DecodeQueryString(p)
	q = builder.Build(qp)
	printResult(t, p, qp, q)

	p = "restaurant-items?foods-price=45"
	qp = decode.DecodeQueryString(p)
	q = builder.Build(qp)
	printResult(t, p, qp, q)*/

}
