package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/go-web-kits/lab/response"
	. "github.com/go-web-kits/testx"
	. "github.com/go-web-kits/testx/api_matchers"
	"github.com/go-web-kits/utils"
	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = Describe("Response", func() {
	BeforeEach(func() {
		CurrentAPI = utils.GetFuncName(Handler)
	})

	Describe("Success", func() {
		It("responses success response", func() {
			Action = func(c *gin.Context) {
				Success(c)
			}
			IsExpected().To(ResponseSuccess())
		})
	})

	Describe("Only", func() {
		It("responses the given data", func() {
			Action = func(c *gin.Context) {
				Only(c, gin.H{"msg": "hello"})
			}
			ExpectRequested().To(Response(http.StatusOK))
			ExpectRequested().To(Response(map[string]interface{}{"msg": "hello"}))
		})
	})

	Describe("Data", func() {
		It("responses success with data", func() {
			Action = func(c *gin.Context) {
				Data(c, gin.H{"msg": "hello"})
			}
			IsExpected().To(ResponseSuccess())
			ExpectRequested().To(ResponseData(map[string]interface{}{"msg": "hello"}))
		})
	})
})
