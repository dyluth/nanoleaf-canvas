package main

import (
	"flag"
	"fmt"

	canvas "github.com/dyluth/nanoleaf-canvas"
)

func main() {
	IPAddr := flag.String("url", "", "ip or URL for connecting to the nanoleaf")
	APIKey := flag.String("apikey", "", "APIKey for talking to the canvas.  If not set, will need to be run with the canvas in pairing mode to get one")

	canvas := canvas.New(*IPAddr, *APIKey)
	if *APIKey == "" {
		key, err := canvas.GetNewAPIKey()
		if err != nil {
			panic(fmt.Sprintf("CANNOT GET APIKEY: %v\n", err.Error()))
		}
		fmt.Printf("NEW APIKEY: %v\n", key)
	}
	fmt.Printf("\n===PANEL INFO===\n")
	err := canvas.GetPanelInfo()
	if err != nil {
		panic(fmt.Sprintf("CANNOT GET PANEL INFO: %v\n", err.Error()))
	}
	fmt.Printf("\n===EFFECTS LIST===\n")
	effects, err := canvas.GetEffectsList()
	if err != nil {
		panic(fmt.Sprintf("CANNOT GET EFFECTS LIST: %v\n", err.Error()))
	}
	fmt.Printf("%+v\n", effects)
	fmt.Println("...done")
}
