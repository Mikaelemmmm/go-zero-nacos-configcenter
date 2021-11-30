/**
	author : mikael
	email : 13247629622@163.com
*/
package configcenter

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/conf"
	"time"
)


var bootstrapFile = flag.String("f", "etc/bootstrap.yaml", "the config file")

func main() {

	//解析bootstrap config
	flag.Parse()
	var bootstrapConfig BootstrapConfig
	conf.MustLoad(*bootstrapFile, &bootstrapConfig)

	//解析业务配置
	var c Config
	nacos := NewNacos(bootstrapConfig.NacosConfig)
	serviceConfigContent := nacos.InitConfig(func(data string) {
		err:= conf.LoadConfigFromYamlBytes([]byte(data), &c)
		if err != nil{
			panic(err)
		}
	})
	err:= conf.LoadConfigFromYamlBytes([]byte(serviceConfigContent), &c)
	if err != nil{
		panic(err)
	}

	for{
		fmt.Printf("c : %+v \n" ,c)
		time.Sleep(3 * time.Second)
	}
}
