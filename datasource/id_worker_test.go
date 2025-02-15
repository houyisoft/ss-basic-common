package datasource

import (
	"fmt"
	"sync"
	"testing"
)

func TestWorker_GetId(t *testing.T) {
	type fields struct {
		mu        sync.Mutex
		timestamp int64
		workerId  int64
		number    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Worker{
				mu:        tt.fields.mu,
				timestamp: tt.fields.timestamp,
				workerId:  tt.fields.workerId,
				number:    tt.fields.number,
			}
			if got := w.GetId(); got != tt.want {
				t.Errorf("Worker.GetId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetId(t *testing.T) {
	// 测试脚本

	// 生成节点实例
	worker, err := NewWorker(1)

	if err != nil {
		fmt.Println(err)
		return
	}

	ch := make(chan int64)
	count := 10000
	// 并发 count 个 goroutine 进行 snowflake ID 生成
	for i := 0; i < count; i++ {
		go func() {
			id := worker.GetId()
			ch <- id
		}()
	}

	defer close(ch)

	m := make(map[int64]int)
	for i := 0; i < count; i++ {
		id := <-ch
		// 如果 map 中存在为 id 的 key, 说明生成的 snowflake ID 有重复
		_, ok := m[id]
		if ok {
			t.Error("ID is not unique!\n")
			return
		}
		// 将 id 作为 key 存入 map
		m[id] = i
	}
	// 成功生成 snowflake ID
	fmt.Println("All", count, "snowflake ID Get successed!")
}
