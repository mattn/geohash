package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/mmcloughlin/geohash"
)

func main() {
	var rev, loc bool
	flag.BoolVar(&rev, "r", false, "decode geohash")
	flag.BoolVar(&loc, "l", false, "show location")
	flag.Parse()
	if rev {
		if flag.NArg() != 1 && flag.NArg() != 2 {
			fmt.Fprintf(os.Stderr, "Usage: %s -r lat lng\n", os.Args[0])
			os.Exit(1)
		}

		var lat, lng float64
		var err error
		if flag.NArg() == 1 {
			lat, lng = geohash.Decode(flag.Arg(0))
		} else {
			lat, err = strconv.ParseFloat(flag.Arg(0), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, err.Error())
				os.Exit(1)
			}
			lng, err = strconv.ParseFloat(flag.Arg(1), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, err.Error())
				os.Exit(1)
			}
		}
		fmt.Printf("Latitude: %f\n", lat)
		fmt.Printf("Longitude: %f\n", lng)
		if loc {
			addr, err := openstreetmap.Geocoder().ReverseGeocode(lat, lng)
			if err != nil {
				fmt.Fprintf(os.Stderr, err.Error())
				os.Exit(1)
			}
			fmt.Printf("FormattedAddress: %s\n", addr.FormattedAddress)
			fmt.Printf("Street: %s\n", addr.Street)
			fmt.Printf("HouseNumber: %s\n", addr.HouseNumber)
			fmt.Printf("Suburb: %s\n", addr.Suburb)
			fmt.Printf("Postcode: %s\n", addr.Postcode)
			fmt.Printf("State: %s\n", addr.State)
			fmt.Printf("StateCode: %s\n", addr.StateCode)
			fmt.Printf("StateDistrict: %s\n", addr.StateDistrict)
			fmt.Printf("County: %s\n", addr.County)
			fmt.Printf("Country: %s\n", addr.Country)
			fmt.Printf("CountryCode: %s\n", addr.CountryCode)
			fmt.Printf("City: %s\n", addr.City)
		}
	} else {
		if flag.NArg() != 1 {
			fmt.Fprintf(os.Stderr, "Usage: %s address\n", os.Args[0])
			os.Exit(1)
		}
		resp, err := openstreetmap.Geocoder().Geocode(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		if resp == nil {
			fmt.Fprintf(os.Stderr, "The location %q not found\n", flag.Arg(0))
			os.Exit(1)
		}
		fmt.Println(geohash.Encode(resp.Lat, resp.Lng))
	}
}
