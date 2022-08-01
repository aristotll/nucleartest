package redis1

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"log"
	"time"
)

type Config struct {
	ShowCmd bool
}

type Stream struct {
	client *redis.Client
	cfg    *Config
}

func NewStream(cli *redis.Client, cfg *Config) *Stream {
	return &Stream{client: cli, cfg: cfg}
}

func (s *Stream) ProduceMessage(ctx context.Context, stream string, val any) (result string, err error) {
	cmd := s.client.XAdd(ctx, &redis.XAddArgs{
		Stream: stream,
		Values: val,
	})
	if s.cfg.ShowCmd {
		log.Println(printCmd(cmd))
	}
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

func (s *Stream) ConsumeMessage(ctx context.Context, streams []string, count int64, blockArgs *BlockArgs) ([]redis.XStream, error) {
	if blockArgs.NoBlock != nil && blockArgs.Block != nil {
		return nil, fmt.Errorf("only one of them (Noblock or Block) can be provided")
	}
	stream := make([]string, len(streams), len(streams))
	copy(stream, streams)
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
	if s.cfg.ShowCmd {
		log.Println(printCmd(cmd))
	}
	return cmd.Result()
}

func (s *Stream) CreatGroup(ctx context.Context, groupName, stream, startId string) error {
	cmd := s.client.XGroupCreate(ctx, stream, groupName, startId)
	if s.cfg.ShowCmd {
		log.Println(printCmd(cmd))
	}
	return cmd.Err()
}

func (s *Stream) GroupInfo(ctx context.Context, stream string) ([]redis.XInfoGroup, error) {
	cmd := s.client.XInfoGroups(ctx, stream)
	if s.cfg.ShowCmd {
		log.Println(cmd.String())
	}
	return cmd.Result()
}

func (s *Stream) GroupConsume(
	ctx context.Context,
	groupName, consumerName string, count int64,
	streams []string, block time.Duration, noAck bool,
) ([]redis.XStream, error) {
	cmd := s.client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    groupName,
		Consumer: consumerName,
		Streams:  streams,
		Count:    count,
		Block:    block,
		NoAck:    noAck,
	})
	return cmd.Result()
}

func printCmd(cmd redis.Cmder) string {
	return cmd.String()
}
