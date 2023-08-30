package main

import (
	"net/http"

	"github.com/daisyelem/RecipeKeeper/data"
	"github.com/daisyelem/RecipeKeeper/handlers"
)

func main() {
	data.FetchAllRecipes()
	//fmt.Println(data.AllRecipes)
	http.HandleFunc("/", handlers.HomePage)
	http.ListenAndServe(":8000", nil)
}
