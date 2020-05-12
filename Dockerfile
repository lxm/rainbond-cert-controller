FROM golang:1.13
RUN mkdir /opt/rainbond-certs-controller/
ADD . /opt/rainbond-certs-controller/
# goproxy
# ENV GOPROXY=https://goproxy.cn
WORKDIR /opt/rainbond-certs-controller/
RUN make


FROM ubuntu:latest as dist
ENV DEBIAN_FRONTEND=noninteractive
ENV ACME_SRORAGE_PATH=/opt/rainbond-certs-controller/storage
ENV ACME_KEY_TYPE=RSA4096

# aliyun mirror
# RUN sed -i "s/archive.ubuntu.com/mirrors.aliyun.com/g" /etc/apt/sources.list && \
# 	sed -i "s/security.ubuntu.com/mirrors.aliyun.com/g" /etc/apt/sources.list

RUN	apt -y update && \
	apt -y install curl cron tzdata && \
	ln -fs /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
	dpkg-reconfigure --frontend noninteractive tzdata

ADD crontab /tmp/crontab
RUN crontab -u root /tmp/crontab
ADD "https://hongyaa-generic.pkg.coding.net/qingjiao/tools/env2file-linux?version=0.1.3" /usr/local/bin/env2file
RUN chmod 755 /usr/local/bin/env2file

RUN mkdir -p /opt/rainbond-certs-controller/bin
WORKDIR /opt/rainbond-certs-controller/

ADD utils/startup.sh /opt/rainbond-certs-controller/startup.sh
RUN chmod 755 /opt/rainbond-certs-controller/startup.sh
COPY --from=0 /opt/rainbond-certs-controller/dist/certs-controller ./bin/
COPY --from=0 /opt/rainbond-certs-controller/cfg.example.json ./
VOLUME [ "/opt/rainbond-certs-controller/storage" ]
CMD ["bash", "-c", "/opt/rainbond-certs-controller/startup.sh"]