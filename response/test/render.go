package test

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/go-web-kits/dbx"
	"github.com/go-web-kits/dbx/dbx_model"
	be "github.com/go-web-kits/lab/business_error"
	. "github.com/go-web-kits/lab/response"
	. "github.com/go-web-kits/testx"
	. "github.com/go-web-kits/testx/api_matchers"
	"github.com/go-web-kits/testx/factory"
	"github.com/go-web-kits/utils"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v8"
)

var _ = Describe("Render", func() {
	var (
		post Post
		p    *MonkeyPatches
	)

	BeforeEach(func() {
		CurrentAPI = utils.GetFuncName(Handler)
	})

	AfterEach(func() {
		CleanData(&Post{})
		Reset(&post)
		p.Check()
	})

	Describe("Error", func() {
		When("giving Renderable error", func() {
			It("calls .Render", func() {
				Action = func(c *gin.Context) {
					Error(c, be.CommonErrors[be.Unauthorized])
				}

				IsExpected().To(Response(be.CommonErrors[be.Unauthorized]))
			})
		})
		When("giving validator.ValidationErrors", func() {
			It("renders params error", func() {
				Action = func(c *gin.Context) {
					Error(c, validator.ValidationErrors{})
				}
				IsExpected().To(Response(be.CommonErrors[be.ParamsError]))
			})
		})
		When("giving GORM not found error", func() {
			It("renders not found error", func() {
				Action = func(c *gin.Context) {
					Error(c, gorm.ErrRecordNotFound)
				}
				IsExpected().To(Response(be.CommonErrors[be.NotFound]))
			})
		})
		When("giving go_redis not found error", func() {
			It("renders not found error", func() {
				Action = func(c *gin.Context) {
					Error(c, redis.Nil)
				}
				IsExpected().To(Response(be.CommonErrors[be.NotFound]))
			})
		})
		When("giving any other error", func() {
			It("renders unknown error", func() {
				Action = func(c *gin.Context) {
					Error(c, errors.New("abc"))
				}
				IsExpected().To(Response(be.CommonErrors[be.Unknown]))
				IsExpected().To(Response(errors.New("abc")))
			})
		})
	})

	Describe("Result", func() {
		BeforeEach(func() {
			factory.Create(&post)
		})

		It("renders dbx.Record's Data by default", func() {
			Action = func(c *gin.Context) {
				Result(c, dbx.Result{Data: post})
			}

			ax := IsExpected()
			ax.To(ResponseSuccess())
			ax.ResponseData().To(Include(gin.H{"bar": "foo", "id": post.ID}))
		})

		It("renders dbx.Record's Data and Count", func() {
			Action = func(c *gin.Context) {
				Result(c, dbx.Result{Data: post, Total: 1})
			}

			ax := IsExpected()
			ax.To(ResponseSuccess())
			ax.ResponseData().To(Include(gin.H{"total": 1}))
			ax.ResponseData().To(Include(gin.H{"list": gin.H{"bar": "foo", "id": post.ID}}))
		})

		It("renders dbx.Result with optional serializations", func() {
			Action = func(c *gin.Context) {
				Result(c, dbx.Result{Data: post, Total: 1}, dbx_model.Serialization{Add: map[string]string{"zoo": "Zoo"}})
			}
			IsExpected().ResponseData().To(Include(gin.H{"list": gin.H{"bar": "foo", "zoo": "zoo", "id": post.ID}}))
		})

		When("error occurs", func() {
			It("renders error", func() {
				Action = func(c *gin.Context) {
					Result(c, dbx.Result{Err: errors.New("")})
				}
				IsExpected().To(Response(be.CommonErrors[be.Unknown]))
			})
		})
	})

	Describe("ResultOf", func() {
		When("error occurs", func() {
			It("renders error", func() {
				Action = func(c *gin.Context) {
					ResultOf(c, dbx.Result{Err: gorm.ErrRecordNotFound})
				}
				IsExpected().To(Response(be.CommonErrors[be.NotFound]))

				Action = func(c *gin.Context) {
					ResultOf(c, errors.New(""))
				}
				IsExpected().To(Response(be.CommonErrors[be.Unknown]))
			})
		})

		When("nothing error", func() {
			It("responses success", func() {
				Action = func(c *gin.Context) {
					ResultOf(c, nil)
				}
				IsExpected().To(ResponseSuccess())
			})
		})
	})
})
