package middleware

import echojwt "github.com/labstack/echo-jwt/v4"

var IsAuthenticated = echojwt.WithConfig(echojwt.Config{

	SigningKey: []byte("secret"),
})

// var IsAuthenticated =  middleware.JWTWithConfig(middleware.JWTConfig{
// 		SigningKey: []byte("secret"),
// 	})
