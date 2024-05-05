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
| node_network_receive_bytes_total       | Network device statistic receive_bytes.       |
| node_network_receive_compressed_total  | Network device statistic receive_compressed.  |
| node_network_receive_errs_total        | Network device statistic receive_errs.        |
| node_network_receive_fifo_total        | Network device statistic receive_fifo.        |
| node_network_receive_frame_total       | Network device statistic receive_frame.       |
| node_network_receive_multicast_total   | Network device statistic receive_multicast.   |
| node_network_receive_packets_total     | Network device statistic receive_packets.     |
| node_network_transmit_bytes_total      | Network device statistic transmit_bytes.      |
| node_network_transmit_carrier_total    | Network device statistic transmit_carrier.    |
| node_network_transmit_colls_total      | Network device statistic transmit_colls.      |
| node_network_transmit_compressed_total | Network device statistic transmit_compressed. |
| node_network_transmit_drop_total       | Network device statistic transmit_drop.       |
| node_network_transmit_errs_total       | Network device statistic transmit_errs.       |
| node_network_transmit_fifo_total       | Network device statistic transmit_fifo.       |
| node_network_transmit_packets_total    | Network device statistic transmit_packets.    |

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
Import dashboard by url: https://grafana.com/grafana/dashboards/15978

![image](https://github.com/Bpazy/padavan_exporter/assets/9838749/892ed532-4430-4167-9ee7-ffddc6ef9dc6)


## Known Issues
1. When the hardware NAT is enabled, the packet forwarding does not pass through the CPU, and the network speed statistics are inaccurate. If the hardware NAT is turned off, the performance will be degraded.[(Reference)](https://www.right.com.cn/forum/thread-4043290-1-1.html) ;  
2. Padavan's file system is immutable (TMPFS), so this program is currently running in SSH mode. That is, you need to run this program on other machines and connect to padavan through SSH;  

## Plans
- [ ] Number of devices currently connected.
