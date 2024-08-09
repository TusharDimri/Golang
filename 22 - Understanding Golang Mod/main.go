package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// go mod tidy: Very important if we're using packages.
	// go list -m all : This command gives a list of deendencies for the module.

	// Output:
	// github.com/TusharDimri/mymodules
	// github.com/gorilla/mux v1.8.1

	// go list -m -versions github.com/gorilla/mux

	//  go mod why github.com/gorilla/mux
	// go mod graph : dependencies and their relation graph

	// go mod edit : to edit mod file from terminal.
	// go mod edit -go 1.20
	// go mod edit -module "new module name"

	// Bring the package to the module folder and now you can edit the third party packages:
	// go mod vendor

	// go run -mod=vendor main.go
	// First check vedor and then the cache for packages.

	// All these are exoensive operations and we need to be cqreful when working with CI/CD.

	// We dont use workspaces anymre and have shifted to modules.

	// Read go mod documentation for more understanding.

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Golang API</h1>"))
}
