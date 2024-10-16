package main

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/lightninglabs/lightning-terminal/firewalldb"
	"github.com/lightninglabs/lightning-terminal/litrpc"
	"github.com/urfave/cli"
)

var privacyMapCommands = cli.Command{
	Name:      "privacy",
	ShortName: "p",
	Usage: "Access the real-pseudo string pairs of the " +
		"privacy mapper",
	Category: "Privacy",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "session_id",
			Usage:    "The id of the session in question",
			Required: true,
		},
		cli.BoolFlag{
			Name: "realtopseudo",
			Usage: "set to true if the input should be " +
				"mapped to its pseudo counterpart. " +
				"Otherwise the input will be taken " +
				"as the pseudo value that should be " +
				"mapped to its real counterpart.",
		},
	},
	Subcommands: []cli.Command{
		privacyMapConvertStrCommand,
		privacyMapConvertUint64Command,
	},
}

var privacyMapConvertStrCommand = cli.Command{
	Name:      "str",
	ShortName: "s",
	Usage:     "convert a string to its real or pseudo counter part",
	Action:    privacyMapConvertStr,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "input",
			Usage:    "the string to convert",
			Required: true,
		},
	},
}

func privacyMapConvertStr(ctx *cli.Context) error {
	ctxb := context.Background()
	clientConn, cleanup, err := connectClient(ctx)
	if err != nil {
		return err
	}
	defer cleanup()
	client := litrpc.NewFirewallClient(clientConn)

	id, err := hex.DecodeString(ctx.GlobalString("session_id"))
	if err != nil {
		return err
	}

	resp, err := client.PrivacyMapConversion(
		ctxb, &litrpc.PrivacyMapConversionRequest{
			SessionId:    id,
			RealToPseudo: ctx.GlobalBool("realtopseudo"),
			Input:        ctx.String("input"),
		},
	)
	if err != nil {
		return err
	}

	printRespJSON(resp)

	return nil
}

var privacyMapConvertUint64Command = cli.Command{
	Name:      "uint64",
	ShortName: "u",
	Usage:     "convert a uint64 to its real or pseudo counter part",
	Action:    privacyMapConvertUint64,
	Flags: []cli.Flag{
		cli.Uint64Flag{
			Name:     "input",
			Usage:    "the uint64 to convert",
			Required: true,
		},
	},
}

func privacyMapConvertUint64(ctx *cli.Context) error {
	ctxb := context.Background()
	clientConn, cleanup, err := connectClient(ctx)
	if err != nil {
		return err
	}
	defer cleanup()
	client := litrpc.NewFirewallClient(clientConn)

	id, err := hex.DecodeString(ctx.GlobalString("session_id"))
	if err != nil {
		return err
	}

	input := firewalldb.Uint64ToStr(ctx.Uint64("input"))

	resp, err := client.PrivacyMapConversion(
		ctxb, &litrpc.PrivacyMapConversionRequest{
			SessionId:    id,
			RealToPseudo: ctx.GlobalBool("realtopseudo"),
			Input:        input,
		},
	)
	if err != nil {
		return err
	}

	output, err := firewalldb.StrToUint64(resp.Output)
	if err != nil {
		return err
	}

	printRespJSON(&litrpc.PrivacyMapConversionResponse{
		Output: fmt.Sprintf("%d", output),
	})
	return nil
}
