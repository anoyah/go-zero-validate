# go-zero-validate

> proto验证参数类型
>
> 基于go-zero的[protoc-gen-validate](https://github.com/envoyproxy/protoc-gen-validate)使用



**手动调用参数校验**

```go
func (l *CreateLogic) Create(in *user.CreateRequest) (*user.CreateResponse, error) {
	if err := in.ValidateAll(); err != nil {
		return &user.CreateResponse{
			Msg: err.Error(),
		}, nil
	}

	return &user.CreateResponse{
		Msg: "Ok",
	}, nil
}
```

**中间件调用参数校验**

> 引用自：[validator.go](https://github.com/grpc-ecosystem/go-grpc-middleware/blob/master/validator/validator.go)

```go
func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewUserServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// 中间件调用验证方式
	s.AddUnaryInterceptors(validate.UnaryServerInterceptor())

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
```