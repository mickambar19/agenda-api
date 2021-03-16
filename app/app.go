package app

import (
	"fmt"
	"net/http"

	"github.com/amjimenez/agenda-api/controllers"
)

func StartApp() {
	fmt.Println("Print ln")
	http.HandleFunc("/users", controllers.GetUsers)

	if err := http.ListenAndServe("localhost:8000", nil); err != nil {
		panic(err)
	}
}
