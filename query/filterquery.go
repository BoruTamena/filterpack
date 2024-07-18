package query

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

// filter string /int types using general equal operation
func (f FilterParam) Filter() string {

	return fmt.Sprintf("where %v =%v ", f.Key, f.Value)

}

// filter int value which is greater than specified value
func (f FilterParam) FilterGt() string {

	return fmt.Sprintf("where %v > %v", f.Key, f.Value)

}

// filter int value which is greater than or equal to  specified value
func (f FilterParam) FilterGtEq() string {

	return fmt.Sprintf("where %v >=%v", f.Key, f.Value)

}

// // filter int value which is less than  specified value
func (f FilterParam) FilterLt() string {
	return fmt.Sprintf("where %v < %v", f.Key, f.Value)
}

// filter int value which ia less than or equal to  specified value
func (f FilterParam) FilterLtEq() string {
	return fmt.Sprintf("where %v >= %v ", f.Key, f.Value)
}

// filter  string value which is start with the   specified  substring value
func (f FilterParam) StartWith() string {
	return fmt.Sprintf(" %s like  %s%%", f.Key, f.Value)
}

// filter  string value which is end with the   specified  substring value
func (f FilterParam) EndWith() string {
	return fmt.Sprintf(" %s like  %%%s", f.Key, f.Value)
}

// filter  string value which is contains with the   specified  substring value
func (f FilterParam) Contain() string {
	return fmt.Sprintf(" %s like  %%%s%%", f.Key, f.Value)
}
