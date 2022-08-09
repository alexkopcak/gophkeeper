package commands

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/alexkopcak/gophkeeper/client/internal/client"
)

type AddCmd struct {
	Command *cobra.Command
}

func NewAddCmd(ctx context.Context, cli *client.ServiceClient) *AddCmd {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "add gophkeeper stored user data",
		Long:  `This command can be used to add user stored data to gophkeeper`,
	}

	// var name string
	// addCmd.PersistentFlags().StringVarP(&name, "user", "u", "", "user name")
	// //addCmd.MarkFlagRequired("user")

	// var password string
	// addCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "user password")
	// //addCmd.MarkFlagRequired("password")

	//var res *pb.LoginResponse
	// var err error

	// var key string

	// var id int64
	// var cardowner string
	// var cardnumber string
	// var carddate string
	// var cardcvv string
	// var meta string

	// cardTypeCmd := &cobra.Command{
	// 	Use:   "card",
	// 	Short: "card type of stored data",
	// 	Long:  `This command can be used to add card type of stored data`,
	// 	PreRun: func(cmd *cobra.Command, args []string) {
	// 		res, err = cli.Client.Login(ctx, &pb.LoginRequest{
	// 			UserName: name,
	// 			Password: password,
	// 		})
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		if res == nil {
	// 			log.Fatal("something went wrong")
	// 		}
	// 		if res.Token == "" {
	// 			log.Fatal("empty token")
	// 		}
	// 	},

	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		card := &models.Card{
	// 			Owner:  cardowner,
	// 			Number: cardnumber,
	// 			Date:   carddate,
	// 			CVV:    cardcvv,
	// 		}
	// 		rec := models.NewRecord(
	// 			id,
	// 			models.Add,
	// 			models.Tcard,
	// 			[]byte(key),
	// 			card,
	// 			meta,
	// 		)

	// 		ares, err := cli.Client.AddCommand(context.WithValue(ctx, client.KeyPrincipalID, res.Token),
	// 			&pb.CommandRequest{
	// 				Type: pb.MessageType_CARD,
	// 				Id:   rec.Id,
	// 				Data: rec.Value,
	// 				Meta: rec.Meta,
	// 			})
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		fmt.Println(rec)
	// 		fmt.Println(ares)

	// 		fmt.Println("add card type executed")
	// 	},
	// }

	// cardTypeCmd.Flags().StringVarP(&key, "key", "", "", "secret key")
	// //	cardTypeCmd.MarkFlagRequired("key")

	// cardTypeCmd.PerPersistentFlags().StringVarP(&cardowner, "cardowner", "", "", "card owner")
	// //	cardTypeCmd.MarkFlagRequired("cardowner")

	// cardTypeCmd.PersistentFlags().StringVarP(&cardnumber, "cardnumber", "", "", "card number")
	// //	cardTypeCmd.MarkFlagRequired("cardnumber")

	// cardTypeCmd.PersistentFlags().StringVarP(&carddate, "carddate", "", "", "card expires date")
	// //	cardTypeCmd.MarkFlagRequired("carddate")

	// cardTypeCmd.PersistentFlags().StringVarP(&cardcvv, "cardcvv", "", "", "card cvv code")
	// //	cardTypeCmd.MarkFlagRequired("cardcvv")

	// cardTypeCmd.Flags().StringVarP(&meta, "meta", "", "", "meta information")

	// var login string
	// var pass string

	// userPasswordTypeCmd := &cobra.Command{
	// 	Use:   "cred",
	// 	Short: "credential type of stored data",
	// 	Long:  `This command can be used to add credential type to stored data`,

	// 	PreRun: func(cmd *cobra.Command, args []string) {
	// 		res, err = cli.Client.Login(ctx, &pb.LoginRequest{
	// 			UserName: name,
	// 			Password: password,
	// 		})
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		if res == nil {
	// 			log.Fatal("something went wrong")
	// 		}
	// 		if res.Token == "" {
	// 			log.Fatal("empty token")
	// 		}
	// 	},

	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		cred := &models.LoginPass{
	// 			Login: login,
	// 			Pass:  pass,
	// 		}
	// 		rec := models.NewRecord(
	// 			id,
	// 			models.Add,
	// 			models.Tcard,
	// 			[]byte(key),
	// 			cred,
	// 			meta,
	// 		)

	// 		ares, err := cli.Client.AddCommand(context.WithValue(ctx, client.KeyPrincipalID, res.Token),
	// 			&pb.CommandRequest{
	// 				Type: pb.MessageType_LOGIN_PASSWORD,
	// 				Id:   rec.Id,
	// 				Data: rec.Value,
	// 				Meta: rec.Meta,
	// 			})
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		fmt.Println(rec)
	// 		fmt.Println(ares)

	// 		fmt.Println("add login/pass type executed")
	// 	},
	// }

	// userPasswordTypeCmd.Flags().StringVarP(&key, "key", "", "", "secret key")
	// userPasswordTypeCmd.MarkFlagRequired("key")

	// userPasswordTypeCmd.Flags().StringVarP(&login, "login", "", "", "stored login")
	// userPasswordTypeCmd.MarkFlagRequired("login")

	// userPasswordTypeCmd.Flags().StringVarP(&pass, "pass", "", "", "stored password")
	// userPasswordTypeCmd.MarkFlagRequired("pass")

	//addCmd.AddCommand(cardTypeCmd)
	//	addCmd.AddCommand(userPasswordTypeCmd)

	return &AddCmd{
		Command: addCmd,
	}
}
