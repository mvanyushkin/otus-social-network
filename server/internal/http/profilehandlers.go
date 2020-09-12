package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mvanyushkin/otus-social-network/internal/domain"
	"net/http"
	"strconv"
)

func getProfileHandler(profileService domain.ProfileService) func(c *gin.Context) {
	return func(c *gin.Context) {
		profileIdParameter := c.Param("id")
		profileId, err := strconv.ParseUint(profileIdParameter, 0, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		profile, err := profileService.GetProfile(c, profileId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, profile)
	}
}

func updateProfile(profileService domain.ProfileService) func(c *gin.Context) {
	return func(c *gin.Context) {
		profileIdParameter := c.Param("id")
		profileId, err := strconv.ParseUint(profileIdParameter, 0, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

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

		err = profileService.UpdateProfile(c, profileId, profile)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.String(200, "")
		}
	}
}

func cancelFriendship(service domain.ProfileService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	}
}

func makeFriendship(service domain.ProfileService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	}
}

func getFriends(profileService domain.ProfileService) gin.HandlerFunc {
	return func(c *gin.Context) {
		profileId := c.Param("id")
		id, _ := strconv.ParseInt(profileId, 0, 64)
		profileService.GetFriendsList(c, uint64(id))
		c.JSON(200, gin.H{
			"hello": "world",
			"id":    profileId,
		})
	}
}
