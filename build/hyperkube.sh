#!/bin/bash
# Copyright 2015 Google Inc. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Starts local hypercube cluster in a Docker container.
# Learn more at https://github.com/kubernetes/kubernetes/blob/master/docs/getting-started-guides/docker.md

# Version of kubernetes to use.
K8S_VERSION="1.1.2"
# Version heapster to use.
# HEAPSTER_VERSION="0.20.0-alpha9"
HEAPSTER_VERSION="1.0.2"
# Port of the apiserver to serve on.
PORT=8080
# Port of the heapster to serve on.
HEAPSTER_PORT=8082

docker run --net=host -d gcr.io/google_containers/etcd:2.0.12 \
   /usr/local/bin/etcd --addr=127.0.0.1:4001 --bind-addr=0.0.0.0:4001 --data-dir=/var/etcd/data


docker run \
    --volume=/:/rootfs:ro \
    --volume=/sys:/sys:ro \
    --volume=/dev:/dev \
    --volume=/var/lib/docker/:/var/lib/docker:ro \
    --volume=/var/lib/kubelet/:/var/lib/kubelet:rw \
    --volume=/var/run:/var/run:rw \
    --net=host \
    --pid=host \
    --privileged=true \
    -d \
    gcr.io/google_containers/hyperkube:v${K8S_VERSION} \
    /hyperkube kubelet --containerized --hostname-override="127.0.0.1" \
        --address="0.0.0.0" --api-servers=http://localhost:${PORT} \
        --config=/etc/kubernetes/manifests

docker run -d --net=host --privileged gcr.io/google_containers/hyperkube:v${K8S_VERSION} \
    /hyperkube proxy --master=http://127.0.0.1:${PORT} --v=2

# Runs Heapster in standalone mode
docker run --net=host -d gcr.io/google_containers/heapster:v${HEAPSTER_VERSION} -port ${HEAPSTER_PORT} \
    --source=kubernetes:http://127.0.0.1:${PORT}?inClusterConfig=false&auth=""