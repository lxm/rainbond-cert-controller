#!/bin/bash
cd /opt/rainbond-cert-controller
cp cfg.example.json cfg.json
env2file conversion -f cfg.json
exec ./bin/cert-checker