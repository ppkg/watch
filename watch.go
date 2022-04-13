package watch

type Watcher interface {
	watch(string) chan *KV
}

type watcher struct {
	er Watcher
}

type KV struct {
	Key   string
	Value []byte
}

func (c *watcher) Watch(key string) chan *KV {
	return c.er.watch(key)
}

func NewWatcher(er Watcher) *watcher {
	return &watcher{
		er: er,
	}
}
