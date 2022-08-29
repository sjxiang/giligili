

# 视频投稿

curl --request POST 'http://localhost:3000/api/v1/videos' \
--header 'Content-Type: application/json' \
--data-raw '{"title": "第一个视频投稿", "info": "音乐的狂，世代的王"}'


# 视频详情
curl 'http://localhost:3000/api/v1/video/2'



# 视频详情更新

curl --request PUT 'http://localhost:3000/api/v1/video/2' \
--header 'Content-Type: application/json' \
--data-raw '{"title": "第二个视频投稿", "info": "华语天后"}'



# 视频删除

curl --request DELETE 'http://localhost:3000/api/v1/video/2' 




# 上传 token（例如 jisoo.jpg 拿到对象存储地址，这样可以上传或者下载）

curl --request POST 'http://localhost:3000/api/v1/upload' 


{"code":0,
"data":{
    "get":"http://giligili-xsj.oss-cn-hongkong.aliyuncs.com/upload%2Favatar%2F7dbdf8d7-d7a3-4fa9-953e-a49b49fa476b.jpg?Expires=1661797547\u0026OSSAccessKeyId=LTAI5t6yrAZBwLTDxdXqmrfM\u0026Signature=KEH461xvtDwZxdGjsWAZOPuS7eY%3D",
    "key":"upload/avatar/7dbdf8d7-d7a3-4fa9-953e-a49b49fa476b.jpg",
    "put":"http://giligili-xsj.oss-cn-hongkong.aliyuncs.com/upload%2Favatar%2F7dbdf8d7-d7a3-4fa9-953e-a49b49fa476b.jpg?Expires=1661797547\u0026OSSAccessKeyId=LTAI5t6yrAZBwLTDxdXqmrfM\u0026Signature=wmeZSKPeB2Va%2B6XsZGqSGwEB3T4%3D"},
"msg":""}x


