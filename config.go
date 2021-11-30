/**
	author : mikael
	email : 13247629622@163.com
*/
package configcenter

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB struct{
		DataSource string
	}
	Cache cache.CacheConf

	IdentityRpcConf zrpc.RpcClientConf
}

