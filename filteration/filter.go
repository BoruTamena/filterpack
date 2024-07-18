package filteration

import (
	"errors"
	"strconv"

	"github.com/BoruTamena/query"
	"github.com/gin-gonic/gin"
)

var (
	c = &gin.Context{}
)

type FilterResponseQuery struct {
	Query     string
	Lt        string
	LtEq      string
	Gt        string
	GtEq      string
	StartWith string
	EndWith   string
	Contain   string
}

func Filter() (error, FilterResponseQuery) {

	var f query.FilterParam

	err, fp := f.BindFilterParam(c)

	if err != nil {
		return err, FilterResponseQuery{}
	}

	if fp.Key != "" || fp.Op != "" || fp.Value != "" {

		_, err := strconv.Atoi(fp.Value)

		if err != nil {
			// perform string operation only

			return nil, FilterResponseQuery{

				Query:     f.Filter(),
				StartWith: f.StartWith(),
				EndWith:   f.EndWith(),
				Contain:   f.Contain(),
			}
		}

		res := FilterResponseQuery{
			Query: f.Filter(),
			Lt:    f.FilterLt(),
			LtEq:  f.FilterLtEq(),
			Gt:    f.FilterGt(),
			GtEq:  f.FilterGtEq(),
		}

		return nil, res

	}

	return errors.New("please specify your filter parameters"), FilterResponseQuery{}

}
