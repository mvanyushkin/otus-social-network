package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mvanyushkin/otus-social-network/internal/domain"
	"net/http"
	"strconv"
)

func searchPeople(profileService domain.ProfileSearcher) func(c *gin.Context) {
	return func(c *gin.Context) {
		searchTypeParameterValue := c.Query("searchtype")
		searchType, err := strconv.ParseUint(searchTypeParameterValue, 0, 8)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if domain.SearchType(searchType) == domain.ByFirstNameLastName {
			var dto SearchByFirstNameLastNameDto
			if err := c.ShouldBind(&dto); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			profiles, err := profileService.FindByFirstNameLastName(c, dto.FirstName, dto.LastName)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(200, profiles)
			return
		}

		if domain.SearchType(searchType) == domain.ByCity {
			var dto SearchByCityDto
			if err := c.ShouldBind(&dto); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			profiles, err := profileService.FindByCity(c, dto.City)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(200, profiles)
			return
		}

		if domain.SearchType(searchType) == domain.ByAge {
			var dto SearchByAgeDto
			if err := c.ShouldBind(&dto); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			profiles, err := profileService.FindByAge(c, dto.MinAge, dto.MaxAge)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(200, profiles)
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong search critera",
		})
	}
}
