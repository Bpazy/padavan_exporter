# padavan_exporter
![Build](https://github.com/Bpazy/padavan_exporter/workflows/Build/badge.svg)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=Bpazy_padavan_exporter&metric=alert_status)](https://sonarcloud.io/dashboard?id=Bpazy_padavan_exporter)
[![Docker Pulls](https://img.shields.io/docker/pulls/bpazy/padavan_exporter)](https://hub.docker.com/r/bpazy/padavan_exporter)

Prometheus Exporter for padavan metrics. Please do not spare any of your comments or suggestions. You can Discuss them in the Issue, or you can submit a Pull Request directly.

这是适用于老毛子固件系统指标的 Prometheus Exporter。请不要吝啬您的任何意见或建议，你可以在 Issue 中讨论她们，也可以直接提交你的 Pull Request.

## Collectors
Name     | Description
---------|-------------
padavan_load1 | CPU 1min load average. 
padavan_load5 | CPU 5min load average. 
padavan_load15 | CPU 15min load average. 
node_cpu_seconds_total | Seconds the cpus spent in each mode.
node_network_receive_bytes_total | Network device statistic receive_bytes.
node_network_receive_compressed_total | Network device statistic receive_compressed.
node_network_receive_errs_total | Network device statistic receive_errs.
node_network_receive_fifo_total | Network device statistic receive_fifo.
node_network_receive_frame_total | Network device statistic receive_frame.
node_network_receive_multicast_total | Network device statistic receive_multicast.
node_network_receive_packets_total | Network device statistic receive_packets.
node_network_transmit_bytes_total | Network device statistic transmit_bytes.
node_network_transmit_carrier_total | Network device statistic transmit_carrier.
node_network_transmit_colls_total | Network device statistic transmit_colls.
node_network_transmit_compressed_total | Network device statistic transmit_compressed.
node_network_transmit_drop_total | Network device statistic transmit_drop.
node_network_transmit_errs_total | Network device statistic transmit_errs.
node_network_transmit_fifo_total | Network device statistic transmit_fifo.
node_network_transmit_packets_total | Network device statistic transmit_packets.

## 使用方法
```shell
$ ./padavan_exporter --help
Flags:
  --help                        Show context-sensitive help (also try
                                --help-long and --help-man).
  --web.listen-address=":9100"  Address on which to expose metrics and web
                                interface
  --padavan.ssh.host="127.0.0.1:22"
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
Add dashboard by url: https://grafana.com/grafana/dashboards/15978

![1](https://user-images.githubusercontent.com/9838749/90956338-3e62d000-e4b8-11ea-9859-185eb09f6820.jpg)

## 已知问题
1. 当开启硬件 NAT 时，包转发不经过 CPU，网速统计数据不准确。若关闭硬件 NAT 则会导致性能下降[（参考论坛）](https://www.right.com.cn/forum/thread-4043290-1-1.html) ；
2. Padavan 的文件系统是不可更改的（ tmpfs ），所以本程序目前运行在 ssh 方式。即需要在其他机器运行本程序，并通过 ssh 连接到 Padavan； 

## 支持计划
- [ ] 当前连接设备数
