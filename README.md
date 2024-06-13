# drawio-export-tool
drawioファイルの全ページをpngエクスポートするCLIアプリ

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

## example

```txt
$ ./build/drawio-export-linux-amd64 -f testdata/valid.drawio -o testdata/
drawio page size:  3
exported drawio-export-tool -> testdata/valid-1.png
exported 花 -> testdata/valid-2.png
exported empty -> testdata/valid-3.png
```

## build
※makeコマンドがインストールされていること

ビルドするとbuildディレクトリ配下にバイナリが生成される

ビルド方法
```
make clean && make all
```

