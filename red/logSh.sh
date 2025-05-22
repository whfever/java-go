GRANT ALL PRIVILEGES ON xxl_job.* TO 'aes'@'%';
FLUSH PRIVILEGES;

-- 创建用户（如果不存在）
CREATE USER 'aes'@'%' IDENTIFIED BY 'your_password';

-- 授予对 xxl_job 数据库的所有权限
GRANT ALL PRIVILEGES ON xxl_job.* TO 'aes'@'%';

-- 刷新权限
FLUSH PRIVILEGES;


export {http,https}_proxy="http://127.0.0.1:7890"

brew install pyenv

pyenv install 3.12.0

pyenv global 3.12.0
pyenv virtualenv <version> <name>。
pyenv activate <name>。
pyenv deactivate

brew  install goenv


go env -w GOPROXY=https://goproxy.cn,direct
go install github.com/go-delve/delve/cmd/dlv@latest

go mod init ginHello
go mod tidy

go run ginHello.go

brew install wrk

wrk -t12 -c400 -d30s http://127.0.0.1:8080/users

pip install fastapi uvicorn

uvicorn main:app --reload