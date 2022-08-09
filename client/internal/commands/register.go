package commands

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/alexkopcak/gophkeeper/api-gateway/pkg/services/pb"
	"github.com/alexkopcak/gophkeeper/client/internal/client"
)

type RegisterCmd struct {
	Command *cobra.Command
}

func NewRegisterCmd(ctx context.Context, client *client.ServiceClient) *RegisterCmd {
	var name string
	var password string

	registerCmd := &cobra.Command{
		Use:   "register",
		Short: "Register new gophkeeper user",
		Long:  `This command can be used add new gophkeeper user`,

		Run: func(cmd *cobra.Command, args []string) {
			resp, err := client.Client.Register(ctx, &pb.RegisterRequest{
				UserName: name,
				Password: password,
			})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(resp)
			fmt.Println("register user func executed")
			fmt.Println("user", name)
			fmt.Println("password", password)
		},
	}

	registerCmd.Flags().StringVarP(&name, "user", "u", "", "new user name")
	registerCmd.MarkFlagRequired("user")

	registerCmd.Flags().StringVarP(&password, "password", "p", "", "new user password")
	registerCmd.MarkFlagRequired("password")

	return &RegisterCmd{
		Command: registerCmd,
	}
}
