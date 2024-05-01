package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "strconv"
    "strings"
)

// Structure of every registry
type product struct {
    ID     string  `json:"id"`
    Name   string  `json:"name"`
    Description  string  `json:"description"`
}

// products slice to seed record product data.
var products = []product {
    {ID: "1", Name: "Telephone", Description: "Motorola Moto G54 telephone"},
    {ID: "2", Name: "Television", Description: "Samsumg 50Cu7700 television"},
}

func main() {
    router := gin.Default()
    router.GET("/products", getProducts)
    router.GET("/product/:id", getProductByID)
    router.POST("/product", addProduct)
    router.DELETE("/product/:id", deleteProductByID)
    router.PUT("/product/:id", updateProduct)

    router.Run("localhost:8081")
}

// getProducts responds with the list of all products as JSON.
func getProducts(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, products)
}

// updating a product
func updateProduct(c *gin.Context) {
    id := c.Param("id")
    var newProduct product

    if err := c.BindJSON(&newProduct); err != nil {
        return
    }

    if strings.Trim(newProduct.Name, " ") == "" {
        c.IndentedJSON(http.StatusUnprocessableEntity, "The name of the product is empty")
        return
    } 

    for index, a := range products {
        if a.ID == id {
            newProduct.ID = id
            products[index] = newProduct
            c.IndentedJSON(http.StatusOK, newProduct)
            return
        }
    }

    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})    
}

// addProduct adds a product from a JSON received in the request body.
func addProduct(c *gin.Context) {
    var newProduct product
    var newId int

    newId , _ = strconv.Atoi(products[len(products)-1].ID)
    newProduct.ID = strconv.Itoa(newId + 1)
    
    if err := c.BindJSON(&newProduct); err != nil {
        return
    }

    newProduct.Name = strings.Trim(newProduct.Name, " ")
    if newProduct.Name == "" {
        newProduct.ID = ""
        c.IndentedJSON(http.StatusUnprocessableEntity, "The name of the product is empty")
        return
    } 

    // Addding the new product
    products = append(products, newProduct)
    c.IndentedJSON(http.StatusCreated, newProduct)
}

// geProductByID locates the product whose ID value matches the id
// parameter sent by the client, then returns that product related
// to the received ID as a response.
func getProductByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of products, looking for
    // a product whose ID value matches the parameter.
    for _, a := range products {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

// Here we delete a product by the informed ID
func deleteProductByID(c *gin.Context) {
    id := c.Param("id")

    var productFound product
    var newProducts []product

    for _, a := range products {
        if a.ID == id {
            productFound = a
        } else{
            newProducts = append(newProducts, a)
        }
    }

    if productFound.ID != "" {
        products = newProducts
        c.IndentedJSON(http.StatusOK, productFound)
    } else {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})    
    }
}