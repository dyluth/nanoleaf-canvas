package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	canvas "github.com/dyluth/nanoleaf-canvas"
)

func main() {
	IPAddr := flag.String("url", "", "ip or URL for connecting to the nanoleaf")
	APIKey := flag.String("apikey", "", "APIKey for talking to the canvas.  If not set, will need to be run with the canvas in pairing mode to get one")
	flag.Parse()

	fmt.Printf("IP ADDR: %v\n", *IPAddr)
	canvas := canvas.New(*IPAddr, *APIKey)
	if *APIKey == "" {
		key, err := canvas.GetNewAPIKey()
		if err != nil {
			flag.Usage()
			panic(fmt.Sprintf("CANNOT GET APIKEY: %v\n", err.Error()))
		}
		fmt.Printf("NEW APIKEY: %v\n", key)
	}
	fmt.Printf("\n===PANEL INFO===\n")
	info, err := canvas.GetPanelInfo()
	if err != nil {
		flag.Usage()
		panic(fmt.Sprintf("CANNOT GET PANEL INFO: %v\n", err.Error()))
	}
	fmt.Printf("current Effect: %v\n", info.Effects.Select)

	fmt.Printf("\n===EFFECTS LIST===\n")
	effects, err := canvas.GetEffectsList()
	if err != nil {
		panic(fmt.Sprintf("CANNOT GET EFFECTS LIST: %v\n", err.Error()))
	}
	fmt.Printf("%v\n", strings.Join(effects, ", "))
	fmt.Println("...done")

	for _, effect := range effects {
		if effect != info.Effects.Select {
			fmt.Printf("Setting effect to: %v\n", effect)
			canvas.SetEffect(effect)
			time.Sleep(10 * time.Second)
		}
	}
	fmt.Printf("resetting effect to: %v\n", info.Effects.Select)
	canvas.SetEffect(info.Effects.Select)
}
