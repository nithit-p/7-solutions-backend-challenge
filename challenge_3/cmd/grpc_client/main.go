package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nithit-p/7-solutions-backend-challenge/challenge_3/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBeefServiceClient(conn)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		raw, _ := reader.ReadString('\n')
		input := raw[:len(raw)-2]
		if input == "q" {
			break
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		beefSummary, err := c.GetBeefSummary(ctx, &pb.BeefSummaryRequest{Text: input})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		jsonBytes, err := json.Marshal(beefSummary)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("GetBeefSummary: %s", string(jsonBytes))
	}
}
