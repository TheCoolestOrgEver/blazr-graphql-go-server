package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
    profileHandlers "./api/profile"
    "./api/rabbit"
)

func Index( w http.ResponseWriter, r *http.Request, _ httprouter.Params ) {
    fmt.Fprint(w, "Welcome!\n")
}

func main() {
    router := httprouter.New()
    router.GET( "/", Index )
    router.GET( "/profile/:userID", profileHandlers.GetProfile )
    router.DELETE( "/profile/:userID", profileHandlers.DeleteProfile )
    router.GET( "/profiles/", profileHandlers.GetProfiles )
    router.POST( "/profile/", profileHandlers.CreateProfile )
    router.PUT( "/profile/", profileHandlers.UpdateProfile )
    router.GET("/matches/:userID", profileHandlers.GetMatches )
    go rabbit.Consume()
    fmt.Println( "\nStarting http server on port 8080 \n" )
    go log.Fatal(http.ListenAndServe(":8080", router))
}