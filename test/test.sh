curl -i "http://localhost:8888/shorten?url=https://go-zero.dev"

# 结果如下:
#HTTP/1.1 200 OK
#Content-Type: application/json; charset=utf-8
#Traceparent: 00-d247fefb567d618abde4179185212761-26d32063ea819d29-00
#Date: Tue, 13 May 2025 12:16:00 GMT
#Content-Length: 20
#
#{"shorten":"b0434f"}

curl -i "http://localhost:8888/expand?shorten=b0434f"
# 结果如下:
#HTTP/1.1 200 OK
#Content-Type: application/json; charset=utf-8
#Traceparent: 00-a92bdcd455142f21579bafe83765e2d7-cd8dece0d249f29b-00
#Date: Tue, 13 May 2025 12:17:05 GMT
#Content-Length: 29
#
#{"url":"https://go-zero.dev"}