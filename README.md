# padavan_exporter
![Build](https://github.com/Bpazy/padavan_exporter/workflows/Build/badge.svg)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=Bpazy_padavan_exporter&metric=alert_status)](https://sonarcloud.io/dashboard?id=Bpazy_padavan_exporter)
[![Docker Pulls](https://img.shields.io/docker/pulls/bpazy/padavan_exporter)](https://hub.docker.com/r/bpazy/padavan_exporter)

## Collectors
Name     | Description
---------|-------------
padavan_load1 | CPU 1min load average from '/system_status_data.asp'. 
padavan_load5 | CPU 5min load average from '/system_status_data.asp'. 
padavan_load15 | CPU 15min load average from '/system_status_data.asp'. 

## 使用方法
```shell
$ ./padavan_exporter --help
Flags:
  --help                        Show context-sensitive help (also try
                                --help-long and --help-man).
  --web.listen-address=":9100"  Address on which to expose metrics and web
                                interface
  --padavan.host="http://192.168.31.1"
                                Padavan address
  --padavan.username="admin"    Padavan username
  --padavan.password="admin"    Padavan password
```
### Grafana 预览
![1](https://user-images.githubusercontent.com/9838749/89121355-c6c10700-d4f0-11ea-92db-499de60bc027.png)

## 已知问题
1. padavan_exporter 运行时会占用登录用户，此时其他设备无法访问 Padavan 控制台。

## 支持计划
- [ ] 网络流量
  1. 硬件 NAT 转发不经过 CPU，那 Padavan 接口返回的数据则不准确。关闭硬件 NAT 则性能下降[（参考论坛）](https://www.right.com.cn/forum/thread-4043290-1-1.html)；
- [ ] 当前连接设备数
