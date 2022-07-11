package main

import (
	"context"
	"os"
	"testing"
	"valid/types/user"

	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

var (
	ctx = context.Background()
)

// NewGrpcClient 创建indicator rpc客户端
func NewGrpcClient() user.UserClient {
	client, err := zrpc.NewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts:              []string{"192.168.2.230:2379"},
			Key:                "user.rpc",
			InsecureSkipVerify: false,
		},
		Timeout: 2000,
	})
	if err != nil {
		os.Exit(0)
	}
	cc := client.Conn()
	ic := user.NewUserClient(cc)
	return ic
}

func TestCreate(t *testing.T) {
	ic := NewGrpcClient()
	want := "Ok"

	cr, err := ic.Create(ctx, &user.CreateRequest{
		Name:  "yother",
		Age:   20,
		Email: "yother233@gmail.com",
	})
	if err != nil {
		t.Error(err.Error())
	}

	t.Logf("response:%+v", cr)

	assert.Equal(t, want, cr.GetMsg())
}
