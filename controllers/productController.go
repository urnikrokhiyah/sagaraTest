package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"test/lib/database"
	"test/middlewares"
	"test/models"
	"test/responses"

	"github.com/labstack/echo/v4"
)

func CreateNewProductController(c echo.Context) error {
	_, checkRole := middlewares.ExtractToken(c)
	if checkRole != "admin" {
		return c.JSON(http.StatusBadRequest, responses.FailedResponse("failed", "unauthorized access"))
	}

	newProduct := models.Product{}
	c.Bind(&newProduct)

	product, err := database.CreateNewProduct(newProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedResponse("failed", "failed to create new product"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponse("success", "success to create new product", product))
}

func GetListProductController(c echo.Context) error {
	listProduct, rowAffected, err := database.GetListProduct()
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedResponse("failed", "failed to get list of products"))
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusOK, responses.FailedResponse("success", "list of products not found"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponse("success", "success to get list of products", listProduct))
}

func GetProductDetailController(c echo.Context) error {
	productId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedResponse("failed", "invalid product id"))
	}
	product, rowAffected, err := database.GetProductDetail(productId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedResponse("failed", "failed to get product detail"))
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusOK, responses.FailedResponse("success", "product detail not found"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponse("success", "success to get product detail", product))
}

func DeleteProductController(c echo.Context) error {
	_, checkRole := middlewares.ExtractToken(c)
	if checkRole != "admin" {
		return c.JSON(http.StatusBadRequest, responses.FailedResponse("failed", "unauthorized access"))
	}

	productId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedResponse("failed", "invalid product id"))
	}
	messageDeleted, rowAffected, err := database.DeleteProduct(productId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedResponse("failed", "failed to delete product data"))
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusOK, responses.FailedResponse("success", "product data not found"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponse("success", "success to delete product data", messageDeleted))
}

func UpdateProductController(c echo.Context) error {
	_, checkRole := middlewares.ExtractToken(c)
	if checkRole != "admin" {
		return c.JSON(http.StatusBadRequest, responses.FailedResponse("failed", "unauthorized access"))
	}

	productId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedResponse("failed", "invalid product id"))
	}

	newDataProduct := models.Product{}
	c.Bind(&newDataProduct)

	updatedProduct, rowAffected, err := database.UpdateProduct(productId, newDataProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedResponse("failed", "failed to update product data"))
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusOK, responses.FailedResponse("success", "product data not found"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponse("success", "success to update product data", updatedProduct))
}

func UpdateImageProductController(c echo.Context) error {
	_, checkRole := middlewares.ExtractToken(c)
	if checkRole != "admin" {
		return c.JSON(http.StatusBadRequest, responses.FailedResponse("failed", "unauthorized access"))
	}

	productId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedResponse("failed", "invalid product id"))
	}

	imageId := c.Param("id")

	file, errorFile := c.FormFile("image")
	if errorFile != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedResponse("failed", "can't upload image of product"))
	}

	source, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedResponse("failed", "can't upload image of product"))
	}

	defer source.Close()

	dir, err := os.Getwd()
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedResponse("failed", "can't upload image of product"))
	}

	filename := fmt.Sprintf("%s%s", imageId, filepath.Ext(file.Filename))

	fileLocation := filepath.Join(dir, "images", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	destination, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer destination.Close()

	if _, err := io.Copy(targetFile, source); err != nil {
		return err
	}

	fmt.Println(fileLocation, destination)

	updatedProduct, rowAffected, err := database.UpdateImage(productId, fileLocation)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.FailedResponse("failed", "failed to update image of product"))
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusOK, responses.FailedResponse("success", "product data not found"))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponse("success", "success to update image of product", updatedProduct))
}
