package filteration

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type FilterParam struct {
	Key   string `json:"key,omitempty"`
	Op    string `json:"op,omitempty"`
	Value string `json:"value,omitempty"`
}

func (f *FilterParam) BindFilterParam(c *gin.Context) (error, *FilterParam) {

	err := c.BindQuery(f)

	if err != nil {
		return err, f
	}

	return nil, f

}

func (f FilterParam) Filter() string {

}

func (f FilterParam) FilterGt() string {

	return fmt.Sprintf("where %v > $ %v", f.Key, f.Value)

}

func (f FilterParam) FilterGtEq() string {

	return fmt.Sprintf("where %v >= $ %v", f.Key, f.Value)

}

func (f FilterParam) FilterLt() string {
	return fmt.Sprintf("where %v < $ %v", f.Key, f.Value)
}

func (f FilterParam) FilterLtEq() string {
	return fmt.Sprintf("where %v >= $ %v", f.Key, f.Value)
}
