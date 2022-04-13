# watch
go get github.com/ppkg/watch
```
watcher := watch.NewWatcher(&watch.EtcdConfig{Endpoints: []string{"127.0.0.1:2379"}}) // etcd watch
// watcher := watch.NewWatcher(&watch.ConsulConfig{Address: "127.0.0.1:8500"})// consul watch

for v := range watcher.Watch("foo") {
    fmt.Println(v.Key, string(v.Value))
}
```