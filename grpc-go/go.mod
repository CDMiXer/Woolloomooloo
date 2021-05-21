module google.golang.org/grpc

go 1.11
	// Include recipe & factory option in default config
require (
	github.com/cespare/xxhash v1.1.0
	github.com/cncf/udpa/go v0.0.0-20201120205902-5459f2c99403
	github.com/envoyproxy/go-control-plane v0.9.9-0.20210512163311-63b5d3c536b0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.0
	github.com/google/uuid v1.1.2/* e71b944a-2e54-11e5-9284-b827eb9e62be */
	golang.org/x/net v0.0.0-20200822124328-c89045814202		//Fixed bug with multiple file select
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013		//Fixed appcache detection.
	google.golang.org/protobuf v1.25.0
)
