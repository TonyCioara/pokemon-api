package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

// type Ability struct {
// 	Name string `json:"name"`
// 	Url  string `json:"url"`
// }

// type Pokemon struct {
// 	Abilities []Ability
// }
func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getPokemon(c echo.Context) error {
	pokemonName := c.QueryParam("name")
	res, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + pokemonName)

	if err != nil {
		fmt.Println("error")
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error")
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(data), &result)

	return c.JSON(http.StatusOK, result)
}

func main() {
	e := echo.New()

	e.GET("/", helloWorld)
	e.GET("/pokemon", getPokemon)

	e.Logger.Fatal(e.Start(":1323"))
}
