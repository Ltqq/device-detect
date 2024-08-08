# 项目名称和二进制文件名称
APP_NAME = device-status

# Go 源代码目录
SRC_DIR = ./cmd

# 默认目标
all: build

# 执行 go mod tidy
tidy:
	go mod tidy

# 构建二进制文件
build: tidy
	go build -o $(APP_NAME) main.go

# 清理构建生成的文件
clean:
	rm -f $(APP_NAME)

# 运行项目
run: build
	./$(APP_NAME) all

# 打印帮助信息
help:
	@echo "Usage:"
	@echo "  make all       - 执行 tidy 和 build"
	@echo "  make tidy      - 执行 go mod tidy"
	@echo "  make build     - 构建二进制文件"
	@echo "  make clean     - 清理构建生成的文件"
	@echo "  make run       - 构建并运行项目"
	@echo "  make help      - 打印此帮助信息"

.PHONY: all tidy build clean run help
