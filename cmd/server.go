package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "kvgo/api/v1"
	"log"
	"net"
)

type databaseServer struct {
	v1.UnimplementedDatabaseServiceServer
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "in-memory kv store server",
	Long:  `Start an in-memory zdb server`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Server starting...")
		log.Printf("Checking args: %v\n", args)

		// Listen on port
		lis, err := net.Listen("tcp", ":9999")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Printf("server listening at %v", lis.Addr())

		// Create Database
		//db.CreateDatabase()

		// Start grpc server
		s := grpc.NewServer()
		v1.RegisterDatabaseServiceServer(s, &databaseServer{})
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func (s *databaseServer) Set(context.Context, *v1.SetRequest) (*v1.SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (s *databaseServer) Get(context.Context, *v1.GetRequest) (*v1.GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
