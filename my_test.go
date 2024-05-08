package test

import (
	"context"
	_ "embed"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
)

// lua脚本
//
//go:embed my.lua
var lua string

//编写一个 Go 程序，30 分钟内完成:
//模拟一个简单的多人参与的抢购活动。有100个商品，每个人最多购买1个商品，但是商品有限，只有 100 个，确保不会出现超卖现象。

func TestOne(t *testing.T) {
	//商品个数存redis
	//已购买的用户信息用redis存一个bool
	//
	var phone string //假设这里拿到了手机号
	var key string   //假设这里是商品的key
	redisclient := redis.NewClient(&redis.Options{
		Addr: "localhost:13306",
	})
	req, err := redisclient.Eval(context.Background(), lua, []string{key}, phone).Int()
	if err != nil || req != 0 {
		assert.True(t, false)
	}

}
