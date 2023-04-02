package middleware

import (
	"account-service/domain"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/elgs/gojq"
	"github.com/gin-gonic/gin"
)

func ValidateJwt(jwtService string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || header == "Bearer " {
			c.JSON(http.StatusUnauthorized, domain.BuildResponse(http.StatusUnauthorized, "Unauthorized, no token", domain.EmptyObj{}))
			c.Abort()
			return
		}

		// server := bootstrap.JwtService(jwtService)
		// url := fmt.Sprintf("http://%s%s/v1/jwt-auth", server.Host, server.Port)
		url := "http://192.168.18.13:4041/v1/jwt-auth"

		req, err := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", header)
		if err != nil {
			log.Println("Request error", err)
			c.JSON(http.StatusInternalServerError, domain.BuildResponse(http.StatusInternalServerError, err.Error(), domain.EmptyObj{}))
			c.Abort()
			return
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response. \n[ERROR] -", err)
			c.JSON(http.StatusUnauthorized, domain.BuildResponse(http.StatusUnauthorized, err.Error(), domain.EmptyObj{}))
			c.Abort()
			return
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error while reading the response bytes:", err)
			c.JSON(http.StatusInternalServerError, domain.BuildResponse(http.StatusInternalServerError, err.Error(), domain.EmptyObj{}))
			c.Abort()
			return
		}

		str := string(body)
		parser, _ := gojq.NewStringQuery(str)
		status, _ := parser.QueryToInt64("status_code")
		id, _ := parser.QueryToString("status_message")
		if parser == nil {
			c.JSON(http.StatusInternalServerError, domain.BuildResponse(http.StatusInternalServerError, "Something went wrong", domain.EmptyObj{}))
			c.Abort()
			return
		}

		if status != 200 {
			c.JSON(http.StatusUnauthorized, domain.BuildResponse(http.StatusUnauthorized, "Not authorized", domain.EmptyObj{}))
			c.Abort()
			return
		}

		c.Set("x-user-id", id)
		c.Next()

	}
}
