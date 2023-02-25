package cmd

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	//v1 "zdb/api/v1"
	"log"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "client for zdb",
	Long:  "Embedded client for zdb client",
	Run: func(cmd *cobra.Command, args []string) {
		//addr1 := flag.String("address", "localhost:91019",
		//	"The server address in the format of host:port")
		addr := config.Client.ServerAddress

		var opts []grpc.DialOption

		log.Printf("Will attempt to connect to address: %v\n", addr)

		conn, err := grpc.Dial(addr, opts...)
		if err != nil {
			log.Fatalf("fail to dial: %v", err)
		}
		defer conn.Close()

		//client := v1.NewDatabaseServiceClient(conn)

		//client.Set(),

	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	clientCmd.Flags().String("server", "localhost:9999", "")
}
