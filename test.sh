

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




# 上传 token

curl --request POST 'http://localhost:3000/api/v1/upload/token' \
--header 'Content-Type: application/json' \
--data-raw '{"filename": "n.jpg"}'
