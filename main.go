package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/johnaoss/htpasswd/apr1"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

func init() {

}

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

func CodeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "code",
		Short: "Generate Random Code",
		Run: func(cmd *cobra.Command, args []string) {
			alphabet := "ABCDEFGHJKLMNPQRSTUVWXYZ"
			number := "23456789"
			fmt.Println(RandStr(alphabet, 1) + RandStr(number, 2) + RandStr(alphabet, 1))
		},
	}
}

func HTPasswordCommand() *cobra.Command {
	var username string
	var password string
	var hashType string

	m := &cobra.Command{
		Use:   "htpasswd",
		Short: "Generate Htpasswd Secret Key",
		Run: func(cmd *cobra.Command, args []string) {

			if username == "" {
				fmt.Println("Please provide a username")
				os.Exit(1)
			}

			if password == "" {
				password = RandStr(BaseYoutubeUniqueIDAlphabet, 11)
			}

			var htpasswd string

			switch hashType {
			case "md5":
				s, err := apr1.Hash(password, "")
				if err != nil {
					panic(err)
				}
				htpasswd = fmt.Sprintf("%v:%v", username, s)
			case "bcrypt":
				generateApachePassword := func(username, password string) (string, error) {
					hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
					if err != nil {
						return "", err
					}
					return fmt.Sprintf("%s:%s", username, hashedPassword), nil
				}
				apachePassword, err := generateApachePassword(username, "123456")
				if err != nil {
					log.Fatalf("Error generating password: %v", err)
				}
				htpasswd = apachePassword
			default:
				fmt.Println("Unsupported hash type: " + hashType)
				os.Exit(1)
			}

			// output
			fmt.Println("Hash Type: " + hashType)
			fmt.Println("Username: " + username)
			fmt.Println("Password: " + password)
			fmt.Println("htpasswd: " + htpasswd)
		},
	}

	m.PersistentFlags().StringVarP(&username, "username", "u", "", "Username to use for authentication")
	m.PersistentFlags().StringVarP(&password, "password", "p", "", "Password to use for authentication")
	m.PersistentFlags().StringVarP(&hashType, "type", "t", "md5", "Hash type")

	return m
}

func PinCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "pin",
		Short: "Generate Pin Secret Key",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(RandStr("0123456789", 6))
		},
	}
}

func SerialNumberCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "sn",
		Short: "Generate SerialNumber Secret Key",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(RandStr("0123456789ABCDEFGHJKLMNPQRSTUVWXYZ", 10))
		},
	}
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
	mainCmd.AddCommand(CodeCommand())
	mainCmd.AddCommand(HTPasswordCommand())
	mainCmd.AddCommand(PinCommand())
	mainCmd.AddCommand(SerialNumberCommand())
	if err := mainCmd.Execute(); err != nil {
		panic(err)
	}
}
