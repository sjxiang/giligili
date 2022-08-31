## giligili



## 目的

本项目采用了一系列Golang中比较流行的组件，可以以本项目为基础快速搭建Restful Web API


## 特色

本项目已经整合了许多开发API所必要的组件：

1. [Gin](https://github.com/gin-gonic/gin): 轻量级Web框架，自称路由速度是golang最快的 
2. [GORM](https://gorm.io/index.html): ORM工具。本项目需要配合Mysql使用 
3. [Gin-Session](https://github.com/gin-contrib/sessions): Gin框架提供的Session操作工具
4. [Go-Redis](https://github.com/go-redis/redis): Golang Redis客户端
5. [godotenv](https://github.com/joho/godotenv): 开发环境下的环境变量工具，方便使用环境变量
6. [Gin-Cors](https://github.com/gin-contrib/cors): Gin框架提供的跨域中间件
7. 自行实现了国际化i18n的一些基本功能
8. 本项目是使用基于cookie实现的session来保存登录状态的，如果需要可以自行修改为token验证 后续改成 JWT


本项目已经预先实现了一些常用的代码方便参考和复用:

1. 创建了用户模型
2. 实现了```/api/v1/user/register```用户注册接口
3. 实现了```/api/v1/user/login```用户登录接口
4. 实现了```/api/v1/user/me```用户资料接口(需要登录后获取session)
5. 实现了```/api/v1/user/logout```用户登出接口(需要登录后获取session)


本项目已经预先创建了一系列文件夹划分出下列模块:

1. api文件夹就是MVC框架的controller，负责协调各部件完成任务
2. model文件夹负责存储数据库模型和数据库操作相关的代码
3. service负责处理比较复杂的业务，把业务代码模型化可以有效提高业务代码的质量（比如用户注册，充值，下单等）
4. serializer储存通用的json模型，把model得到的数据库模型转换成api需要的json对象
5. cache负责redis缓存相关的代码
6. auth权限控制文件夹
7. util一些通用的小工具
8. conf放一些静态存放的配置文件，其中locales内放置翻译相关的配置文件

## Godotenv

项目在启动的时候依赖以下环境变量，但是在也可以在项目根目录创建.env文件设置环境变量便于使用(建议开发环境使用)

```shell
MYSQL_DSN="db_user:db_password@/db_name?charset=utf8&parseTime=True&loc=Local" # Mysql连接地址
REDIS_ADDR="127.0.0.1:6379" # Redis端口和地址
REDIS_PW="" # Redis连接密码
REDIS_DB="" # Redis库从0到10
SESSION_SECRET="setOnProducation" # Seesion密钥，必须设置而且不要泄露
GIN_MODE="debug"
```




## 准备工作

```sql

CREATE DATABASE IF NOT EXISTS `giligili` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;

USE `giligili`;

DROP TABLE IF EXISTS `users`;

show create table users\G

describe users;
```

## 运行

```shell
go mod tidy
make run
```

项目运行后启动在3000端口（可以修改，参考gin文档)


## 接口 API 
```
投稿视频
视频详情
视频列表
更新视频
删除视频
```

# 阿里云 oss 对象存储


/api/v1/upload 向阿里云申请 1 个 Token




# 模块划分
```
模型
控制器 一组
服务 数据校验 数据处理
视图 序列化器 数据组装，返回统一的 response
```

阿里云 OSS 对象存储 

    后端只负责把签名返还给 vue，让前端处理（image/jpeg）上传（putURL）下载（getURL）。

    Postman 实现类似效果
        即，
        
        Method [PUT] 
        path [http://endpoint/upload%2Favatar%2F7dbdf8d7-d7a3-4fa9-953e-a49b49fa476b.jpg?Expires=1661797547\u0026OSSAccessKeyId=LTAI5t6yrAZBwLTDxdXqmrfM\u0026Signature=wmeZSKPeB2Va%2B6XsZGqSGwEB3T4%3D]

        header [content-type: image/jpeg]
        body [binary: ]


redis 

    缓存
        1. 普通缓存 点击数
        2. 有序数组 排行榜

    分布式锁

定时任务

跨域
安全
翻页


