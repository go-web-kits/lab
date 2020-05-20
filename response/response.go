/*
Render JSON response in a unified format

{
	"result": {
		"code": n,
		"message": "xx"
	},
	"data": {
		...
	}
}
*/
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Structure struct {
	Result interface{} `json:"result"`
	Data   interface{} `json:"data"`
}

// Success success with empty data
func Success(c *gin.Context) {
	JSON(c, http.StatusOK, &Structure{
		Result: &result{0, "success"},
		Data:   struct{}{},
	})
}

// Data success with data
func Data(c *gin.Context, data interface{}) {
	JSON(c, http.StatusOK, &Structure{
		Result: &result{0, "success"},
		Data:   data,
	})
}

func Only(c *gin.Context, data interface{}) {
	JSON(c, http.StatusOK, data)
}

func JSON(c *gin.Context, status int, data interface{}) {
	c.Set("response_body", data)
	c.JSON(status, data)
}
