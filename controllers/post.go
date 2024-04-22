package controllers

import (
	"ci_cd/models"
	"ci_cd/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllPosts(c echo.Context) error {
	posts, err := services.GetAllPosts(c.Request().Context())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse[any]{
			Status:  false,
			Message: "error when fetching posts",
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse[[]models.Post]{
		Status:  true,
		Message: "all posts",
		Data:    posts,
	})
}

func CreatePost(c echo.Context) error {
	var postRequest models.PostRequest

	if err := c.Bind(&postRequest); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse[any]{
			Status:  false,
			Message: "invalid request",
		})
	}

	post, err := services.CreatePost(c.Request().Context(), postRequest)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse[any]{
			Status:  false,
			Message: "error when creating post",
		})
	}

	return c.JSON(http.StatusCreated, models.BaseResponse[models.Post]{
		Status:  true,
		Message: "post created",
		Data:    post,
	})
}
