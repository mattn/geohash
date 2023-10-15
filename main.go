package main

import (
	"fmt"
	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/mmcloughlin/geohash"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s address\n", os.Args[0])
		os.Exit(1)
	}
	resp, err := openstreetmap.Geocoder().Geocode(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	if resp == nil {
		fmt.Fprintf(os.Stderr, "The location %q not found\n", os.Args[1])
		os.Exit(1)
	}
	fmt.Println(geohash.Encode(resp.Lat, resp.Lng))
}
