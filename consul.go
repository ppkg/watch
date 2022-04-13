package watch

import (
	"time"

	"github.com/maybgit/glog"

	consulapi "github.com/hashicorp/consul/api"
)

type ConsulConfig consulapi.Config

func (c *ConsulConfig) watch(key string) chan *KV {
	ch := make(chan *KV)
	go func() {
		var lastIndex uint64
		for {
			func() {
				if err := recover(); err != nil {
					glog.Error(err)
				}

				if c.Scheme == "" {
					c.Scheme = "http"
				}

				if client, err := consulapi.NewClient((*consulapi.Config)(c)); err != nil {
					glog.Error(err)
				} else {
					pair, meta, err := client.KV().Get(key, &consulapi.QueryOptions{WaitIndex: lastIndex})
					if err != nil {
						glog.Error(err)
					} else {
						if pair != nil {
							lastIndex = meta.LastIndex
							ch <- &KV{Key: pair.Key, Value: pair.Value}
						}
					}
				}
			}()
			time.Sleep(3 * time.Second)
		}
	}()
	return ch
}
