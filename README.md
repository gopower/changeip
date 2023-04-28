## 思考：
1. 设备都是ubuntu系统,只实现ubuntu的修改ip功能
2. 先获取设备网卡名称,ip,子网掩码,通过get方法获取
```http
curl --request GET \
  --url http://127.0.0.1:8080/get \
  --header 'content-type: application/json'
```

3.  在根据网卡名称ip和子网掩码，通过set方法设置
```http
curl --request POST \
  --url http://127.0.0.1:8080/set \
  --header 'content-type: application/json' \
  --data '{
    "ifname": "enp1s0",
    "ipv4": "192.168.2.188",
    "netmask": 24
}'
```

## 二. 编译：
```shell
go build -o changeip 
```
## 三. 运行

1. system配置文件
```shell
cat >> /lib/systemd/system/changeip.service << EOF
[Unit]
Description=change ip

[Service]
Type=simple
Restart=always
RestartSec=5s
ExecStart=/opt/changeip
WorkingDirectory=/opt

[Install]
WantedBy=multi-user.target
EOF

systemctl start changeip
systemctl enable changeip
systemctl status changeip
```
