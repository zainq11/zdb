package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	v1 "zdb/api/v1"
	"zdb/pkg/config"

	//v1 "zdb/api/v1"
	"log"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "client for zdb",
	Long:  "Embedded client for zdb client",
	Run: func(cmd *cobra.Command, args []string) {
		addr := config.GetServerAddress()

		// Connect to server
		log.Printf("Will attempt to connect to address: %v\n", addr)
		conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		defer conn.Close()
		if err != nil {
			log.Fatalf("fail to dial: %v", err)
		}
		log.Printf("Connection established: %v\n", conn.GetState().String())

		// Create grpc client
		client := v1.NewDatabaseServiceClient(conn)

		// SET
		k := "foo"
		v := "bar"
		req := v1.SetRequest{Key: k, Value: v}
		resSet, err := client.Set(context.Background(), &req)
		if err != nil {
			log.Fatalf("An error occured after calling SET: %v", err)
		}
		log.Printf("SET succeeded, key: %v, value: %v \n", resSet.Key, resSet.Value)

		// GET
		resGet, err := client.Get(context.Background(), &v1.GetRequest{
			Key: k,
		})
		if err != nil {
			log.Fatalf("An error occured after calling GET: %v", err)
		}
		log.Printf("GET succeeded, key: %v, value: %v \n", resGet.Key, resGet.Value)
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
