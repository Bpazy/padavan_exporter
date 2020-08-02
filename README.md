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

## 已知问题
1. padavan_exporter 运行时会占用登录用户，此时其他设备无法访问 Padavan 控制台。

## 支持计划
- [x] CPU Load average
- [ ] 网络流量（瞬时流量、总流量）
- [ ] 当前连接设备数
