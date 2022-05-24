package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"strconv"
	"github.com/gorrila/mux"
)

type Moive struct {
	ID string `json: "id"`
	Isbn string `json: "isbn"`
	Title string `json: "title"`
	Director *Director `json:"director"`
}

type Director struct {

}

func main() {

}