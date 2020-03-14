# make: 执行go build 生成二进制文件 
# make gotool: 执行 gofmt -w . 和 go tool vet .（格式化代码和源码静态检查）
# make clean: 清除二进制文件
# make ca: 生成证书
# make help: 打印help信息
# .PHONY
all: gotool
	@go build -v .
clean:
	rm -f apiserver
	find . -name "[._]*.s[a-w][a-z]" | xargs -i rm -f {}
gotool:
	gofmt -w .
	go tool vet . |& grep -v vendor;true
ca:
	openssl req -new -nodes -x509 -out conf/server.crt -keyout conf/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1/emailAddress=xxxxx@qq.com"

help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make gotool - run go tool 'fmt' and 'vet'"
	@echo "make ca - generate ca files"

.PHONY: clean gotool ca help