package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"ss-basic-common/datastructure/circlequeue"
	"ss-basic-common/datastructure/skiplist"
	"ss-basic-common/datastructure/timeround/simple"
	"ss-basic-common/datastructure/tree/rbtree"
	"ss-basic-common/datastructure/tree/trietree"
)

func skipList() {
	intCompare := func(a interface{}, b interface{}) bool {
		return a.(int) < b.(int)
	}
	l := skiplist.New(intCompare)

	// insert
	l.Insert(3, "value 3")
	l.Insert(1, "value 1")
	l.Insert(2, "value 2")

	// delete
	l.Delete(2)

	// get
	fmt.Println(l.Search(1))

	// foreach
	l.Foreach(func(key interface{}, value interface{}) {
		fmt.Println(key, ":", value)
	})
}

func cycleQueue() {
	//初始化一个队列
	queue := &circlequeue.CircleQueue{
		MaxSize: 5,
		Head:    0,
		Tail:    0,
	}
	var key string
	var val int
	for {
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从队列获取数据")
		fmt.Println("3. 输入show 表示显示队列")
		fmt.Println("4. 输入exit 表示显示队列")

		fmt.Print("请输入:")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {

				fmt.Println("加入队列ok")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数=", val)
			}
		case "show":
			queue.ListQueue()
		case "exit":
			os.Exit(0)
		}
	}

}

func rbTree(){
	var rbt rbtree.RBTree
	rbt.Add(9)
	rbt.Add(8)
	rbt.Add(7)
	rbt.Add(6)
	rbt.Add(5)
	rbt.Add(4)
	rbt.Add(3)
	rbt.Add(2)
	rbt.Add(1)
	fmt.Println(rbt.GetDeepth())
}

func TrieTree()  {
	t := trietree.Constructor()
	t.Insert("Hello")
	t.Insert("Hell")
	fmt.Print(t.Search("Hello"), "\n")
	fmt.Print(t.Search("Hell"), "\n")
	fmt.Print(t.Search("Hallo"), "\n")
}

func callback1(args interface{}) {
	//只执行一次的事件
	if values, ok := args.([]string); ok {
		var str1 string = values[0]
		var str2 string = values[1]
		log.Println("callback1(" + str1 + "," + str2 + ")")
	} else {
		log.Println("callback1()")
	}
}

func callback2(args interface{}) {
	//每次在当前时间点之后5s插入一个定时器，这样就能形成每隔5秒调用一次callback2回调函数，可以用于周期性事件
	simple.SetTimer("callback2", 5, callback2, args)
	log.Println("callback2")
}

func TimeRound(){
	fmt.Println("offset=%v",1)

	// cpu多核
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 定时器1，传入两个参数
	simple.SetTimer("callback1", 3, callback1, []string{"hello", "world"})
	simple.SetTimer("callback1", 4, callback1, []string{"hello", "worldword"})
	simple.SetTimer("callback1", 257, callback1, []string{"hello", "worldwordword"})
	// 定时器2，不传参数
	simple.SetTimer("callback2", 6, callback2, nil)
	// 移除定时器
	//simple.Delete(simple.TimerMap["callback2"])
	//运行计时器
	simple.Run()
}

func main() {
	TimeRound()
}
