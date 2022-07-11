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