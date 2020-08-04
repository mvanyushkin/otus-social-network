package http

import (
	"backend/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Serve(addr string, profileService domain.ProfileService, searcher domain.ProfileSearcher, accountService domain.AccountService) error {
	router := gin.Default()
	router.POST("/api/register-new", capturePanic(registerNewHandler(accountService)))
	router.POST("/api/login", capturePanic(loginHandler(accountService)))
	router.POST("/api/logout", capturePanic(logoutHandler(accountService)))
	router.GET("/api/get-profile/:id", capturePanic(getProfileHandler(profileService)))
	router.POST("/api/update-profile/:id", capturePanic(updateProfile(profileService)))
	router.GET("/api/get-friends/:id", capturePanic(getFriends(profileService)))
	router.POST("/api/make-friendship", capturePanic(makeFriendship(profileService)))
	router.POST("/api/cancel-friendship", capturePanic(cancelFriendship(profileService)))
	router.POST("/api/search-people", capturePanic(searchPeople(searcher)))
	return router.Run(addr)
}

func capturePanic(callee func(c *gin.Context)) func(c *gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				c.String(http.StatusInternalServerError, "Internal server error")
			}
		}()
		callee(c)
	}
}

type ProfileDto struct {
	Id           uint64 `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordhash"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Age          uint8  `json:"age"`
	Gender       uint8  `json:"gender"`
	City         string `json:"city"`
	Hobby        string `json:"hobby"`
}

type SearchByFirstNameLastNameDto struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type SearchByCityDto struct {
	City string `json:"city"`
}

type SearchByAgeDto struct {
	MinAge uint8 `json:"minage"`
	MaxAge uint8 `json:"maxage"`
}
