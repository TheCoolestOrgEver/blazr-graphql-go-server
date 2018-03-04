package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
    profileHandlers "./api/profile"
)

func Index( w http.ResponseWriter, r *http.Request, _ httprouter.Params ) {
    fmt.Fprint(w, "Welcome!\n")
}

func main() {
    router := httprouter.New()
    router.GET( "/", Index )
    router.GET( "/profile/:userID", profileHandlers.GetProfile )
    router.GET( "/profiles/", profileHandlers.GetProfiles )
    router.POST( "/profile/", profileHandlers.CreateProfile )
    router.PUT( "/profile/", profileHandlers.UpdateProfile )

    log.Fatal(http.ListenAndServe(":8080", router))
}