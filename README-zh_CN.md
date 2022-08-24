[English](./README.md) | 简体中文
<h1 align="center">padavan_exporter</h1>

<div align="center">

![Build](https://github.com/Bpazy/padavan_exporter/workflows/Build/badge.svg)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=Bpazy_padavan_exporter&metric=alert_status)](https://sonarcloud.io/dashboard?id=Bpazy_padavan_exporter)
[![Go Report Card](https://goreportcard.com/badge/github.com/Bpazy/padavan_exporter)](https://goreportcard.com/report/github.com/Bpazy/padavan_exporter)
[![Docker Pulls](https://img.shields.io/docker/pulls/bpazy/padavan_exporter)](https://hub.docker.com/r/bpazy/padavan_exporter)

这是适用于老毛子固件系统指标的 Prometheus Exporter。请不要吝啬您的任何意见或建议，你可以在 Issue 中讨论她们，也可以直接提交你的 Pull Request.
</div>

## Collectors
| Name                                   | Description                 |
|----------------------------------------|-----------------------------|
| padavan_load1                          | CPU 1min 平均负载.              | 
| padavan_load5                          | CPU 5min 平均负载.              | 
| padavan_load15                         | CPU 15min 平均负载.             | 
| node_cpu_seconds_total                 | 在每种模式下花费的CPU秒数              |
| node_network_receive_bytes_total       | 网络设备统计 receive_bytes.       |
| node_network_receive_compressed_total  | 网络设备统计 receive_compressed.  |
| node_network_receive_errs_total        | 网络设备统计 receive_errs.        |
| node_network_receive_fifo_total        | 网络设备统计 receive_fifo.        |
| node_network_receive_frame_total       | 网络设备统计 receive_frame.       |
| node_network_receive_multicast_total   | 网络设备统计 receive_multicast.   |
| node_network_receive_packets_total     | 网络设备统计 receive_packets.     |
| node_network_transmit_bytes_total      | 网络设备统计 transmit_bytes.      |
| node_network_transmit_carrier_total    | 网络设备统计 transmit_carrier.    |
| node_network_transmit_colls_total      | 网络设备统计 transmit_colls.      |
| node_network_transmit_compressed_total | 网络设备统计 transmit_compressed. |
| node_network_transmit_drop_total       | 网络设备统计 transmit_drop.       |
| node_network_transmit_errs_total       | 网络设备统计 transmit_errs.       |
| node_network_transmit_fifo_total       | 网络设备统计 transmit_fifo.       |
| node_network_transmit_packets_total    | 网络设备统计 transmit_packets.    |

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
通过这个 URL 添加 Dashboard: https://grafana.com/grafana/dashboards/15978

![asd](https://user-images.githubusercontent.com/9838749/160107401-1afbcb8a-8338-476a-8f0e-ed7d801e37b7.png)


## 已知问题
1. 当开启硬件 NAT 时，包转发不经过 CPU，网速统计数据不准确。若关闭硬件 NAT 则会导致性能下降[（参考论坛）](https://www.right.com.cn/forum/thread-4043290-1-1.html) ；
2. Padavan 的文件系统是不可更改的（ tmpfs ），所以本程序目前运行在 ssh 方式。即需要在其他机器运行本程序，并通过 ssh 连接到 Padavan；

## 支持计划
- [ ] 当前连接设备数
