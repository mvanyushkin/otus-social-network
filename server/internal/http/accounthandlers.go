package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mvanyushkin/otus-social-network/internal/domain"
	"net/http"
)

func registerNewHandler(profileService domain.AccountService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var profileDto ProfileDto
		if err := c.ShouldBind(&profileDto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		profile := domain.Profile{
			Email:     profileDto.Email,
			FirstName: profileDto.FirstName,
			LastName:  profileDto.LastName,
			Age:       profileDto.Age,
			Gender:    profileDto.Gender,
			City:      profileDto.City,
			Hobby:     profileDto.Hobby,
		}

		id, err := profileService.RegisterNew(c, profile, profileDto.PasswordHash)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"id": id,
			})
		}
	}
}

func loginHandler(service domain.AccountService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var dto LoginDto
		if err := c.ShouldBind(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err := service.Login(c, dto.Email, dto.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{})
		}
	}
}

func logoutHandler(service domain.AccountService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var dto LoginDto
		if err := c.ShouldBind(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err := service.Login(c, dto.Email, dto.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{})
		}
	}
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"passwordhash"`
}
