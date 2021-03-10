package main

import (
	"context"
	"errors"
	"os"

	logger "github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
	"github.com/woshidama323/config"
	"google.golang.org/grpc"
)

var rlog = logger.New()

func main() {
	app := &cli.App{
		Name:  "robotclient",
		Usage: "can config info for robot",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "update",
				Value: false,
				Usage: "update config for system",
			},
			&cli.BoolFlag{
				Name:  "approve",
				Value: false,
				Usage: "approve for a token",
			},
		},
		Action: func(ctx *cli.Context) error {
			rlog.Info("communicating to robot server:", ctx.Args())

			if ctx.NArg() <= 0 {

				return errors.New("Wrong parameter number")
			}
			var conn *grpc.ClientConn
			conn, err := grpc.Dial(":9000", grpc.WithInsecure())
			if err != nil {
				rlog.Fatalln("failed to dial server,err:", err)
			}

			defer conn.Close()

			newconn := config.NewConfigUpdateClient(conn)
			if ctx.Bool("update") {
				message := config.Message{
					Body: "updateconfig",
				}

				response, err := newconn.ReloadConfig(context.Background(), &message)
				if err != nil {
					rlog.Errorf("Error when calling reloadconfig,err:%s\n", err)
					return err
				}
				rlog.Info("Response from server:", response, " err:", err)
				return nil
			}

			if ctx.Bool("approve") {
				message := config.Message{
					Body: "approve",
				}

				response, err := newconn.ApprovalToOneSplitAudit(ctx.Context, &message)
				if err != nil {
					rlog.Errorf("Error when calling reloadconfig,err:%s\n", err)
					return err
				}
				rlog.Info("Response from server:", response)
				return nil
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		rlog.Fatal(err)
	}
}
