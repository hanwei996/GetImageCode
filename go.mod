module GetImageCode

go 1.15

require (
	gitee.com/hanwei66/IHomeWeb v0.0.0-20210425015259-9289cbb53fb2
	github.com/afocus/captcha v0.0.0-20191010092841-4bd1f21c8868
	github.com/astaxie/beego v1.12.3 // indirect
	github.com/garyburd/redigo v1.6.2 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/micro/micro/v3 v3.2.1
	google.golang.org/protobuf v1.25.0

)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

// replace IHome/IHomeWeb => ../IHomeWeb
