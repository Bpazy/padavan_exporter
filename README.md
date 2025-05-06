English | [简体中文](./README-zh_CN.md)
<h1 align="center">padavan_exporter</h1>

<div align="center">

![Build](https://github.com/Bpazy/padavan_exporter/workflows/Build/badge.svg)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=Bpazy_padavan_exporter&metric=alert_status)](https://sonarcloud.io/dashboard?id=Bpazy_padavan_exporter)
[![Go Report Card](https://goreportcard.com/badge/github.com/Bpazy/padavan_exporter)](https://goreportcard.com/report/github.com/Bpazy/padavan_exporter)
[![Docker Pulls](https://img.shields.io/docker/pulls/bpazy/padavan_exporter)](https://hub.docker.com/r/bpazy/padavan_exporter)

Prometheus Exporter for padavan metrics. Please do not spare any of your comments or suggestions. You can Discuss them in the Issue, or you can submit a Pull Request directly.
</div>

## Collectors

| Name                                   | Description                                   |
|----------------------------------------|-----------------------------------------------|
| padavan_load1                          | CPU 1min load average.                        | 
| padavan_load5                          | CPU 5min load average.                        | 
| padavan_load15                         | CPU 15min load average.                       | 
| node_cpu_seconds_total                 | Seconds the cpus spent in each mode.          |
| node_network_receive_bytes_total     | Total number of bytes received by the network interface.                 |
| node_network_receive_compressed_total| Total number of compressed packets received by the network interface.    |
| node_network_receive_errs_total      | Total number of receive errors detected by the network interface.        |
| node_network_receive_fifo_total      | Total number of FIFO buffer errors on receive by the network interface.  |
| node_network_receive_frame_total     | Total number of frame alignment errors received by the network interface.|
| node_network_receive_multicast_total | Total number of multicast packets received by the network interface.     |
| node_network_receive_packets_total   | Total number of packets received by the network interface.               |
| node_network_transmit_bytes_total    | Total number of bytes transmitted by the network interface.              |
| node_network_transmit_carrier_total  | Total number of carrier errors while transmitting by the network interface. |
| node_network_transmit_colls_total    | Total number of collisions detected while transmitting by the network interface. |
| node_network_transmit_compressed_total | Total number of compressed packets transmitted by the network interface. |
| node_network_transmit_drop_total     | Total number of packets dropped while transmitting by the network interface. |
| node_network_transmit_errs_total     | Total number of transmit errors detected by the network interface.       |
| node_network_transmit_fifo_total     | Total number of FIFO buffer errors on transmit by the network interface. |
| node_network_transmit_packets_total  | Total number of packets transmitted by the network interface.            |
| node_memory_buffers_bytes     | Memory information field Buffers_bytes.    |
| node_memory_cached_bytes     | Memory information field Cached_bytes.    |
| node_memory_free_bytes     | Free memory in bytes.    |
| node_memory_total_bytes     | Total memory in bytes.    |
| node_netstat_Icmp6_InErrors     | Number of ICMPv6 messages failed to receive due to errors.                  |
| node_netstat_Icmp6_InMsgs       | Number of ICMPv6 messages successfully received.                            |
| node_netstat_Icmp6_OutMsgs      | Number of ICMPv6 messages successfully sent.                                |
| node_netstat_Icmp_InErrors      | Number of ICMP messages failed to receive due to errors.                    |
| node_netstat_Icmp_InMsgs        | Number of ICMP messages successfully received.                              |
| node_netstat_Icmp_OutMsgs       | Number of ICMP messages successfully sent.                                  |
| node_netstat_Ip6_InOctets       | Number of incoming octets/packets on IPv6.                                  |
| node_netstat_Ip6_OutOctets      | Number of outgoing octets/packets on IPv6.                                  |
| node_netstat_IpExt_InOctets     | Number of incoming octets/packets on network interfaces, including IPv4.    |
| node_netstat_IpExt_OutOctets    | Number of outgoing octets/packets on network interfaces, including IPv4.    |
| node_netstat_Ip_Forwarding      | Whether IP forwarding is enabled (1) or not (0).                            |
| node_netstat_TcpExt_ListenDrops | Number of TCP listening queue drops.                                        |
| node_netstat_TcpExt_ListenOverflows | Number of TCP listening queue overflows.                                  |
| node_netstat_TcpExt_SyncookiesFailed | Number of failed TCP connections due to invalid SYN cookies.             |
| node_netstat_TcpExt_SyncookiesRecv | Number of received TCP connections with SYN cookies.                       |
| node_netstat_TcpExt_SyncookiesSent | Number of sent TCP SYN cookies.                                            |
| node_netstat_TcpExt_TCPTimeouts | Number of TCP timeouts.                                                    |
| node_netstat_Tcp_ActiveOpens    | Number of active TCP connections openings.                                  |
| node_netstat_Tcp_CurrEstab      | Number of currently established TCP connections.                            |
| node_netstat_Tcp_InErrs         | Number of incoming TCP segments that contained errors.                      |
| node_netstat_Tcp_InSegs         | Number of received TCP segments.                                            |
| node_netstat_Tcp_OutRsts        | Number of TCP segments sent with the RST flag.                              |
| node_netstat_Tcp_OutSegs        | Number of TCP segments sent out.                                            |
| node_netstat_Tcp_PassiveOpens   | Number of passive TCP connections openings.                                 |
| node_netstat_Tcp_RetransSegs    | Number of retransmitted TCP segments.                                       |
| node_netstat_Udp6_InDatagrams   | Number of received UDP6 datagrams.                                          |
| node_netstat_Udp6_InErrors      | Number of UDP6 datagrams that could not be delivered for reasons other than the lack of an application at the destination port. |
| node_netstat_Udp6_NoPorts       | Number of received UDP6 datagrams for which there was no application at the destination port. |
| node_netstat_Udp6_OutDatagrams  | Number of sent UDP6 datagrams.                                              |
| node_netstat_Udp6_RcvbufErrors  | Number of receive buffer errors in UDP6.                                    |
| node_netstat_Udp6_SndbufErrors  | Number of send buffer errors in UDP6.                                       |
| node_netstat_Udp_InDatagrams    | Number of received UDP datagrams.                                           |
| node_netstat_Udp_InErrors       | Number of UDP datagrams that could not be delivered for reasons other than the lack of an application at the destination port. |
| node_netstat_Udp_NoPorts        | Number of received UDP datagrams for which there was no application at the destination port. |
| node_netstat_Udp_OutDatagrams   | Number of sent UDP datagrams.                                               |
| node_netstat_Udp_RcvbufErrors   | Number of receive buffer errors in UDP.                                     |
| node_netstat_Udp_SndbufErrors   | Number of send buffer errors in UDP.                                        |


## Usage
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
If you want to manager padavan_exporter by systemd, [Refer This](https://blog.csdn.net/hanziyuan08/article/details/107749078).
### Docker Compose（Recommend）
Of course the best practice is Docker Compose. You can refer to the preset [docker-compose.yml](./docker-compose.yml).

### Grafana dashboard
Import dashboard by file: [15978_rev3.json](./15978_rev3.json)

Import dashboard by url: https://grafana.com/grafana/dashboards/15978

![image](https://github.com/Bpazy/padavan_exporter/assets/9838749/892ed532-4430-4167-9ee7-ffddc6ef9dc6)


## Known Issues
1. When the hardware NAT is enabled, the packet forwarding does not pass through the CPU, and the network speed statistics are inaccurate. If the hardware NAT is turned off, the performance will be degraded.[(Reference)](https://www.right.com.cn/forum/thread-4043290-1-1.html) ;  
2. Padavan's file system is immutable (TMPFS), so this program is currently running in SSH mode. That is, you need to run this program on other machines and connect to padavan through SSH;  

## Plans
- [ ] Number of devices currently connected.
