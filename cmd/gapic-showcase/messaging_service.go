// Code generated. DO NOT EDIT.

package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gapic "github.com/googleapis/gapic-showcase/client"
)

var MessagingConfig *viper.Viper
var MessagingClient *gapic.MessagingClient
var MessagingSubCommands []string = []string{
	"create-room",
	"get-room",
	"update-room",
	"delete-room",
	"list-rooms",
	"create-blurb",
	"get-blurb",
	"update-blurb",
	"delete-blurb",
	"list-blurbs",
	"search-blurbs",
	"poll-search-blurbs", "stream-blurbs",
	"send-blurbs",
	"connect",
}

func init() {
	rootCmd.AddCommand(MessagingServiceCmd)

	MessagingConfig = viper.New()
	MessagingConfig.SetEnvPrefix("GAPIC-SHOWCASE_MESSAGING")
	MessagingConfig.AutomaticEnv()

	MessagingServiceCmd.PersistentFlags().Bool("insecure", false, "Make insecure client connection. Or use GAPIC-SHOWCASE_MESSAGING_INSECURE. Must be used with \"address\" option")
	MessagingConfig.BindPFlag("insecure", MessagingServiceCmd.PersistentFlags().Lookup("insecure"))
	MessagingConfig.BindEnv("insecure")

	MessagingServiceCmd.PersistentFlags().String("address", "", "Set API address used by client. Or use GAPIC-SHOWCASE_MESSAGING_ADDRESS.")
	MessagingConfig.BindPFlag("address", MessagingServiceCmd.PersistentFlags().Lookup("address"))
	MessagingConfig.BindEnv("address")

	MessagingServiceCmd.PersistentFlags().String("token", "", "Set Bearer token used by the client. Or use GAPIC-SHOWCASE_MESSAGING_TOKEN.")
	MessagingConfig.BindPFlag("token", MessagingServiceCmd.PersistentFlags().Lookup("token"))
	MessagingConfig.BindEnv("token")

	MessagingServiceCmd.PersistentFlags().String("api_key", "", "Set API Key used by the client. Or use GAPIC-SHOWCASE_MESSAGING_API_KEY.")
	MessagingConfig.BindPFlag("api_key", MessagingServiceCmd.PersistentFlags().Lookup("api_key"))
	MessagingConfig.BindEnv("api_key")
}

var MessagingServiceCmd = &cobra.Command{
	Use:       "messaging",
	Short:     "A simple messaging service that implements chat...",
	Long:      "A simple messaging service that implements chat rooms and profile posts.   This messaging service showcases the features that API clients  generated...",
	ValidArgs: MessagingSubCommands,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		var opts []option.ClientOption

		address := MessagingConfig.GetString("address")
		if address != "" {
			opts = append(opts, option.WithEndpoint(address))
		}

		if MessagingConfig.GetBool("insecure") {
			if address == "" {
				return fmt.Errorf("Missing address to use with insecure connection")
			}

			conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				return err
			}
			opts = append(opts, option.WithGRPCConn(conn))
		}

		if token := MessagingConfig.GetString("token"); token != "" {
			opts = append(opts, option.WithTokenSource(oauth2.StaticTokenSource(
				&oauth2.Token{
					AccessToken: token,
					TokenType:   "Bearer",
				})))
		}

		if key := MessagingConfig.GetString("api_key"); key != "" {
			opts = append(opts, option.WithAPIKey(key))
		}

		MessagingClient, err = gapic.NewMessagingClient(ctx, opts...)
		return
	},
}
