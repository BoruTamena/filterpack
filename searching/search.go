package searching

import (
	"errors"

	"github.com/BoruTamena/query"
	"github.com/gin-gonic/gin"
)

var (
	c = &gin.Context{}
)

type SearchResponseQuery struct {
	Query     string
	StartWith string
	EndWith   string
	Contain   string
}

func Search() (error, SearchResponseQuery) {

	var s query.SearchParam

	err, param := s.BindParam(c)

	if err != nil {

		return err, SearchResponseQuery{}
	}

	if param.Key != "" && param.Value != "" {

		res := SearchResponseQuery{
			Query:     s.SearchQuery(),
			StartWith: s.StartWith(),
			EndWith:   s.EndWith(),
			Contain:   s.Contain(),
		}
		return nil, res

	}

	return errors.New("please specify your query parmateries "), SearchResponseQuery{}

}
