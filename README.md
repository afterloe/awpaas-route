# Awpass-route

> create by afterloe <lm6289511@gmail.com>  
> Apache License 2.0

## 前置数据网关
&emsp;前置数据网关的作用是针对数据请求进行权限拦截、服务转发。他是SOA架构的主入口，用户通过访问前置数据网关，传入URL实现从SOA众多微服务之中获取数据，SOA中部署的服务包括空间地图服务、数据关联分析服务、大数据服务等。网关的特点是请求访问量大，吞吐量大，在设计过程中要充分考虑到请求到响应的时间问题。  
&emsp;值得注意的是，前置数据网关是基础支撑平台的数据对接顶层，不对数据进行应用层面的加工处理，也不对数据的权限进行校验，核心是只进行服务转发，和鉴权字段拦截，而权限验证这一工作是由部署在SOA中相应的应用系统完成的。  
功能包括：
* 能够实现SOA中各类服务的请求转发包括HTTP、WEBSOCKET、RPC。
* 支持访问日志输出，以提供给其他服务使用。（如：流量统计，活跃用户统计等）
* 支持白名单访问，即在白名单之中的请求无需鉴权即可实现转发。
* 支持在线实时修改配置，高效灵活的适应环境。
* 支持横向扩展部署，占用资源少，简化部署和配置。
* 采用高效的通知机制，保障实时修改。
* 可以主动和被动的获取信息。
* 支持多种请求的双向数据传输。
* 异常情况自动重启
* 符合整体安全策略，保障网关与其SOA及底层系统的两方面安全。

## Hook 参数
用于docker 启动
```sbtshell
env REDIS_ADDR
env REDIS_PORT
```

## 使用
### 已有镜像进行服务部署
```sbtshell
docker service create \
--replicas 4 \
--network awpaas \
--detach=false \
--name api-gateway \
--env REDIS_ADDR=192.168.2.13 \
--host cluster-1:192.168.2.13 \
--publish 8080:8080 \
--publish 8081:8081 \
awpaas/awpaas-route:version
```
### 服务更新
```sbtshell
# git pull && make -m src=/data/data-2/go/src
# docker tag awpaas/awpaas-route:version 127.0.0.1/awpaas/awpaas-route:version
# docker push 127.0.0.1/awpaas/awpaas-route:version
# docker service update --image awpaas/awpaas-route:version api-gateway
```
