package middleware

import "github.com/gkzy/gow"

func Auth() gow.HandlerFunc{
	return func(c *gow.Context){
		c.Header("Access-Control-Allow-Origin","*")
		c.Next()
	}

}
