package main

import (
	"net/http"

	"github.com/TanmayKumar005/2302900100247/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJNYXBDbGFpbXMiOnsiYXVkIjoiaHR0cDovLzIwLjI0NC41Ni4xNDQvZXZhbHVhdGlvbi1zZXJ2aWNlIiwiZW1haWwiOiJ0YW5tYXkyMjYwNUBnbWFpbC5jb20iLCJleHAiOjE3ODA0NjM1NjAsImlhdCI6MTc4MDQ2MjY2MCwiaXNzIjoiQWZmb3JkIE1lZGljYWwgVGVjaG5vbG9naWVzIFByaXZhdGUgTGltaXRlZCIsImp0aSI6ImViNDdjZDhiLWNmNmYtNGUxMS04NzU3LWNhZjAwN2M2NjNlNSIsImxvY2FsZSI6ImVuLUlOIiwibmFtZSI6InRhbm1heSBrdW1hciIsInN1YiI6ImYzYTg3NDJlLWNhMGMtNGY5Yy04ZTZkLWM3OTJiMWMzNGU2YSJ9LCJlbWFpbCI6InRhbm1heTIyNjA1QGdtYWlsLmNvbSIsIm5hbWUiOiJ0YW5tYXkga3VtYXIiLCJyb2xsTm8iOiIyMzAyOTAwMTAwMjQ3IiwiYWNjZXNzQ29kZSI6InNkV1dnYyIsImNsaWVudElEIjoiZjNhODc0MmUtY2EwYy00ZjljLThlNmQtYzc5MmIxYzM0ZTZhIiwiY2xpZW50U2VjcmV0IjoiQVdtd2JGVWdqanBOTkhoSiJ9.06lCkoISWr4sI-oD-1PsOmZihOaXt7zAU_jVwI2g-vk"

	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		err := middleware.Log(
			token,
			"backend",
			"info",
			"middleware",
			"Logger stup successful",
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Log sent successfully",
		})
	})

	r.Run(":7070")

}
