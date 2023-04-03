// Package main implements a server for movieinfo service.
package main

import (
	"context"
	"fmt"
	"log"
	"movieapiserver/protogenfiles"
	"net"
	"strconv"
	"strings"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement .MovieInfoServer
type server struct {
	protogenfiles.UnimplementedMovieInfoServer
}

// Map representing a database
var moviedb = map[string][]string{"Pulp fiction": []string{"1994", "Quentin Tarantino", "John Travolta,Samuel Jackson,Uma Thurman,Bruce Willis"}}

// GetMovieInfo implements .MovieInfoServer
func (s *server) GetMovieInfo(ctx context.Context, in *protogenfiles.MovieRequest) (*protogenfiles.MovieReply, error) {
	title := in.GetTitle()
	log.Printf("Received: %v", title)
	reply := &protogenfiles.MovieReply{}
	if val, ok := moviedb[title]; !ok { // Title not present in database
		return reply, nil
	} else {
		if year, err := strconv.Atoi(val[0]); err != nil {
			reply.Year = -1
		} else {
			reply.Year = int32(year)
		}
		reply.Director = val[1]
		cast := strings.Split(val[2], ",")
		reply.Cast = append(reply.Cast, cast...)

	}

	return reply, nil

}

func (s *server) SetMovieInfo(c context.Context, md *protogenfiles.MovieData) (*protogenfiles.Status, error) {
	fmt.Println("Set Movie Info Method Invoked..........")

	// Setting the data in the map above ----
	// creating slice
	var slc []string
	slc = append(slc, md.Title, md.Year, md.Director, md.Cast)
	moviedb[md.Title] = slc

	return &protogenfiles.Status{Code: "Success"}, nil

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protogenfiles.RegisterMovieInfoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
