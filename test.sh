

# 视频投稿

curl --request POST 'http://localhost:3000/api/v1/videos' \
--header 'Content-Type: application/json' \
--data-raw '{"title": "第一个视频投稿", "info": "音乐的狂，世代的王", "url": "和图片一个原理，略过", "avatar": "upload/avatar/65f8d983-0898-49e2-8293-053f06ade07d.jpg"}'


# 视频展示

curl 'http://localhost:3000/api/v1/video/4'


# 视频列表

curl 'http://localhost:3000/api/v1/videos'

# 视频更新

curl --request PUT 'http://localhost:3000/api/v1/video/2' \
--header 'Content-Type: application/json' \
--data-raw '{"title": "第二个视频投稿", "info": "华语天后"}'



# 视频删除

curl --request DELETE 'http://localhost:3000/api/v1/video/2' 




# 上传 token（例如 jisoo.jpg 拿到对象存储地址，这样可以上传或者下载）

curl --request POST 'http://localhost:3000/api/v1/upload/token' 



# 视频排行榜

curl POST 'http://localhost:3000/api/v1/rank/daily'