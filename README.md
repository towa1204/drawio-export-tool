# drawio-export-tool
drawioファイルの全ページをエクスポートするCLIアプリ

Linux, Windows対応

※[デスクトップ版drawio](https://github.com/jgraph/drawio-desktop)をインストールしておく必要がある

## usage
```
Usage of drawio-export:
  -f string
        specify drawio filename
  -o string
        specify output directory (default ".")
```
## build
※makeコマンドがインストールされていること

ビルドするとbuildディレクトリ配下にバイナリが生成される

ビルド方法
```
make clean && make all
```

