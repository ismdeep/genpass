package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func JWTCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "jwt",
		Short: "Generate JWT Secret Key",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(RandStr(BaseWithoutFuzzy, 32))
		},
	}
}

func GAuthCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "gauth",
		Short: "Generate GAuth Secret Key",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(RandStr("ABCDEFGHIJKLMNOPQRSTUVWXYZ234567", 32))
		},
	}
}

func HexCommand() *cobra.Command {
	var n int64
	m := &cobra.Command{
		Use:   "hex",
		Short: "Generate Hex Secret Key",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(RandHex(n))
		},
	}

	m.PersistentFlags().Int64VarP(&n, "length", "n", 32, "Length of secret key")

	return m
}

func UUIDCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "uuid",
		Short: "Generate UUID Secret Key",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(uuid.NewString())
		},
	}
}

func IDCommand() *cobra.Command {
	var youtubeStyle bool
	var uuidStyle bool
	m := &cobra.Command{
		Use:   "id",
		Short: "Generate Unique ID",
		Run: func(cmd *cobra.Command, args []string) {
			switch {
			case youtubeStyle:
				fmt.Println(RandStr(BaseYoutubeUniqueIDAlphabet, 11))
			case uuidStyle:
				fmt.Println(uuid.NewString())
			}
		},
	}

	m.PersistentFlags().BoolVar(&youtubeStyle, "youtube", false, "Use youtube style id")
	m.PersistentFlags().BoolVar(&uuidStyle, "uuid", false, "Use uuid style id")

	return m
}

func MainCommand() *cobra.Command {
	return &cobra.Command{
		Use:           "genpass",
		Short:         "Generate password",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
}

func main() {
	mainCmd := MainCommand()
	mainCmd.AddCommand(JWTCommand())
	mainCmd.AddCommand(GAuthCommand())
	mainCmd.AddCommand(HexCommand())
	mainCmd.AddCommand(UUIDCommand())
	mainCmd.AddCommand(IDCommand())
	if err := mainCmd.Execute(); err != nil {
		panic(err)
	}
}
