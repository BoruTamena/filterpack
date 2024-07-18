package query

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

func (p *Pagination) BindPagination(c *gin.Context) (error, *Pagination) {

	err := c.BindQuery(p)

	if err != nil {
		return err, p
	}

	return nil, p
}

func (p *Pagination) Pagenate(q string) string {

	limit := p.PageSize
	offset := (p.Page - 1) * p.PageSize

	if q != "" {
		return q + fmt.Sprintf("limit %v offset %v", limit, offset)
	}

	return fmt.Sprintf("limit %v offset %v", limit, offset)

}
