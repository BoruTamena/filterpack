package query

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type SortParam struct {
	// sort field
	Key string `json:"key,omitempty"`
}

func (st *SortParam) BindSort(c *gin.Context) (error, *SortParam) {

	err := c.BindQuery(st)

	if err != nil {
		return err, st
	}

	return nil, st

}

// sort by ascending order
func (st *SortParam) AscSort() string {

	return fmt.Sprintf(" order by %s ASC", st.Key)
}

// sort by descending order
func (st *SortParam) DscSort() string {

	return fmt.Sprintf(" order by %s DSC", st.Key)
}
