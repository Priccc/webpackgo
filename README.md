# vue-webpack-go

因为项目为go项目，请严格按照go项目目录结构构建项目，并在此目录下执行一下命令！！！

``` bash
$ npm install -g vue-cli
$ vue init Priccc/webpackgo my-project
$ cd my-project
```

### 服务端

根据自己的需求更改目录及文件

如果你的设备上8800端口已经被占用，请到 `conf/config.json` 修改 port 值

(初始化没有这个文件，需要自己新建一个，配置信息请参照 `conf/config_default.json` )

启动服务：

``` bash
$ go run main.go
```
如果遇到package找不到的error，请根据提示 `go get package-path`

### web端

``` bash
$ cd web
$ npm install
$ npm run dev
```

如果你的设备上8801端口已经被占用，请到 `web/config/index.js` 修改 port 值，然后 `npm run dev`

## 怎么使用

- `npm run dev`: 启动前端服务 并实时更新到浏览器上

- `npm run build`: 编译文件到／dist目录下

- `npm run copy` : 复制／dist目录下文件到服务端／assets目录下

- `npm run bdc` : 编译并复制

- `npm run unit`: Unit tests run in PhantomJS with [Karma](http://karma-runner.github.io/0.13/index.html) + [Mocha](http://mochajs.org/) + [karma-webpack](https://github.com/webpack/karma-webpack).

- `npm run e2e`: End-to-end tests with [Nightwatch](http://nightwatchjs.org/).
