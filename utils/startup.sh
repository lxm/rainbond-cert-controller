#!/bin/bash
cd /opt/rainbond-certs-controller
cp cfg.example.json cfg.json
env2file conversion -f cfg.json
flock -xn /tmp/certs-sign.lock -c "cd /opt/rainbond-certs-controller && ./bin/certs-controller >> /proc/1/fd/1 2>&1"
cron -f