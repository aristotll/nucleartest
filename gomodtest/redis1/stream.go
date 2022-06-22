package redis1

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type Stream struct {
	client *redis.Client
}

func NewStream(cli *redis.Client) *Stream {
	return &Stream{client: cli}
}

func (s *Stream) ProduceMessage(ctx context.Context, key string, val any) (result string, err error) {
	cmd := s.client.XAdd(ctx, &redis.XAddArgs{
		Stream: key,
		Values: val,
	})
	log.Println("cmd: ", printCmd(cmd))
	return cmd.Result()
}

type BlockArgs struct {
	NoBlock *struct {
		readStartId string
	}
	Block *struct {
		maxWaitTime time.Duration
	}
}

func (s *Stream) ConsumeMessage(ctx context.Context, keys []string, count int64, blockArgs *BlockArgs) ([]redis.XStream, error) {
	if blockArgs.NoBlock != nil && blockArgs.Block != nil {
		return nil, fmt.Errorf("only one of them (Noblock or Block) can be provided")
	}
	stream := make([]string, len(keys), len(keys))
	copy(stream, keys)
	readArgs := &redis.XReadArgs{
		Streams: stream,
		Count:   count,
	}
	if blockArgs.Block != nil {
		readArgs.Block = blockArgs.Block.maxWaitTime
		readArgs.Streams = append(readArgs.Streams, "$")
	}
	if blockArgs.NoBlock != nil {
		readArgs.Streams = append(readArgs.Streams, blockArgs.NoBlock.readStartId)
	}
	cmd := s.client.XRead(ctx, readArgs)
	log.Println("cmd: ", printCmd(cmd))
	return cmd.Result()
}

func printCmd(cmd redis.Cmder) string {
	return cmd.String()
}
