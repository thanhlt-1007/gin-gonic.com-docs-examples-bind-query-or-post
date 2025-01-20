package main

import (
    "log"
    "time"
    "github.com/gin-gonic/gin"
)

type Person struct {
    Name string `form:"name"`
    Address string `form:"address"`
    Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func getPerson(context *gin.Context) {
    var person Person
    err := context.Bind(&person)
    if err == nil {
        log.Println(person.Name)
        log.Println(person.Address)
        log.Println(person.Birthday)
    }

    context.String(200, "Success")
}

func main() {
    router := gin.Default()
    router.GET("/person", getPerson)
    router.Run()
}
