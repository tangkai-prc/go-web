
### 单服伪集群

#### 关闭
`ps -ef | grep "[e]tcd" | awk -F ' ' '{print $2}' | xargs -I {} kill {}`
#### 基于服务器发现

官方的服务发现地址
`curl -w "\n" 'https://discovery.etcd.io/new?size=3'`
```bash
#!/bin/sh

TOKEN=token-01
CLUSTER_STATE=new
NAME_1="etcd-1"
NAME_2="etcd-2"
NAME_3="etcd-3"

HOST_1="192.168.1.151"
HOST_2="192.168.1.151"
HOST_3="192.168.1.151"


CLIENT_PORT_1=2377
CLIENT_PORT_2=2378
CLIENT_PORT_3=2379

PEER_PORT_1=2380
PEER_PORT_2=2381
PEER_PORT_3=2382

DISCOVERY="https://discovery.etcd.io/6a13232d9058d7a9895e7dd5f8759d0b"

THIS_NAME=${NAME_1}
THIS_IP=${HOST_1}
nohup /opt/etcd-v3.3.13-linux-amd64/etcd --data-dir=${THIS_NAME}".etcd" --name ${THIS_NAME} \
    --initial-advertise-peer-urls http://${THIS_IP}:${PEER_PORT_1} --listen-peer-urls http://${THIS_IP}:${PEER_PORT_1} \
    --advertise-client-urls http://${THIS_IP}:${CLIENT_PORT_1} --listen-client-urls http://${THIS_IP}:${CLIENT_PORT_1} \
    --discovery ${DISCOVERY} \
    --initial-cluster-state ${CLUSTER_STATE} --initial-cluster-token ${TOKEN} >> ${THIS_NAME}".nohup" &
echo "启动"${NAME_1}"成功"

THIS_NAME=${NAME_2}
THIS_IP=${HOST_2}
nohup /opt/etcd-v3.3.13-linux-amd64/etcd --data-dir=${THIS_NAME}".etcd" --name ${THIS_NAME} \
    --initial-advertise-peer-urls http://${THIS_IP}:${PEER_PORT_2} --listen-peer-urls http://${THIS_IP}:${PEER_PORT_2} \
    --advertise-client-urls http://${THIS_IP}:${CLIENT_PORT_2} --listen-client-urls http://${THIS_IP}:${CLIENT_PORT_2} \
    --discovery ${DISCOVERY} \
    --initial-cluster-state ${CLUSTER_STATE} --initial-cluster-token ${TOKEN} >> ${THIS_NAME}".nohup" &
echo "启动"${NAME_2}"成功"


THIS_NAME=${NAME_3}
THIS_IP=${HOST_3}
nohup /opt/etcd-v3.3.13-linux-amd64/etcd --data-dir=${THIS_NAME}".etcd" --name ${THIS_NAME} \
    --initial-advertise-peer-urls http://${THIS_IP}:${PEER_PORT_3} --listen-peer-urls http://${THIS_IP}:${PEER_PORT_3} \
    --advertise-client-urls http://${THIS_IP}:${CLIENT_PORT_3} --listen-client-urls http://${THIS_IP}:${CLIENT_PORT_3} \
    --discovery ${DISCOVERY} \
    --initial-cluster-state ${CLUSTER_STATE} --initial-cluster-token ${TOKEN} >> ${THIS_NAME}".nohup" &
echo "启动"${NAME_3}"成功"
```



```bash
#!/bin/sh

NAME_1="etcd-1"
NAME_2="etcd-2"
NAME_3="etcd-3"

HOST_1="192.168.1.151"
HOST_2="192.168.1.151"
HOST_3="192.168.1.151"


CLIENT_PORT_1=2377
CLIENT_PORT_2=2378
CLIENT_PORT_3=2379

PEER_PORT_1=2380
PEER_PORT_2=2381
PEER_PORT_3=2382
 
echo "----------------------------"
THIS_NAME=${NAME_1}
THIS_IP=${HOST_1}
CLIENT_PORT=${CLIENT_PORT_1}
PEER_PORT=${PEER_PORT_1}
nohup /opt/etcd-v3.3.13-linux-amd64/etcd \
    --data-dir=${THIS_NAME}".etcd" \
    --name ${THIS_NAME} \
    --initial-advertise-peer-urls http://${THIS_IP}:${PEER_PORT} \
    --listen-peer-urls http://${THIS_IP}:${PEER_PORT} \
    --advertise-client-urls http://${THIS_IP}:${CLIENT_PORT} \
    --listen-client-urls http://${THIS_IP}:${CLIENT_PORT} \
    --initial-cluster ${NAME_1}=http://${HOST_1}:${PEER_PORT_1},${NAME_2}=http://${HOST_2}:${PEER_PORT_2},${NAME_3}=http://${HOST_3}:${PEER_PORT_3} \
    --initial-cluster-state new \
    --initial-cluster-token docker-etcd &> ${THIS_NAME}".nohup" &
echo "启动"${NAME_1}"成功"

echo "----------------------------"
THIS_NAME=${NAME_2}
THIS_IP=${HOST_2}
CLIENT_PORT=${CLIENT_PORT_2}
PEER_PORT=${PEER_PORT_2}

nohup /opt/etcd-v3.3.13-linux-amd64/etcd \
    --data-dir=${THIS_NAME}".etcd" \
    --name ${THIS_NAME} \
    --initial-advertise-peer-urls http://${THIS_IP}:${PEER_PORT} \
    --listen-peer-urls http://${THIS_IP}:${PEER_PORT} \
    --advertise-client-urls http://${THIS_IP}:${CLIENT_PORT} \
    --listen-client-urls http://${THIS_IP}:${CLIENT_PORT} \
    --initial-cluster ${NAME_1}=http://${HOST_1}:${PEER_PORT_1},${NAME_2}=http://${HOST_2}:${PEER_PORT_2},${NAME_3}=http://${HOST_3}:${PEER_PORT_3} \
    --initial-cluster-state new \
    --initial-cluster-token docker-etcd &> ${THIS_NAME}".nohup" &
echo "启动"${NAME_2}"成功"


echo "----------------------------"
THIS_NAME=${NAME_3}
THIS_IP=${HOST_3}
CLIENT_PORT=${CLIENT_PORT_3}
PEER_PORT=${PEER_PORT_3}
nohup /opt/etcd-v3.3.13-linux-amd64/etcd \
    --data-dir=${THIS_NAME}".etcd" \
    --name ${THIS_NAME} \
    --initial-advertise-peer-urls http://${THIS_IP}:${PEER_PORT} \
    --listen-peer-urls http://${THIS_IP}:${PEER_PORT} \
    --advertise-client-urls http://${THIS_IP}:${CLIENT_PORT} \
    --listen-client-urls http://${THIS_IP}:${CLIENT_PORT} \
    --initial-cluster ${NAME_1}=http://${HOST_1}:${PEER_PORT_1},${NAME_2}=http://${HOST_2}:${PEER_PORT_2},${NAME_3}=http://${HOST_3}:${PEER_PORT_3} \
    --initial-cluster-state new \
    --initial-cluster-token docker-etcd &> ${THIS_NAME}".nohup" &
echo "启动"${NAME_3}"成功"

```