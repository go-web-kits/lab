package response

import (
	"github.com/gin-gonic/gin"
	"github.com/go-web-kits/dbx"
)

func ResultOf(c *gin.Context, result interface{}) {
	switch r := result.(type) {
	case dbx.Result:
		if r.Err != nil {
			Error(c, r.Err)
			return
		}
	case error:
		if r != nil {
			Error(c, r)
			return
		}
	}

	Success(c)
}

var SuccessOrNot = ResultOf
