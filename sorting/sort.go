package sorting

import (
	"errors"

	"github.com/BoruTamena/query"
	"github.com/gin-gonic/gin"
)

var (
	c = &gin.Context{}
)

type SortResponseQuery struct {
	Asc string
	Dsc string
}

func Sort() (error, SortResponseQuery) {

	var st query.SortParam

	err, sp := st.BindSort(c)

	if err != nil {

		return err, SortResponseQuery{}
	}

	if sp.Key != "" {
		return errors.New("please specify a filed name "), SortResponseQuery{}
	}

	res := SortResponseQuery{
		Asc: sp.AscSort(),
		Dsc: sp.DscSort(),
	}

	return nil, res

}
