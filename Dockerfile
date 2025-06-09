FROM alpine:latest
WORKDIR /REMIND37
ADD ./remind37 ./
RUN mkdir ./web
ADD ./web.tar ./web
# 1. 安装 tzdata 包
RUN apk add --no-cache tzdata

# 2. 设置时区为上海
ENV TZ=Asia/Shanghai

# 3. （可选）创建软链接，某些应用可能依赖 /etc/localtime
RUN ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime

# 验证时区（测试用）
RUN echo "当前时区: $(date)"
EXPOSE 8080
ENTRYPOINT ["/REMIND37/remind37"]
