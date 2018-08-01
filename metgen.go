package main

import (
	"fmt"
	"bufio"
	"os"
	"time"
	"io/ioutil"
	"strings"
	"math/rand"
    "path/filepath"
)

var start time.Time
var lastMet string

func standard() {
	reader := bufio.NewReader(os.Stdin)
    fmt.Print("Type: ")
    text, _ := reader.ReadString('\n')
    if text != "f" {
    	generateMetaphor()
    	time.Sleep(time.Second)
    	standard()
    } else {
        d1 := []byte(lastMet)
        err := ioutil.WriteFile(directory() + "favorites.txt", d1, 0664)
        if err != nil {
            fmt.Println(err)
        }
    }
}

func generateMetaphor() {
    path := directory()
	n, err := ioutil.ReadFile(path + "nouns.txt")
	if err != nil {
        fmt.Print(err)
        fmt.Println("\nPlease report this incident to Colin!")
    }

    a, err := ioutil.ReadFile(path + "adjectives.txt")
	if err != nil {
        fmt.Print(err)
        fmt.Println("\nPlease report this incident to Colin!")
    }

    nouns := strings.Split(string(n),"\n")
    adjectives := strings.Split(string(a),"\n")

    rand.Seed(int64(time.Since(start)))

    noun1 := rand.Intn(len(nouns)-1)
    adj1 := rand.Intn(len(adjectives)-1)
    noun2 := rand.Intn(len(nouns)-1)

    lastMet = nouns[noun1] + " is " + adjectives[adj1] + " " + nouns[noun2]

    fmt.Println(lastMet)

}

func directory() string{
    ex, err := os.Executable()

    if err != nil {
        panic(err)
    }

    exPath := filepath.Dir(ex) + "/"
    return exPath
}

func main() {
	start = time.Now()

	fmt.Println("Welcome to the Metaphor Generator!")
	time.Sleep(time.Second)
	fmt.Println("Press enter to generator a new metaphor")
	time.Sleep(time.Second)
	standard()
}