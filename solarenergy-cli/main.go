package main

import (
	"fmt"
	"log"
	"os"

	api "github.com/Bialson/solarenergy/proto"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	app := &cli.App{
		EnableBashCompletion: true,
		Name:                 "solarenergy",
		Usage:                "Power consumption from homes by area in Poland",
		Commands: []*cli.Command{
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
						Value:   0,
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
					conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
					if err != nil {
						fmt.Printf(" * Cannot connect to server: %v \n Exiting", err)
						os.Exit(1)
					} else {
						fmt.Println(" * Connected to server!")
					}
					defer conn.Close()
					client := api.NewSolarServiceClient(conn)
					fmt.Printf(" * Creating request...\n\n")
					params := &api.PowerConsumptionRequest{
						Year:           cCtx.Int64("year"),
						ResponseAmount: cCtx.Int64("amount"),
						Region:         cCtx.String("region"),
						Character:      cCtx.String("character"),
					}
					givePowerByArea(client, params)
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
