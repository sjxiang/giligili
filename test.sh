

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

curl --request POST 'http://localhost:3000/api/v1/upload' \
-F "file=@/home/xsj/go/src/github.com/sjxiang/giligili/static/img/jisoo_2.jpg" \
-H 'Content-Type: multipart/form-data' 




