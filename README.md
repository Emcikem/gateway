## 目的
实现一个网关的管理：
1. 作为一个代理中间件，需要实现限流，熔断，权限认证(IP白名单，JWT权限认证)，流量统计
2. 负载均衡上，需要进行权值轮询、哈希一致性轮询
3. 网络代理上：基于Go的ReverseProxy实现http的代理。TCP，grpc的代理。利用TCP去实现Thrift的代理

## 前端地址
https://github.com/Emcikem/gateway-vue


## 技术栈
后端：GO + Redis + MySQL + Gin + gRPC
前端：vue + element-ui 

慕课网Go的项目，学习使用


## 学习笔记
### 中间件
中间件一般都是封装在路由上，路由是URL请求分发的管理器

中间件的实现方式
- 基于链表构建中间件，但实现复杂，调用方式不灵活
- 使用数据构建中间件，控制灵活方便

### ReverseProxy
对于简单的http，无法实现一些功能
1. 错误回调及错误日志处理
2. 更改代理返回内容
3. 负载均衡
4. url重写
5. 限流、熔断、降级
6. 数据统计
7. 权限验证

ReverseProxy是Go里官方的一个代理工具

我希望扩展这个工具实现：
- 4种负载轮询类型实现及接口封装
- 扩展中间件支持：限流、熔断实现、权限、数据统计

### connection请求头
是发起方与第一代里的状态
- keep-alive: 不关闭网络
- close： 关闭网络
- Upgrade：协议升级

### X-forwarded-For
记录最后直连实际服务器之前，整个代理过程，可能被篡改

### X-Real-IP
实际请求的ip,不会被改变的

### 负载均衡
1. 随机负载：随机挑选目标服务器IP
2. 轮询负载：ABC三台服务器依次轮询
3. 加权负载：给目标设置访问权重，按照权值轮询
4. 一致性hash负载：请求固定URL访问指定IP

### Websocket
有两个协议：
1. 连接建立协议
2. 数据传输协议

和http的区别就是，http是一问一答的形式
如果说服务器发生改变，需要让客户端及时更新，那么就要去轮询，但这样耗内存，
所以需要websocket，进行连接之后，如果服务器有数据更新，就发送给客户端进行数据更新

### 服务发现
使用tcp进行心跳检测，查询服务的状态，观察者模式通知对应的负载均衡进行服务发现

### 注意点
要求保证：
1. 负载均衡策略池是单例
2. 限流是单例

### 限流
1. 计数器
2. 漏桶
桶的容量是一定的，桶出去的速度是固定的，桶进入的流量是未知的。
保证了请求运行的速度

3. 令牌桶
产生令牌的速度是一定的，每次请求都需要申请令牌，如果令牌不足，那么久拒绝请求，令牌桶也有数量大小限制

