// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	// models "qcode/models"

	// pb "qcode/rpc"

	"github.com/spf13/cobra"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/reflection"
)

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Serve the API through GRPC",
	Long:  `Starts GRPC server for clients with grpc connection`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GRPC Server started at port: 2012")

		// lis, err := net.Listen("tcp", ":8080")
		// if err != nil {
		// 	log.Fatalf("failed to listen: %v", err)
		// }
		// s := grpc.NewServer()
		// pb.RegisterUserServiceHandler(s, &models.User{})
		// // Register reflection service on gRPC server.
		// reflection.Register(s)
		// if err := s.Serve(lis); err != nil {
		// 	log.Fatalf("failed to serve: %v", err)
		// }
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grpcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
