package main

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// RegisterServerToETCD 将服务注册到etcd上
func RegisterServerToETCD(serviceTarget string, value string) {
	//dir := strings.TrimRight(serviceTarget, "/") + "/"
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		log.Fatalln(err)
	}

	kv := clientv3.NewKV(cli)
	lease := clientv3.NewLease(cli)

	var curLeaseId clientv3.LeaseID

	for {
		// 还没有租约
		if curLeaseId == 0 {
			// 创建一个新租约
			leaseResp, err := lease.Grant(context.TODO(), 10)
			if err != nil {
				log.Fatalln(err)
			}
			curLeaseId = leaseResp.ID
			key := serviceTarget + fmt.Sprintf("%d", curLeaseId)

			_, err = kv.Put(
				context.TODO(),
				key, value,
				clientv3.WithLease(leaseResp.ID))
			if err != nil {
				log.Fatalln(err)
			}

		} else {
			_, err := lease.KeepAliveOnce(context.TODO(), curLeaseId)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
