FROM tnozicka/s2i-centos7-golang

USER 0

ENV RUN_MODE DEV
ENV APP_NAME golang-wxwork-sdk
ENV ORACLE_HOME /usr/lib/oracle/11.2/client64
ENV PATH $ORACLE_HOME/bin:$PATH
ENV TNS_ADMIN $ORACLE_HOME/network/admin
ENV ORACLE_SID ORCL
ENV LD_LIBRARY_PATH $ORACLE_HOME/lib:$LD_LIBRARY_PATH


# RUN yum install -y gcc
RUN yum install -y libaio-devel
RUN yum install -y expect
RUN yum install -y ImageMagick

ADD ./oracle/11.2/client64 /usr/lib/oracle/11.2/client64
ADD ./oracle/oci8.pc /usr/lib64/pkgconfig/oci8.pc

# ADD ./conf /opt/conf
# ADD ./$APP_NAME /opt/$APP_NAME
# RUN chmod +x /opt/$APP_NAME

ADD ./run.sh /opt/run.sh
RUN chmod +x /opt/run.sh

WORKDIR /opt/

ENTRYPOINT "./$APP_NAME" "${RUN_MODE}"



# docker build -t 192.168.12.35:5000/你的镜像名 .
# docker push 192.168.12.35:5000/你的镜像名:latest
# docker pull 192.168.12.35:5000/你的镜像名:latest

# docker run -it --rm --env "RUN_MODE=DEV APP_NAME=golang-wxwork-sdk" -p 7998:7998 192.168.12.35:5000/你的镜像名:latest /bin/bash
# docker run -it -d --rm --env "RUN_MODE=DEV APP_NAME=golang-wxwork-sdk" -p 7998:7998 192.168.12.35:5000/你的镜像名:latest

# docker images|grep none|awk '{print $3 }'|xargs docker rmi --force

# 测试团队端口6998
# docker stop "golang-wxwork-sdk_6998"
# docker run -it -d --rm --env RUN_MODE=TEST --name="golang-wxwork-sdk_6998" -p 6998:6998 192.168.12.35:5000/你的镜像名:latest

# 开发团队端口7998
# docker stop "golang-wxwork-sdk_7998"
# docker run -it -d --rm --env RUN_MODE=DEV --name="golang-wxwork-sdk_7998" -p 7998:7998 192.168.12.35:5000/你的镜像名:latest

# 正式端口8998
# docker stop "golang-wxwork-sdk_8998"
# docker run -it -d --rm --env RUN_MODE=PRD --name="golang-wxwork-sdk_8998" -p 8998:8998 192.168.12.35:5000/你的镜像名:latest
