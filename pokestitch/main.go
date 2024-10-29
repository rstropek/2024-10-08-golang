package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"io"
	"math"
	"net/http"
	"os"
	"reflect"
	"sync"

	gim "github.com/ozankasikci/go-image-merge" // Importing a third-party package for image manipulation
)

const POKEMON = "pikachu"

// pokemon struct to unmarshal JSON data from PokeAPI.
type pokemon struct {
	Sprites pokemonSprites `json:"sprites"`
}

// pokemonSprites struct to map the JSON sprite data.
type pokemonSprites struct {
	// Fields for different sprite URLs
	BackDefault      string `json:"back_default"`
	BackFemale       string `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  string `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}

func main() {
	res, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", POKEMON))
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close() 
	if err != nil {
		panic(err)
	}

	var pokemon pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		panic(err)
	}

	imageUrls := make([]string, 0)
	spritesValue := reflect.ValueOf(pokemon.Sprites)
	
	for i := 0; i < spritesValue.NumField(); i++ {
		field := spritesValue.Field(i)
		imageUrls = append(imageUrls, field.String())
	}

	images := make(chan image.Image, 8)
	wg := sync.WaitGroup{}
	wg.Add(len(imageUrls))

	for _, url := range imageUrls {
		go getImage(url, &wg, images)
	}

	wg.Wait()
	close(images) // Closing the channel

	// Preparing images for merging
	grids := make([]*gim.Grid, 0)
	for img := range images {
		grids = append(grids, &gim.Grid{Image: img})
	}

	// Merging images
	rgba, err := gim.New(grids, 2, int(math.Ceil(float64(len(grids))/float64(2)))).Merge()
	if err != nil {
		return
	}

	// Encode image into png
	b := new(bytes.Buffer)
	wr := bufio.NewWriter(b)
	png.Encode(wr, rgba)
	wr.Flush()

	// write to file
	os.WriteFile("pokemon.png", b.Bytes(), 0644)
}

func getImage(url string, wg *sync.WaitGroup, images chan<- image.Image) {
	defer wg.Done()

	if len(url) == 0 {
		return
	}

	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()

	img, _, err := image.Decode(res.Body)
	if err != nil {
		return
	}

	images <- img
}
