package controllers

import (
	"strconv"
	"wareHouse/dao"
	"wareHouse/service"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	SetupRoutes(r *gin.Engine)
}

type productController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &productController{productService}
}

func (c *productController) SetupRoutes(r *gin.Engine) {
	r.POST("/products", c.registerProduct)
	r.GET("/products", c.getAllProducts)
	r.GET("/products/:id", c.getProductByID)
	r.DELETE("/products/:id", c.deleteProduct)
	r.PUT("/products/:id", c.updateProduct)
}

func (c *productController) registerProduct(ctx *gin.Context) {
	var product dao.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := c.productService.Save(product); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, product)
}

func (c *productController) getAllProducts(ctx *gin.Context) {
	products, err := c.productService.GetAllProducts()
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, products)
}

func (c *productController) getProductByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	product, err := c.productService.GetProductByID(uint(id))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, product)
}

func (c *productController) deleteProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	err = c.productService.Delete(uint(id))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Product deleted successfully"})
}

func (c *productController) updateProduct(ctx *gin.Context) {
	var product dao.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := c.productService.Update(product)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Product updated successfully"})
}
