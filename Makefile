# Makefile

# ターゲット名
TARGET = drawio-export

# ソースコードのエントリーポイント
SOURCE = .

# ビルド時の追加のフラグ（必要に応じて設定）
BUILD_FLAGS =

# 出力ディレクトリ
OUTPUT_DIR = build

# ターゲットとなるOSとアーキテクチャ
WINDOWS_TARGET = $(OUTPUT_DIR)/$(TARGET)-windows-amd64.exe
LINUX_TARGET = $(OUTPUT_DIR)/$(TARGET)-linux-amd64

# デフォルトターゲット
.DEFAULT_GOAL := all

# 全てのターゲットをビルド
all: windows linux

# Windows向けにクロスコンパイル
windows:
	GOOS=windows GOARCH=amd64 go build $(BUILD_FLAGS) -o $(WINDOWS_TARGET) $(SOURCE)

# Linux向けにクロスコンパイル
linux:
	GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -o $(LINUX_TARGET) $(SOURCE)

clean:
	rm -rf $(OUTPUT_DIR)

.PHONY: all windows linux clean
