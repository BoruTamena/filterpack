package searching

import (
	"errors"

	"github.com/BoruTamena/query"
	"github.com/gin-gonic/gin"
)

var (
	c = &gin.Context{}
)

func Search() (error, string) {

	var s query.SearchParam

	err, param := s.BindParam(c)

	if err != nil {

		return err, ""
	}

	if param.Key != "" && param.Value != "" {

		res := s.SearchQuery()
		return nil, res

	}

	return errors.New("please specify your query parmateries "), ""

}
