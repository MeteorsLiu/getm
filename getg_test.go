package getm

import (
	"sync"
	"testing"
	"time"
)

func TestGetM(t *testing.T) {
	t.Log(GetM(), MID(), GetG(), mOffset, mIDOffset)
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			_ = i + 1
			time.Sleep(5 * time.Second)
		}()
	}
	t.Log(GetM(), MID(), GetG(), mOffset, mIDOffset)
	wg.Wait()
	t.Log(GetM(), MID(), GetG(), mOffset, mIDOffset)
	t.Log(CustomInG[uint64]("goid"))
	SetCustomInG[uint64]("goid", 123456)
	t.Log(CustomInG[uint64]("goid"))
}
