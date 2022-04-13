package watch

import (
	"context"
	"time"

	"github.com/maybgit/glog"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdConfig clientv3.Config

func (c *EtcdConfig) watch(key string) chan *KV {
	ch := make(chan *KV)
	go func() {
		for {
			func() {
				if err := recover(); err != nil {
					glog.Error(err)
				}
				cli, err := clientv3.New(clientv3.Config(*c))
				if err != nil {
					glog.Error(err)
				}

				defer cli.Close()
				rch := cli.Watch(context.Background(), key)
				for wresp := range rch {
					for _, ev := range wresp.Events {
						ch <- &KV{Key: string(ev.Kv.Key), Value: ev.Kv.Value}
					}
				}
			}()
			time.Sleep(3 * time.Second)
		}
	}()
	return ch
}
