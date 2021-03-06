# 数字健康实验室学术垃圾回收站

**伪·毕业传承池**  
**真·屎山集散地**

**实在**是受不了Beego了，用Gin搓一个轻量一点的网页（内部用的，不是骗钱项目，不赶工期）

## TODO

- [ ] 注册登陆等
- [ ] 文件服务器
- [ ] 文件上传下载
- [ ] 访问控制

可能会做：

- [ ] markdown在线编辑修改
- [ ] 图片视频在线预览
- [ ] 类似web版的数据库图形界面（dashboard）
- [ ] RPC接口和相关API
- [ ] 断点续传
- [ ] 文件加密
- [ ] 文件校验

## 开发

```bash
git clone https://github.com/xi102/collection
```

### 后端

需要安装go, 在根目录

```bash
go mod tidy
go run main.go
```

或者用[Makefile](./Makefile)

```bash
make run
```

### 前端

需要安装nodejs、vue-cli、yarn/npm

```bash
cd frontend
yarn install && yarn serve
```

## 部署

复制一份配置文件并改名`config.toml`，修改数据库密码端口等配置信息

```bash
cp conf/default.toml config.toml
```

```bash
cd frontend && yarn install
yarn build
```

于是你有了前端静态文件

```bash
make build
#或者直接
go build -o collection main.go
```

直接运行编译好的二进制文件即可

```bash
./collection
```

## 库和框架

后端：

- gin 用来处理路由和网络请求，自带的orm和log等都没用
- jwt-go, 用于登陆鉴权
- gin-contrib/cors，gin的官方middleware,用于配置CORS
- xorm，用于orm
- logrus、rotatelogs, 用于日志处理
- viper, 功能很多的一个用来读取配置文件的库，这里只是简单读取一下toml配置

前端：

- vue
- vuetify

测试：

- testify
- httptest

## 关于

今年又有一批师兄师姐要毕业跑路了。  
老板：做一个管理交接文档和数据的网站呗，这个对你们来说应该很简单吧。有不同的用户角色和密码，能上传文件和下载，最好是那种点进网页就能看到那个项目数据库相关的数据表在哪里，在网页上展示出来。。。
数据安全特别重要的，一定要管好，可不能泄漏，最近正在查。。。  

我：还嫌服务器漏洞不够多。。。直接把数据表都暴露出来这不是作死。。。
老板：就我们内部用用，不开放出来不行吗  
我：102那几个路由器和服务器又不在一个局域网里面，  （iptables白名单就行但我没说）  
老板：。。。  
室友：要不我们买个交换机（超小声）  
我：这样真的不安全，要不就把数据库里的数据导出成文件再写个建库脚本放在那呗。给文件设个密码，这样安全点（被问我为什么不说gpg加密文件。。）  
老板：行吧，到时候把密码发我一份。  
室友：什么时候做完啊，他们毕业之前还是毕业以后啊？  
老板：当然是毕业之前，一个月以内。  
我，室友：行。。。  
