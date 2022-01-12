


```bash
# 替换包名
singo => giligili 


# 解决依赖
go mod tidy

# 创建数据库
create database giligili;
更改字符集编码 utf8 utf8_general_ci


# 解决数据库连接问题
Q：连接时，密码加密
A：设为原生字符串



use mysql;
select user,plugin from user;
UPDATE mysql.user SET authentication_string=PASSWORD('密码'), PLUGIN='mysql_native_password' WHERE USER='root';

service mysql restart

```


```bash
GiliGili 视频
    后端
        CRUD 组件
            Gin、GORM、Go-Redis、其它中间件和实用工具

        系列接口
            接口层面
                投稿视频
                视频详情
                视频列表
                更新视频
                删除视频
                
            模块划分层面
                模型
                控制器
                服务
                视图
                    序列化器

    前端
        vue、Node.js

        视频投稿
        视频详情
            视频播放
            视频直播

        视频列表



从配置文件读配置
    mysql、redis 连接

装载路由
    中间件
    注册路由 => api
    
运行



[GIN-debug] POST   /api/v1/videos   --> giligili/api.CreateVideo (6 handlers)



// c.Json 函数签名 => func (c *Context) Json(code int, obj interface{}) 
```
