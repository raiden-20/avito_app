FROM debian

RUN apt update -y && apt upgrade -y 
RUN apt install lsb-release curl gpg -y
RUN curl -fsSL https://packages.redis.io/gpg | gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg
RUN echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" | tee /etc/apt/sources.list.d/redis.lis
RUN apt-get update -y
RUN apt-get install redis -y

EXPOSE 6379

COPY  ./cert.pem /usr/local/redis/cert.pem
COPY  ./key.pem /usr/local/redis/key.pem
COPY redis.conf /usr/local/redis/redis.conf

WORKDIR /data
RUN  rm -f ~/dump.rdb /dump.rdb
# RUN sed '/port 6379/d' /etc/redis/redis.conf
# RUN echo 'tls-port 6379' >> /etc/redis/redis.conf
# RUN echo 'tls-client-cert-file /usr/local/redis/redis.crt' >> /etc/redis/redis.conf
# RUN echo 'tls-client-key-file /usr/local/redis/redis.key' >> /etc/redis/redis.conf

CMD ["redis-server", "/usr/local/redis/redis.conf"]