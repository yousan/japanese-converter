[![Deploy to Vercel](https://vercel.com/button)](https://vercel.com/import/project?template=https://github.com/yousan/japanese-converter/)

# Japanese Converter
日本語の文字コードを変換するAPIを作成します。

![demo](https://github.com/yousan/japanese-converter/blob/master/.readme/demo.gif?raw=true)


[デプロイ済みのエンドポイント](https://japanese-converter.vercel.app/api)

# テスト

```
$ curl --data-binary '@./testdata/utf8.txt' https://japanese-converter.vercel.app/api
```

# デプロイ
Vercelで自動的にデプロイされています。

[![Deploy to Vercel](https://vercel.com/button)](https://vercel.com/import/project?template=https://github.com/yousan/japanese-converter/)

Vercelのアカウントがある状態でフォークしてコミットプッシュすると自動的にデプロイされます。

# 出典
変換を行うコードは下記のライブラリを利用しています。

https://github.com/spiegel-im-spiegel/text/
