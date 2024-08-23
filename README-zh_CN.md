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
| node_memory_buffers_bytes     | 内存信息 Buffers_bytes.    |
| node_memory_cached_bytes     | 内存信息 field Cached_bytes.    |
| node_memory_free_bytes     | 剩余内存 in bytes.    |
| node_memory_total_bytes     | 总内存 in bytes.    |
| node_netstat_Icmp6_InErrors     | 由于错误未能成功接收的 ICMPv6 消息数量。                                 |
| node_netstat_Icmp6_InMsgs       | 成功接收的 ICMPv6 消息数量。                                             |
| node_netstat_Icmp6_OutMsgs      | 成功发送的 ICMPv6 消息数量。                                             |
| node_netstat_Icmp_InErrors      | 由于错误未能成功接收的 ICMP 消息数量。                                   |
| node_netstat_Icmp_InMsgs        | 成功接收的 ICMP 消息数量。                                               |
| node_netstat_Icmp_OutMsgs       | 成功发送的 ICMP 消息数量。                                               |
| node_netstat_Ip6_InOctets       | IPv6 上的传入字节数/数据包数量。                                         |
| node_netstat_Ip6_OutOctets      | IPv6 上的传出字节数/数据包数量。                                         |
| node_netstat_IpExt_InOctets     | 网络接口上传入的字节数/数据包数量，包含 IPv4。                           |
| node_netstat_IpExt_OutOctets    | 网络接口上传出的字节数/数据包数量，包含 IPv4。                           |
| node_netstat_Ip_Forwarding      | 是否启用了 IP 转发，启用 (1) 或未启用 (0)。                              |
| node_netstat_TcpExt_ListenDrops | TCP 监听队列丢弃的数量。                                                 |
| node_netstat_TcpExt_ListenOverflows | TCP 监听队列溢出的数量。                                             |
| node_netstat_TcpExt_SyncookiesFailed | 由于无效 SYN cookies 而失败的 TCP 连接数量。                        |
| node_netstat_TcpExt_SyncookiesRecv | 接收到的带有 SYN cookies 的 TCP 连接数量。                               |
| node_netstat_TcpExt_SyncookiesSent | 发送的 TCP SYN cookies 数量。                                           |
| node_netstat_TcpExt_TCPTimeouts | TCP 超时的数量。                                                        |
| node_netstat_Tcp_ActiveOpens    | 主动打开的 TCP 连接数量。                                                |
| node_netstat_Tcp_CurrEstab      | 当前已建立的 TCP 连接数量。                                              |
| node_netstat_Tcp_InErrs         | 包含错误的接收到的 TCP 段数量。                                          |
| node_netstat_Tcp_InSegs         | 接收到的 TCP 段数量。                                                    |
| node_netstat_Tcp_OutRsts        | 发送的带有 RST 标志的 TCP 段数量。                                       |
| node_netstat_Tcp_OutSegs        | 发送出的 TCP 段数量。                                                    |
| node_netstat_Tcp_PassiveOpens   | 被动打开的 TCP 连接数量。                                                |
| node_netstat_Tcp_RetransSegs    | 重传的 TCP 段数量。                                                      |
| node_netstat_Udp6_InDatagrams   | 接收到的 UDP6 数据报数量。                                               |
| node_netstat_Udp6_InErrors      | 由于原因不明（非目的端口无应用程序）未能投递的 UDP6 数据报数量。         |
| node_netstat_Udp6_NoPorts       | 接收到的目的端口无应用程序的 UDP6 数据报数量。                           |
| node_netstat_Udp6_OutDatagrams  | 发送的 UDP6 数据报数量。                                                 |
| node_netstat_Udp6_RcvbufErrors  | UDP6 收到的接收缓冲区错误数量。                                          |
| node_netstat_Udp6_SndbufErrors  | UDP6 发送缓冲区错误的数量。                                              |
| node_netstat_Udp_InDatagrams    | 接收到的 UDP 数据报数量。                                                |
| node_netstat_Udp_InErrors       | 由于原因不明（非目的端口无应用程序）未能投递的 UDP 数据报数量。          |
| node_netstat_Udp_NoPorts        | 接收到的目的端口无应用程序的 UDP 数据报数量。                            |
| node_netstat_Udp_OutDatagrams   | 发送的 UDP 数据报数量。                                                  |
| node_netstat_Udp_RcvbufErrors   | UDP 收到的接收缓冲区错误数量。                                           |
| node_netstat_Udp_SndbufErrors   | UDP 发送缓冲区错误的数量。                                               |

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

### Grafana 面板
通过这个 URL 导入 Dashboard: https://grafana.com/grafana/dashboards/15978

![image](https://github.com/Bpazy/padavan_exporter/assets/9838749/d5769689-47a6-41bd-b859-d168fd19ec50)


## 已知问题
1. 当开启硬件 NAT 时，包转发不经过 CPU，网速统计数据不准确。若关闭硬件 NAT 则会导致性能下降[（参考论坛）](https://www.right.com.cn/forum/thread-4043290-1-1.html) ；
2. Padavan 的文件系统是不可更改的（ tmpfs ），所以本程序目前运行在 ssh 方式。即需要在其他机器运行本程序，并通过 ssh 连接到 Padavan；

## 支持计划
- [ ] 当前连接设备数
