#!/bin/bash
cd /opt/rainbond-cert-controller
cp cfg.example.json cfg.json
env2file conversion -f cfg.json
echo "first run"
flock -xn /tmp/certs-sign.lock -c "cd /opt/rainbond-cert-controller && ./bin/cert-controller >> /dev/stdout 2>&1"
echo "start cron"
exec cron -f -L /dev/stdout