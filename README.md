# padavan_exporter
![Build](https://github.com/Bpazy/padavan_exporter/workflows/Build/badge.svg)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=Bpazy_padavan_exporter&metric=alert_status)](https://sonarcloud.io/dashboard?id=Bpazy_padavan_exporter)
[![Docker Pulls](https://img.shields.io/docker/pulls/bpazy/padavan_exporter)](https://hub.docker.com/r/bpazy/padavan_exporter)

## Collectors
Name     | Description
---------|-------------
padavan_load1 | CPU 1min load average. 
padavan_load5 | CPU 5min load average. 
padavan_load15 | CPU 15min load average. 

## 使用方法
```shell
$ ./padavan_exporter --help
Flags:
  --help                        Show context-sensitive help (also try
                                --help-long and --help-man).
  --web.listen-address=":9100"  Address on which to expose metrics and web
                                interface
  --padavan.ssh.host="192.168.31.1:22"
                                Padavan ssh host
  --padavan.ssh.username="admin"
                                Padavan ssh username
  --padavan.ssh.password="admin"
                                Padavan ssh password
  --debug                       Debug mode
```
### systemd
如果你想要通过 systemd 来管理 padavan_exporter，请参考 [这篇文章](https://blog.csdn.net/hanziyuan08/article/details/107749078) 。
### Docker Compose（推荐）
当然更好的方式是使用 Docker Compose，你可以参考本项目预置的 [docker-compose.yml](./docker-compose.yml) 文件。

### Grafana 预览
![1](https://user-images.githubusercontent.com/9838749/89121355-c6c10700-d4f0-11ea-92db-499de60bc027.png)

## 已知问题
1. padavan_exporter 运行时会占用登录用户，此时其他设备无法访问 Padavan 控制台。

## 支持计划
- [ ] 网络流量
  1. 硬件 NAT 转发不经过 CPU，那 Padavan 接口返回的数据则不准确。关闭硬件 NAT 则性能下降[（参考论坛）](https://www.right.com.cn/forum/thread-4043290-1-1.html)；
- [ ] 当前连接设备数
