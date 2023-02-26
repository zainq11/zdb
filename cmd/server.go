package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	v1 "zdb/api/v1"
	"zdb/pkg/db"
)

type databaseServer struct {
	v1.UnimplementedDatabaseServiceServer
	db db.Database
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "in-memory kv store server",
	Long:  `Start a zdb server`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Server starting...")
		log.Printf("Checking args: %v\n", args)

		// Listen on port
		lis, err := net.Listen("tcp", ":9999")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Printf("server listening at %v", lis.Addr())

		// Create DB server
		dbServer := databaseServer{db: db.CreateDatabase("zdb")}

		// Start grpc server
		gServer := grpc.NewServer()
		v1.RegisterDatabaseServiceServer(gServer, &dbServer)
		if err := gServer.Serve(lis); err != nil {
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

func (s *databaseServer) Set(_ context.Context, req *v1.SetRequest) (*v1.SetResponse, error) {
	log.Printf("Received SET command for key: %s, value: %s", req.Key, req.Value)

	err := s.db.Set(req.Key, req.Value)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "SET error")
	}

	log.Printf("Processed SET command for key: %s, value: %s", req.Key, req.Value)
	return &v1.SetResponse{Key: req.Key, Value: req.Value}, err
}
func (s *databaseServer) Get(_ context.Context, req *v1.GetRequest) (*v1.GetResponse, error) {
	log.Printf("Processed GET command for key: %s", req.Key)

	v, err := s.db.Get(req.Key)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "GET error")
	}

	log.Printf("Processed GET command for key: %s, value: %s", req.Key, v)
	return &v1.GetResponse{Key: req.Key, Value: v}, err
}
