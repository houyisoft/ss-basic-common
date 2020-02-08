package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
)

func callback1(master *redis.Client) error {
	it := master.Scan(0, "", 1).Iterator()
	for it.Next() {

		fmt.Println(it.Val())
	}
	return nil
}


func TestShow(t *testing.T)  {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"47.100.127.159:8001", "47.100.127.159:8002","47.100.127.159:8003","47.100.127.159:8004","47.100.127.159:8005","47.100.127.159:8006"},
	})
	err := client.ForEachMaster(callback1)
	if err != nil {
		fmt.Println(err.Error())
	}

	client.LPush("list","one","two","three") //rpush则在尾部插入
	client.LRem("list",2,"three") //删除list中前2个value为 ‘three’的元素
	client.LPop("list") //删除头部的值，同理RPop删除尾部的值。
	list, _ := client.LRange("list", 0, 2).Result()
	fmt.Println("List: ", list)
	//output:
	//List:  [three two one]

	client.LRem("list",1,"three")   //删除一个“three”
	list1, _ := client.LRange("list", 0, 2).Result()
	fmt.Println("List: ", list1)  //从输出发现list中只剩下two 和one，遍历完list后又从头开始遍历再输出一个two

	//hash
	user := make(map[string]interface{})
	user["name"] = "jim"
	user["gender"] = "man"
	user["age"] = 23
	client.HMSet("user",user)


	client.HSet("user", "name","tom")
	name := client.HGet("user","name")
	fmt.Print(name)

	hash, _ := client.HGetAll("user").Result()
	for k, v:= range hash{
		fmt.Printf("key: %v, value: %v ",k, v)
	}

	client.SAdd("set", 7, 6, 5, 3)

	count1 := client.SCard("set")
	fmt.Println("count1:", count1)


	nums:= client.SMembers("set")
	fmt.Println("Set:", nums)

	result, _ := client.Keys("*").Result()
	fmt.Println("Redis value: ", result)

	pl := client.Pipeline()
	pl.Set("pipe", 0,0)
	pl.Incr("pipe")
	pl.Incr("pipe")
	pl.Incr("pipe")
	pl.Exec()
	p11,_ := client.Get("pipe").Result()
	fmt.Println("Pipe: ",p11)


	done := make(chan struct{})
	client.Publish("mychannel", "hello budy!\n")
	go func() {
		pubsub := client.Subscribe("mychannel")
		msg,_ := pubsub.Receive()
		fmt.Println("Receive from channel:", msg)
		done <- struct {}{}
	}()

}