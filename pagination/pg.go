package pagination

import (
	"errors"

	"github.com/BoruTamena/query"
	"github.com/gin-gonic/gin"
)

var (
	c = &gin.Context{}
)

type PaginationResponse struct {
	Query string
}

func Paginate() (error, PaginationResponse) {

	var q query.Pagination

	err, qt := q.BindPagination(c)

	if err != nil {
		return err, PaginationResponse{}
	}

	if qt.Page != 0 && qt.PageSize != 0 {

		return nil, PaginationResponse{
			Query: q.Pagenate(""),
		}

	}

	return errors.New("please specify your page and page size "), PaginationResponse{}

}

func PaginateQuery(qs string) (error, string) {

	var q query.Pagination

	err, qt := q.BindPagination(c)

	if err != nil {
		return err, ""
	}

	if qt.Page != 0 && qt.PageSize != 0 {

		return nil, q.Pagenate(qs)

	}

	return errors.New("no parameters specified"), ""
}
