package main

import (
	"os"
	"testing"
	"time"
)

func TestMain(t *testing.M) {
	SetUp()
	os.Exit(t.Run())
}

// 测试brpop命令
func TestBRPop(t *testing.T) {
	client, err := GetClient()
	if err != nil {
		t.Fatal(err)
	}

	for {

		t.Log(time.Now(), "blocking pop")
		result, err := client.RandomR().BRPop(time.Second*5, "list_test").Result()
		if err != nil {
			t.Error(err.Error())
			continue
		}

		for _, v := range result {
			t.Log(v)
		}
	}
}

// 测试rpop命令
func TestRPop(t *testing.T) {
	client, err := GetClient()
	if err != nil {
		t.Fatal(err)
	}

	tick := time.NewTicker(time.Second * 5)
	timer := time.NewTimer(time.Second * 60)
	done := make(chan struct{})

	for {
		select {
		case <-tick.C:
			result, err := client.RandomR().RPop("list_test").Result()
			if err != nil {
				t.Error(err.Error())
				continue
			}

			t.Log(result)
		case <-timer.C:
			close(done)

		case <-done:
			t.Log("over")
			return
		}
	}
}
