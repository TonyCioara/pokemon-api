package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/labstack/echo"
)

type Ability struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// Routes
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

func getRandomAbility(c echo.Context) error {
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

	abilities := result["abilities"].([]interface{})

	for _, ability := range abilities {
		fmt.Println(ability)
	}

	cap := cap(abilities)
	randnum := rand.Intn(cap)
	ability := abilities[randnum]

	return c.JSON(http.StatusOK, ability)
}

func getAbilities(c echo.Context) error {
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

	abilities := result["abilities"].([]interface{})

	return c.JSON(http.StatusOK, abilities)
}

func main() {
	e := echo.New()

	e.GET("/", helloWorld)
	e.GET("/pokemon", getPokemon)
	e.GET("/abilities", getAbilities)
	e.GET("/abilities/random", getRandomAbility)

	e.Logger.Fatal(e.Start(":1323"))
}
