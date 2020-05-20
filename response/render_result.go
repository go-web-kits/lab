package response

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-web-kits/dbx"
	"github.com/go-web-kits/dbx/dbx_model"
)

func Result(c *gin.Context, result dbx.Result, serializations ...dbx_model.Serialization) {
	if result.Err != nil {
		Error(c, result.Err)
		return
	}

	data, errs := dbx_model.SerializeData(result.Data, serializations...)
	for _, err := range errs {
		fmt.Println(err)
	}

	if result.Total != nil {
		Data(c, gin.H{
			"total": result.Total,
			"list":  data,
		})
	} else {
		Data(c, data)
	}
}

func Records(c *gin.Context, records interface{}, serializations ...dbx_model.Serialization) {
	data, errs := dbx_model.SerializeData(records, serializations...)
	for _, err := range errs {
		fmt.Println(err)
	}

	Data(c, data)
}

var Record = Records
