package main

import (
	"fmt"
	"log"
	"os"

	api "github.com/Bialson/solarenergy/proto"
	"github.com/urfave/cli/v2"
)

const (
	port = ":8080"
)

func main() {
	// Create a new CLI app
	app := &cli.App{
		EnableBashCompletion: true,
		Name:                 "solarenergy",
		Usage:                "Power consumption from homes by area in Poland",
		Commands: []*cli.Command{
			{
				Name:    "eco-energy",
				Usage:   "Get energy from renewable sources in Poland",
				Aliases: []string{"eco"},
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "year",
						Aliases: []string{"y"},
						Usage:   "Set year to get energy from",
						Value:   2020,
						Action: func(cCtx *cli.Context, v int) error {
							if v < 2010 {
								return fmt.Errorf("Year cannot be lower than 2010")
							}
							return nil
						},
					},
					&cli.IntFlag{
						Name:    "amount",
						Aliases: []string{"a"},
						Usage:   "Set amount of results to get",
						Value:   0,
						Action: func(cCtx *cli.Context, v int) error {
							if v > 80 {
								return fmt.Errorf("Results amount cannot be higher than 204")
							}
							return nil
						},
					},
					&cli.StringFlag{
						Name:    "type",
						Aliases: []string{"t"},
						Usage:   "Set type of renewable source to get energy from",
						Value:   "",
					},
					&cli.StringFlag{
						Name:    "unit",
						Aliases: []string{"u"},
						Usage:   "Set unit of energy amount",
						Value:   "",
					},
				},
				Action: func(cCtx *cli.Context) error {
					fmt.Printf("SolarEnergy Service Client v0.5.0\n\n")
					fmt.Println("Given parameters:")
					fmt.Printf("Type: %s, Year: %d, Amount: %d\n", cCtx.String("type"), cCtx.Int("year"), cCtx.Int("amount"))
					fmt.Println(" * Connecting to server...")
					//Create client connection to server
					client, err, conn := CreateClient()
					defer conn.Close()
					if err != nil {
						fmt.Printf("Error while connecting to server: %v", err)
						os.Exit(1)
					}
					fmt.Printf(" * Creating request...\n\n")
					//Create request based on CLI given flags
					params := &api.EcoEnergyRequest{
						Year:           cCtx.Int64("year"),
						ResponseAmount: cCtx.Int64("amount"),
						Type:           cCtx.String("type"),
						Unit:           cCtx.String("unit"),
					}
					GetEcoEnergy(client, params) //Calling request method
					return nil
				},
			},
			{
				Name:    "get-power",
				Usage:   "Get power consumption from homes by area in Poland",
				Aliases: []string{"gp"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "region",
						Aliases: []string{"r"},
						Usage:   "Set region to get power consumption from",
						Value:   "",
					},
					&cli.StringFlag{
						Name:    "character",
						Aliases: []string{"c"},
						Usage:   "Set character of area to get power consumption from",
						Value:   "",
					},
					&cli.IntFlag{
						Name:    "year",
						Aliases: []string{"y"},
						Usage:   "Set year to get power consumption from",
						Value:   2020,
						Action: func(cCtx *cli.Context, v int) error {
							if v < 2000 {
								return fmt.Errorf("Year cannot be lower than 2000")
							}
							return nil
						},
					},
					&cli.IntFlag{
						Name:    "amount",
						Aliases: []string{"a"},
						Usage:   "Set amount of results to get",
						Value:   204,
						Action: func(cCtx *cli.Context, v int) error {
							if v > 204 {
								return fmt.Errorf("Results amount cannot be higher than 204")
							}
							return nil
						},
					},
				},
				Action: func(cCtx *cli.Context) error {
					fmt.Printf("SolarEnergy Service Client v0.5.0\n\n")
					fmt.Println("Given parameters:")
					fmt.Printf("Region: %s, Character: %s, Year: %d, Amount: %d\n", cCtx.String("region"), cCtx.String("character"), cCtx.Int("year"), cCtx.Int("amount"))
					fmt.Println(" * Connecting to server...")
					//Create client connection to server
					client, err, conn := CreateClient()
					defer conn.Close()
					if err != nil {
						fmt.Printf("Error while connecting to server: %v", err)
						os.Exit(1)
					}
					fmt.Printf(" * Creating request...\n\n")
					//Create request based on CLI given flags
					params := &api.PowerConsumptionRequest{
						Year:           cCtx.Int64("year"),
						ResponseAmount: cCtx.Int64("amount"),
						Region:         cCtx.String("region"),
						Character:      cCtx.String("character"),
					}
					GetEnergyFromHomes(client, params) //Calling request method
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil { //Running app
		log.Fatal(err)
	}
}
