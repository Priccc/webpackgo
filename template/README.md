### 服务端

根据自己的需求更改目录及文件

如果你的设备上8800端口已经被占用，请到 `main.go` 修改 port 值

或将配置信息文件 config 移步到 `conf/config.json` 中

启动服务：

``` bash
$ go run main.go
```
如果遇到package找不到的error，请根据提示 `go get package-path`
