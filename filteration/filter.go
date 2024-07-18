package filteration

import "github.com/gin-gonic/gin"

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
