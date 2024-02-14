package middlewares

import (
	userRepository "blog/internal/modules/user/repositories"
	"blog/pkg/sessions"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IsGuest() gin.HandlerFunc {
	userRepo := userRepository.New()
	return func(c *gin.Context) {
		authID := sessions.Get(c, "auth")
		userID, _ := strconv.Atoi(authID)

		user := userRepo.FindByID(userID)

		if user.ID != 0 {
			c.Redirect(http.StatusFound, "/")
			return
		}

		c.Next()
	}
}
