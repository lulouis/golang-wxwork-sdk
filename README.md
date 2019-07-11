# golang-wxwork-sdk

#### 介绍
企业微信，企业方服务API提供，企业自行开发扩展，GO语言服务端+docker自动化构建

#### 软件架构
1、GOLANG

2、Swagger generater

3、OAuth2.0

4、Docker File

5、Jenkins with docker-compose.yml


#### 安装教程
1、step up fisrt.
- input your corp. and database info (at conf/app.conf);
- your redis db for token (at main.go line 74);
- cmd "swag init ",generate your restful style site (at your workdir);
- cmd "go run main.go TEST"
- input "http://localhost:6998" at your chrome brower.

2、your can get token with any client_id
http://localhost:6998/oauth2/token?grant_type=client_credentials&client_id=123

if you want connect to your own database , it's todo at line 52 , utils/TokenManager.go


#### 使用说明

1. your can get token use any client_id
http://localhost:6998/oauth2/token?grant_type=client_credentials&client_id=123

#### 参与贡献
1. something important you can email me : zhuaiman@hotmail.com

