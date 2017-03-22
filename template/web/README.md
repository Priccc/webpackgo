# web端

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
