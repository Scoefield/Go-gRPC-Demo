package main

import (
	"bufio"
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	pb "practiceProject/grpcDemo/bigDataTransport/proto"
)

func GetChunkData(client pb.ChunkSrvClient) {
	stream, err := client.GetChunk(context.Background(), &pb.ChunkRequest{Md5:"hello"})
	if err != nil {
		log.Fatal(err)
	}
	var blob []byte
	for {
		c, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Printf("Transfer of %d bytes successful", len(blob))
				return
			}
			log.Fatal(err)
		}
		blob = append(blob, c.Data...)
	}
}

func GetZipFileData(client pb.ChunkSrvClient) {
	fileName := "victory.zip"
	stream, err := client.GetZipFile(context.Background(), &pb.FileRequest{Filename:fileName})
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(fileName, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for {
		data, err := stream.Recv()
		if err != nil && err != io.EOF {
			panic(err)
		}
		if data != nil {
			nn, err := writer.Write(data.Filedata)
			if err != nil {
				panic(err)
			}
			if nn == 0 {
				break
			}
		}else {
			break
		}
	}
	_ = writer.Flush()
	log.Printf("Transport file %s success", fileName)
}

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewChunkSrvClient(conn)

	//GetChunkData(client)
	GetZipFileData(client)
}