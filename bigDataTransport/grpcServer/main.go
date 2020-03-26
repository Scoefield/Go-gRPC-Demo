package main

import (
	"bufio"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	pb "practiceProject/grpcDemo/bigDataTransport/proto"
)

type chunkByte struct {

}

func (c chunkByte) GetZipFile(in *pb.FileRequest, srv pb.ChunkSrv_GetZipFileServer) error {
	fileData := &pb.FileResponse{}
	file, err := os.Open("C:/Users/sangfor/go/src/practiceProject/grpcDemo/bigDataTransport/data/cxg.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	oneCap := make([]byte, 1024)
	for {
		n, err := reader.Read(oneCap)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		fileData.Filedata = oneCap[:n]
		if err := srv.Send(fileData); err != nil {
			panic(err)
		}
	}

	return nil
}

func (c chunkByte) GetChunk(in *pb.ChunkRequest, srv pb.ChunkSrv_GetChunkServer) error {
	chunk := &pb.ChunkResponse{}
	file, err := os.Open("C:/Users/sangfor/go/src/practiceProject/grpcDemo/bigDataTransport/data/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		chunk.Data = []byte(str)
		if err := srv.Send(chunk); err != nil {
			panic(err)
		}
	}
	return nil
}

func ReadFile() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterChunkSrvServer(s, chunkByte{})
	log.Println("serving on localhost:8888")
	log.Fatal(s.Serve(listen))
}
