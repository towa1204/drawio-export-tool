# drawio-export-tool
drawioファイルの全ページをpngエクスポートするCLIアプリ

Linux, Windows対応

※[デスクトップ版drawio](https://github.com/jgraph/drawio-desktop)をインストールしておく必要がある

## usage
```
Usage of drawio-export:
  -o string
        specify output directory (default ".")
```

## example

```txt
$ ./build/drawio-export-linux-amd64 -o testdata/ testdata/valid.drawio  testdata/valid2
```

## build
※makeコマンドがインストールされていること

ビルドするとbuildディレクトリ配下にバイナリが生成される

ビルド方法
```
make clean && make all
```

## test
テスト実施
```
make test
```

テストで生成されたファイルの削除
```
make test-clean
```
