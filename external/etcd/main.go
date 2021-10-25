package main

import (
	"context"
	"encoding/json"
	"fmt"
	etcd_client "github.com/coreos/etcd/clientv3"
	"time"
)

func demo() {
	/**
	创建一个 etcd 的客户端
	*/
	cli, err := etcd_client.New(etcd_client.Config{
		// endpoints 是 etcd 的集群的服务端口，也就是 host+port，可以换成域名，这样减少当 ip 更换造成的麻烦
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed err:", err)
	}
	fmt.Println("connect succ")

	defer cli.Close()

	// 往 etcd 里写数据
	put(cli)

	// 从 etcd 里读数据
	get(cli)

	// 监听 etcd 里面的数据变更（特别是作为配置中心的使用场景时）
	watch(cli)
}

func put(cli *etcd_client.Client) {
	/**
	在 ctx 背景下往 etcd 里面存值
	*/
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := cli.Put(ctx, "/logagent/conf/", "sample_value")
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}
}

func get(cli *etcd_client.Client) {
	/**
	在 ctx 背景下从 etcd 里面取值
	*/
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "/logagent/conf/")
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}

func watch(cli *etcd_client.Client) {
	for {
		rch := cli.Watch(context.Background(), "/logagent/conf/")
		for wresp := range rch {
			for _, ev := range wresp.Events {
				fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		}
	}
}

type LogConf struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

const (
	EtcdKey = "/oldboy/backend/logagent/192.168.3.49"
)

func testCreateData() {
	cli, err := etcd_client.New(etcd_client.Config{
		// endpoints 是 etcd 的集群的服务端口，也就是 host+port，可以换成域名，这样减少当 ip 更换造成的麻烦
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed err:", err)
	}
	fmt.Println("connect succ")

	defer cli.Close()

	/**
	在 ctx 背景下往 etcd 里面存值
	*/
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	logConfArr := []LogConf{
		{
			Path:  "/etc/nginx/logs/access.log",
			Topic: "nginx_access_log",
		},
		{
			Path:  "/etc/nignx/logs/error.log",
			Topic: "nginx_error_log",
		},
	}
	data, err := json.Marshal(logConfArr)
	if err != nil {
		fmt.Println("json failed,", err)
		return
	}
	_, err = cli.Put(ctx, EtcdKey, string(data))
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}

	/**
	在 ctx 背景下从 etcd 里面取值
	*/
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, EtcdKey)
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}

func main() {
	// 简单的 etcd 使用 demo
	//demo()

	// 测试创建数据到 etcd
	testCreateData()
}
