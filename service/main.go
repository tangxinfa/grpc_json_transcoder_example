package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"sync"

	"./gen"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type KV struct {
	sync.Mutex
	store map[string]string
}

func (k *KV) Get(ctx context.Context, in *kv.GetRequest) (*kv.GetResponse, error) {
	log.Printf("get: %s", in.Key)
	resp := new(kv.GetResponse)
	if val, ok := k.store[in.Key]; ok {
		resp.Value = val
	}

	return resp, nil
}

func (k *KV) Set(ctx context.Context, in *kv.SetRequest) (*kv.SetResponse, error) {
	log.Printf("set: %s = %s", in.Key, in.Value)
	k.Lock()
	defer k.Unlock()

	k.store[in.Key] = in.Value

	return &kv.SetResponse{true}, nil
}

func (k *KV) Count(in *kv.CountRequest, stream kv.KV_CountServer) error {
	var count int
	var err error

	for i := 0; i < 10; i += 1 {
		k.Lock()
		count = len(k.store)
		k.Unlock()
		log.Printf("Sending count[%d]: %d", i, count)
		stream.Send(&kv.CountResponse{Count: uint32(count)})
		log.Printf("Sent count[%d]: %d", i, count)
		time.Sleep(time.Second * 1)
	}

	log.Printf("Sent count done")

	return err
}

func NewKVStore() (kv *KV) {
	kv = &KV{
		store: make(map[string]string),
	}

	return
}

func main() {
	port := flag.Int("port", 8081, "grpc port")

	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	gs := grpc.NewServer()
	kv.RegisterKVServer(gs, NewKVStore())

	log.Printf("starting grpc on :%d\n", *port)

	gs.Serve(lis)
}
