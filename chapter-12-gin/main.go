package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	var group errgroup.Group
	enSecureServer := &http.Server{
		Addr:         ":8080",
		Handler:      router(),
		ReadTimeout:  time.Second * 6,
		WriteTimeout: time.Second * 6,
	}
	secureServer := &http.Server{
		Addr:         ":8443",
		Handler:      router(),
		ReadTimeout:  time.Second * 6,
		WriteTimeout: time.Second * 6,
	}

	group.Go(func() error {
		err := enSecureServer.ListenAndServe()
		if err != nil {
			log.Fatalln(err)
		}
		return nil
	})
	var filePath string
	var err error
	if filePath, err = os.Getwd(); err != nil {
		log.Fatalln(err)
	}

	group.Go(func() error {

		err := secureServer.ListenAndServeTLS(filePath+"/chapter-12-gin/server.pem", filePath+"/chapter-12-gin/server.key")
		if err != nil {
			log.Fatalln(err)
		}
		return nil
	})
	if err := group.Wait(); err != nil {
		log.Fatalln(err)
	}
}

type Product struct {
	Username    string    `json:"username" binding:"required"`
	Name        string    `json:"name,omitempty" binding:"required"`
	Category    string    `json:"category,omitempty" binding:"required"`
	Price       int       `json:"price,omitempty" binding:"gte=0"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

type productHandler struct {
	sync.RWMutex
	products map[string]Product
}

func newProductHandler() *productHandler {
	return &productHandler{
		products: make(map[string]Product),
	}

}

func (u *productHandler) Create(c *gin.Context) {
	u.Lock()
	defer u.Unlock()

	// 1. 参数解析
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. 参数校验
	if _, ok := u.products[product.Name]; ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("product %s already exist", product.Name)})
		return
	}
	product.CreatedAt = time.Now()

	// 3. 逻辑处理
	u.products[product.Name] = product
	log.Printf("Register product %s success", product.Name)

	// 4. 返回结果
	c.JSON(http.StatusOK, product)
}

func (u *productHandler) Get(c *gin.Context) {
	u.Lock()
	defer u.Unlock()

	product, ok := u.products[c.Param("name")]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Errorf("can not found product %s", c.Param("name"))})
		return
	}

	c.JSON(http.StatusOK, product)
}

func router() http.Handler {
	router := gin.Default()
	productHandler := newProductHandler()
	// 路由分组、中间件、认证
	v1 := router.Group("/v1")
	{
		productv1 := v1.Group("/products")
		{
			// 路由匹配
			productv1.POST("", productHandler.Create)
			productv1.GET(":name", productHandler.Get)
		}
	}

	return router
}
