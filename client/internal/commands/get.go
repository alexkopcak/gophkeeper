package commands

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/alexkopcak/gophkeeper/api-gateway/pkg/services/pb"
	"github.com/alexkopcak/gophkeeper/client/internal/client"
)

type GetCmd struct {
	Command *cobra.Command
}

func NewGetCmd(ctx context.Context, cli *client.ServiceClient) *GetCmd {
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "get gophkeeper stored user data",
		Long:  `This command can be used to get user stored data at gophkeeper`,
	}

	var name string
	getCmd.PersistentFlags().StringVarP(&name, "user", "u", "", "user name")

	var password string
	getCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "user password")

	var res *pb.LoginResponse
	var err error

	anyType := &cobra.Command{
		Use:   "any",
		Short: "any type of stored data",
		Long:  `This command can be used to get any type of stored data`,
		PreRun: func(cmd *cobra.Command, args []string) {
			res, err = cli.Client.Login(ctx, &pb.LoginRequest{
				UserName: name,
				Password: password,
			})
			if err != nil {
				log.Fatal(err)
			}

			if res == nil {
				log.Fatal("something went wrong")
			}
			if res.Token == "" {
				log.Fatal("empty token")
			}
			fmt.Println("get any type prerun executed")
		},

		Run: func(cmd *cobra.Command, args []string) {
			qres, err := cli.Client.Query(context.WithValue(ctx, client.KeyPrincipalID, res.Token), &pb.QueryRequest{
				Type: pb.MessageType_ANY,
			})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(qres)
		},
	}

	binType := &cobra.Command{
		Use:   "bin",
		Short: "bin type of stored data",
		Long:  `This command can be used to get binnary type of stored data`,

		PreRun: func(cmd *cobra.Command, args []string) {
			res, err = cli.Client.Login(ctx, &pb.LoginRequest{
				UserName: name,
				Password: password,
			})
			if err != nil {
				log.Fatal(err)
			}

			if res == nil {
				log.Fatal("something went wrong")
			}
			if res.Token == "" {
				log.Fatal("empty token")
			}
			fmt.Println("get bin type prerun executed")
		},
		Run: func(cmd *cobra.Command, args []string) {
			qres, err := cli.Client.Query(context.WithValue(ctx, client.KeyPrincipalID, res.Token), &pb.QueryRequest{
				Type: pb.MessageType_BINNARY,
			})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(qres)
		},
	}

	cardType := &cobra.Command{
		Use:   "card",
		Short: "card type of stored data",
		Long:  `This command can be used to get card type of stored data`,

		PreRun: func(cmd *cobra.Command, args []string) {
			res, err = cli.Client.Login(ctx, &pb.LoginRequest{
				UserName: name,
				Password: password,
			})
			if err != nil {
				log.Fatal(err)
			}

			if res == nil {
				log.Fatal("something went wrong")
			}
			if res.Token == "" {
				log.Fatal("empty token")
			}
			fmt.Println("get card type prerun executed")
		},
		Run: func(cmd *cobra.Command, args []string) {
			qres, err := cli.Client.Query(context.WithValue(ctx, client.KeyPrincipalID, res.Token), &pb.QueryRequest{
				Type: pb.MessageType_CARD,
			})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(qres)
		},
	}

	credType := &cobra.Command{
		Use:   "cred",
		Short: "login pasword type of stored data",
		Long:  `This command can be used to get login password type of stored data`,

		PreRun: func(cmd *cobra.Command, args []string) {
			res, err = cli.Client.Login(ctx, &pb.LoginRequest{
				UserName: name,
				Password: password,
			})
			if err != nil {
				log.Fatal(err)
			}

			if res == nil {
				log.Fatal("something went wrong")
			}
			if res.Token == "" {
				log.Fatal("empty token")
			}
			fmt.Println("get login password type prerun executed")
		},
		Run: func(cmd *cobra.Command, args []string) {
			qres, err := cli.Client.Query(context.WithValue(ctx, client.KeyPrincipalID, res.Token), &pb.QueryRequest{
				Type: pb.MessageType_CARD,
			})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(qres)
		},
	}

	textType := &cobra.Command{
		Use:   "text",
		Short: "text type of stored data",
		Long:  `This command can be used to get text type of stored data`,

		PreRun: func(cmd *cobra.Command, args []string) {
			res, err = cli.Client.Login(ctx, &pb.LoginRequest{
				UserName: name,
				Password: password,
			})
			if err != nil {
				log.Fatal(err)
			}

			if res == nil {
				log.Fatal("something went wrong")
			}
			if res.Token == "" {
				log.Fatal("empty token")
			}
			fmt.Println("get text type prerun executed")
		},
		Run: func(cmd *cobra.Command, args []string) {
			qres, err := cli.Client.Query(context.WithValue(ctx, client.KeyPrincipalID, res.Token), &pb.QueryRequest{
				Type: pb.MessageType_CARD,
			})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(qres)
		},
	}

	getCmd.AddCommand(anyType)
	getCmd.AddCommand(binType)
	getCmd.AddCommand(cardType)
	getCmd.AddCommand(credType)
	getCmd.AddCommand(textType)

	return &GetCmd{
		Command: getCmd,
	}
}
