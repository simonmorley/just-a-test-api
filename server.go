package main

import (
	"flag"
	"fmt"
	"github.com/PolkaSpots/puffin-api-v2/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var (
	ProjectID = flag.String("project-id", "emulator-project-id", "Set the GCE projcet")
)

func HomeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Gorilla!\n"))
	})
}

type NewUser struct {
	Username  string
	Password1 string
	Password2 string
}

func RegHandler(db db.DbManager) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

			var user NewUser

			err := db.RegisterUser("Steve")
			if err != nil {
				http.Error(
					w,
					"unable to add user",
					http.StatusInternalServerError,
				)
				return
			}
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "Successfully registered %v", user.Username)
		}
	})
}

func init() {
	flag.Parse()

	// if *Creds != "" {
	// 	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", *Creds)
	// }
	// os.Setenv("ENV", *Environment)
	os.Setenv("PROJECT_ID", *ProjectID)
}

func main() {
	d, err := db.NewDatastore()
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.Handle("/cheese", HomeHandler())
	r.Handle("/reg", RegHandler(d))
	// http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":3000", r))

}
