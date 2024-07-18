package query

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type SearchParam struct {
	// search struct that accept filed
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

func (s *SearchParam) BindParam(ctx *gin.Context) (error, *SearchParam) {

	err := ctx.BindQuery(&s)

	if err != nil {
		return err, s
	}
	return nil, s

}

func (s *SearchParam) SearchQuery() string {
	return fmt.Sprintf("where %v=%v", s.Key, s.Value)
}

// start with a give substring value
func (s *SearchParam) StartWith() string {
	return fmt.Sprintf(" %v like  %v%%", s.Key, s.Value)
}

// end with a give substring value
func (s *SearchParam) EndWith() string {

	return fmt.Sprintf(" %v like  %% %v", s.Key, s.Value)

}

// contain a give substring value

func (s *SearchParam) Contain() string {
	return fmt.Sprintf(" %v like  %% %v %%", s.Key, s.Value)

}
