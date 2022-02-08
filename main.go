package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"github.com/richardbieringa/dblp-parser/types"
	"golang.org/x/net/html/charset"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: dblp-parser DIRECTORY")
	}

	parseXmlFile("data/dblp.xml")
}

func parseXmlFile(path string) {
	fileReader, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	var data types.DBLP

	decoder := xml.NewDecoder(fileReader)

	// Converts the ISO encoding to UTF-8
	decoder.CharsetReader = charset.NewReaderLabel

	// Throw an error if something can not be parsed
	decoder.Strict = true

	// Defines a default set of HTML entities to escape
	decoder.Entity = xml.HTMLEntity

	err = decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Articles: %v\n", len(data.Articles))
	fmt.Printf("Inproceeding: %v\n", len(data.Inproceedings))
	fmt.Printf("Proceeding: %v\n", len(data.Proceedings))
	fmt.Printf("Books: %v\n", len(data.Books))
	fmt.Printf("Incollection: %v\n", len(data.Incollections))
	fmt.Printf("PHD Thesis: %v\n", len(data.PHDThesis))
	fmt.Printf("MSC Thesis: %v\n", len(data.MasterThesis))
	fmt.Printf("Data: %v\n", len(data.Data))
	fmt.Printf("WWW: %v\n", len(data.WWW))

}
