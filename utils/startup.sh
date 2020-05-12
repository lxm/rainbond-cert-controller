#!/bin/bash
cd /opt/rainbond-certs-controller
cp cfg.example.json cfg.json
env2file conversion -f cfg.json
echo "first run"
flock -xn /tmp/certs-sign.lock -c "cd /opt/rainbond-certs-controller && ./bin/certs-controller >> /dev/stdout 2>&1"
echo "start cron"
exec cron -f -L /dev/stdout