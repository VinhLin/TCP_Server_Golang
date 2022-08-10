# TCP_Server
Simple Test TCP_Server use Golang

### Documents
- [Golang TCP Server](https://linuxhint.com/golang-tcp-server/
- Tham khao khac:
```
https://blog.katastros.com/a?ID=00850-94f7e646-73a1-4f7c-a52a-3d49b27b39a1
https://www.youtube.com/watch?v=s1oh-vl0fqI
```

### Library
- [sms](github.com/warthog618/sms)
> go get github.com/warthog618/sms


### Build and Run
- Build to **tcp_server
> go build -o tcp_server *.go
- Run file:
> ./tcp_server


## Test: File code Bin, Ngrok, Netcat
- In my PC, run code:
> ./tcp_server
- Run Ngrok: 
> ngrok tcp 9001
### Trên thiết bị muốn gửi data mình thực hiện
- Giả sử mình có địa chỉ sau khi chạy ngrok là: `tcp://0.tcp.ap.ngrok.io:11433`
- Thực hiện câu lệnh **netcat**:
> echo "Test ngrok" |nc 0.tcp.ap.ngrok.io 11433

echo 07911614220991F1040B911605935713F200008140806113912304D7F79B0E |nc 0.tcp.ap.ngrok.io 11433
