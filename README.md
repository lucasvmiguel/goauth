
# OAUTH2  [![GoDoc](http://godoc.org/github.com/lucasvmiguel/goauth1?status.png)](http://godoc.org/github.com/lucasvmiguel/goauth)

"With the resource owner password credentials grant type, the user provides their service credentials (username and password) directly to the application, which uses the credentials to obtain an access token from the service. This grant type should only be enabled on the authorization server if other flows are not viable. Also, it should only be used if the application is trusted by the user (e.g. it is owned by the service, or the user's desktop OS)." Mitchell Anicas

## Install

  go get github.com/dghubble/oauth1

## Documentation

  Read [GoDoc](http://godoc.org/github.com/lucasvmiguel/goauth)

### Authorization Flow

###### SIMPLE REQUEST WITH TOKEN

  <img src="https://pbs.twimg.com/media/CN3P89gWEAEMcuV.png">

###### REQUEST TO GET TOKEN

  <img src="https://pbs.twimg.com/media/CN3P89UW8AAMrvj.png">

###### REQUEST TO REFRESH TOKEN

  <img src="https://pbs.twimg.com/media/CN3P8-WWgAU1llO.png">

###### REQUEST TO DESTROY TOKEN

  <img src="https://pbs.twimg.com/media/CN3P898WoAAIAP_.png">

#### CODE ON SERVER

      package main
      
      import (
      	"github.com/gin-gonic/gin"
      	"github.com/lucasvmiguel/goauth"
      )
      
      func main() {
      	router := gin.Default()
      
      	//Initialize the goauth
      	//router - gin.Engine
      	//ctrlAuthentication - controller used to authenticate
      	//ctrlAuthorization - controller used to authorize
        //true to debug or false to not debug
      	a := goauth.Init(router, ctrlAuthentication, ctrlAuthorization, true)
      
      	//create route with goauth init return if you DO NOT want to authenticate the route
      	//AUTHENTICATE
      	a.GET("/admin", func(c *gin.Context) {
      		c.String(200, "ADMIN")
      	})
      
      	//create route with router(gin.Engine) if you DO NOT want to authenticate the route
      	//NOT AUTHENTICATE
      	router.GET("/teste", func(c *gin.Context) {
      		c.String(200, "TESTE")
      	})
      
      	router.Run(":8082")
      }
      
      //controller authentication gives you the username, password and clientID
      //you need return the id and the object user(optional) if you want authenticate
      func ctrlAuthentication(username string, password string, clientID string) (string, interface{}) {
      	if username == "lucas" && password == "123456" {
      		//AUTHENTICATE
      		return "1", nil
      	}
      	//NOT AUTHENTICATE
      	return "", nil
      }
      
      //controller authentication gives you the path and id user
      //you just need return true to authorize or false to unauthorize
      func ctrlAuthorization(path string, id string) bool {
      	if id == "1" && path == "/profile" {
      		return true
      	}
      	if id == "1" && path == "/admin" {
      		return true
      	}
      	if id == "2" && path == "/teste" {
      		return true
      	}
      
      	return false
      }
      
## License

[MIT License](LICENSE)
