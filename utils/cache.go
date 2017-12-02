/*
	并发非堵塞缓存
*/
package utils

type Func func(string) (interface{}, error)

type Memo struct{ requests chan request }

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

type request struct {
	key      string
	response chan<- result
}

func NewMemoryCache(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	// 启动一个 goroutines 服务
	go memo.server(f)
	return memo
}

func (m *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	// 使用通信来共享数据， 不要使用共享数据来通信
	m.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (m *Memo) Close() { close(memo.requests) }

//
func (m *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range m.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}
