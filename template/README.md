# 服务端

根据自己的需求更改目录及文件

如果你的设备上8800端口已经被占用，请到 `conf/config.json` 修改 port 值

(初始化没有这个文件，需要自己新建一个，配置信息请参照 `conf/config_default.json` )

启动服务：

``` bash
$ go run main.go
```
如果遇到package找不到的error，请根据提示 `go get package-path`

build_release.py 是将代码编包进 release 目录下，生成可执行文件，发布于正式服务器环境
