# Awpass-route

> create by afterloe <lm6289511@gmail.com>  
> Apache License 2.0

## 前置数据网关
* 分发请求
* 权限拦截
* 请求日志记录

## Hook 参数
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
