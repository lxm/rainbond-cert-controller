#!/bin/bash
cd /opt/rainbond-certs-controller
cp cfg.example.json cfg.json
env2file conversion -f cfg.json
cron -f