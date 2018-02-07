FROM alpine

# Update curl
RUN apk --update add curl

# Download analytics beat
RUN ( curl -L https://github.com/sameer1703/efsbeat/releases/download/v1.0.0/efsbeat-5.2.1-SNAPSHOT-linux-x86_64.tar.gz -o /tmp/efsbeat.tar.gz  && cd /tmp/ && tar -xvf efsbeat.tar.gz  && mv /tmp/efsbeat-5.2.1-SNAPSHOT-linux-x86_64 /usr/bin/efsbeat && mkdir /etc/efsbeat && mv /usr/bin/efsbeat/efsbeat.yml /etc/efsbeat/efsbeat.yml && rm -rf /tmp/* )
COPY wsuiteCA.crt /etc/efsbeat/ssl/wsuiteCA.crt
COPY efsbeat.crt /etc/efsbeat/ssl/efsbeat.crt
COPY efsbeat.key /etc/efsbeat/ssl/efsbeat.key

CMD ["/usr/bin/efsbeat/efsbeat", "-c", "/etc/efsbeat/efsbeat.yml", "-e"]


