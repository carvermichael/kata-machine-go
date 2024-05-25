package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Config struct {
	Algos []string
}

func main() {

	if len(os.Args) > 1 && os.Args[1] == "generate" {
		fmt.Println("generate!")

		bytes, err := os.ReadFile("config.yaml")
		if err != nil {
			log.Fatal("Failed to find config file.")
			return
		}

		config := Config{}
		err = yaml.Unmarshal(bytes, &config)
		if err != nil {
			log.Fatal("Failed to parse config file.")
			return
		}

		// Figure out what Day it is
		files, err := os.ReadDir(".")
		if err != nil {
			log.Fatal("Failed to read contents of base-level folder.")
			return
		}

		days := make([]int, 5)
		for _, v := range files {
			if v.IsDir() && strings.Contains(v.Name(), "day") {
				dayStr := string(v.Name()[3])
				dayInt, err := strconv.Atoi(dayStr)
				if err != nil {
					log.Fatal("Failed to parse existing day folders.")
					return
				}
				days = append(days, dayInt)
			}
		}

		sort.Ints(days)
		newDay := days[len(days)-1] + 1
		dayFolder := fmt.Sprintf("day%d", newDay)

		// Create new Day folder
		err = os.Mkdir(dayFolder, 0750)
		if err != nil {
			log.Fatal("Failed to create Day folder.")
			return
		}

		// Copy contents of relevant files (empty impl and test) into day folder
		kataPath := "./kata/"
		for _, algoName := range config.Algos {

			kataFiles, err := os.ReadDir(kataPath)
			if err != nil {
				log.Fatal("Failed to read contents of kata folder.")
				return
			}
			for _, w := range kataFiles {
				if !w.IsDir() && strings.Contains(w.Name(), algoName) {
					pathStr := kataPath + w.Name()
					fileBytes, err := os.ReadFile(pathStr)
					if err != nil {
						log.Fatalf("Failed to read file: %s", pathStr)
					}

					fileString := string(fileBytes)
					fileString = strings.Replace(fileString, "package kata", "package "+dayFolder, 1)

					newPath := dayFolder + "/" + w.Name()
					err = os.WriteFile(newPath, []byte(fileString), 0750)
					if err != nil {
						log.Fatalf("Failed to write file: %s", newPath)
					}
				}
			}
		}
	} else {
		fmt.Println("degenerate")
	}
}
