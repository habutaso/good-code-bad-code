---
marp: true
paginate: true
theme: default
style: |
  pre { font-size: 32px; }
  p, li, ol, ul { font-size: 28px; }
  h3 { font-size: 32px; }
  img[alt~="center"] {
  display: block;
  margin: 0 auto;
  }
  .columns {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: 1rem;
  }
  .wider-right {
    grid-template-columns: 1fr 2fr;
  }
---

<!--
class: lead invert
-->

# 名前設計
\- あるべき構造を見破る名前 \-

---

<!--
class: lead invert
-->

# ビジネス目的に寄せた名前設計

---

<!--
class: lead invert
-->

# 目的駆動名前設計

---

<!--
class: top
-->

# 目次

1. 悪魔を呼び寄せる名前とは
2. 目的ベースで名づけする
3. 名づけのリスク
4. 意味が伝わりにくい名前
---

<!--
class: lead invert
-->

# 悪魔を呼び寄せる名前とは

---

<!--
class: top
-->

# 意味をまとめてしまう
あるECサイトを作ると仮定する
巨大な「商品」クラスがある

---

<!--
class: top
-->

# 意味をまとめてしまう
結びついたクラスに関係するロジックができ始め、複雑化
-> 仕様変更のたびにすべてに影響！
-> 毎回確認しなくてはいけないため、生産性は低下・・・

---

<!--
class: lead
-->

# どうすればよい？

---

<!--
class: lead
-->

# 関心ごとに分ける

---

<!--
class: top
-->

# もう一度同じ図
「商品」クラスは、4つの関心ごとを持っている。これらはユースケースごとに意味が異なる。
設計上は、**関心の分離**を行う

---

<!--
class: top
-->

# 関心事に分けるメリット
1. 疎結合高凝集を果たすことができる
2. 仕様変更時の影響度を少なくできる


---
# 背景にある考え方(参考)
**ドメイン駆動設計**における「境界づけられたコンテキスト」が当てはまる
この本での関心事=コンテキスト

## コンテキストってなに？
正直言葉にするのは難しいですが、システム開発においては「あるサービスを構成し、そのサービスの状態や文化、状況を共有する構成要素のこと」を指すのだと思います。

<footer>
https://little-hands.hatenablog.com/entry/2017/11/28/bouded-context-concept
</footer>

---
# ドメイン駆動
[Eric Evans本](https://www.amazon.co.jp/gp/product/B00GRKD6XU/ref=as_li_tl?ie=UTF8&camp=247&creative=1211&creativeASIN=B00GRKD6XU&linkCode=as2&tag=majackyy-22&linkId=29ec226ada515b83691ac3b54222fa6c)が最初
名前設計に関しては、ドメイン駆動が背景にあるとおもいます。

---

<!--
class: lead invert
-->

# 目的ベースで名づけする 

---

<!--
class: top
-->

# 名前を付けるときの考え方
ECサイトを作るため、設計しているとする。
まず何が存在するだろう。。「住所」「商品」「ユーザー」があって。。。
このような考え方は「存在ベース」と言う。

単純に存在を示すだけの名前は、意味が多重になりがちで目的不明オブジェクトになってしまいやすい
+ 単に「ユーザー」でも、「個人」「法人」といったバリエーションをもっている。
+ なので、名前を「ユーザー」にすると、「個人」ロジックと「法人」ロジックが混じる

---

<style scoped>
  table {
    table-layout: fixed;
    border-collapse: collapse;
    display: table;
    font-size: 24px;
  }
</style>

# 目的ベース
例えば「住所」は「**商品の配送**」をするために使用する -> これが目的
ECサイトを構成する要素(存在)がどのような目的で使用されるのかが名づけのベースとなり、この手法がロジック実装の混乱を避ける

### 例

| 存在ベース | 目的ベース |
| --------- | -------- |
| 住所 | 配送元、配送先、勤務先、本籍地 |
| 商品 | 入庫品、予約品、注文品、配送品 |
| ユーザー | アカウント、個人プロフィール、職務履歴 |
| ユーザー名 | アカウント名、表示名、本名、法人名 |
| 金額 | 請求金額、消費税額、延滞保証料、キャンペーン割引料金 |

---
# どんなビジネス目的があるか分析
登場人物や事柄を列挙して、関係性を整理、分析する
分析手法は本には載っていませんのでいくつか調べてみました。

1. ビジネスモデル・キャンバスとピクト図解を合わせた手法
ビジネスチームとの提供価値、ヒト、カネ、モノ、それら関連の共通認識に使用できそう
https://dhbr.diamond.jp/articles/-/2447
2. ビジネスを表すデータモデル図
ビジネスチーム->開発チームへの構造詳細化に使えそう
https://www.it-innovation.co.jp/2017/01/22-231759/

---

# 会話
開発A「ちょっと認識の確認をしたく。今回のアプリ、フリー版は14日の制限をする。あってますか？」
ビジネスB「そうですねー、有料プランの収益は非常に大きいですー」
開発A「皆さんの認識合ってましたが、図にちゃんと書いてなかったですね。追加します。**個人客と法人客のフリー版の線に14日の制限を追加っと**」
ビジネスB「あ、**法人の制限は60日**にしようと思ってます」
開発A「わお」

新しく、**法人の制限**という単語が出てきました！みたいなパターン

積極的に話し合って、会話の中に特化した名前がないか注意深く観察しましょう。
-> 議事録を取るのは大事です。

---
# ユビキタス言語
チーム全体で意図を共有するための言葉、それを集めた辞書
同じ意図の名前を会話、ドキュメント、コードにまで反映することで、意図と成果物のつながりを保つことができる
-> ユビキタス言語のメンテナンス厳しいと思います
-> プルリクレビューではコードと名前の整合性チェックに結構ストレスがかかりそうです
-> なのでこの辺うまく管理したり、つながりをすぐチェック方法を知りたいです

- 会話議事 ドメインナレッジ ドキュメント を一覧で見れる
- ナレッジグラフ的な奴がすぐ見れる
- コードからドキュメント(ビジネスの意図がわかるもの)に飛ぶ

---

# 利用規約を読んでみる
利用規約は、サービスの取り扱いやルールが極めて厳密な言い回しで表現されている。->特化型名前の参考になる

> **購入者**が商品購入手続きを完了した時点をもって、**売買契約**が**締結**されたものとします。
> **売買契約**が**締結**した場合、**出品者**は当社に**サービス使用料**を支払うものとします。
> サービス利用料は、売買契約が締結した時点の商品の**販売価格**に、**販売手数料**を乗じた金額となります。

---
# 利用規約をコードに起こしてみる

<style scoped>
  pre { font-size: 24px}
</style>

```go
type ServiceUsageFee struct {
    // 料金金額
    amount int
}

// @param
func NewServiceUsageFee(amount int) (*ServiceUsageFee, error) {
    if amount < 0 {
        return nil, errors.New("amount must be non-negative")
    }
    return &ServiceUsageFee{amount: amount}, nil
}

// サービス利用料金を確定する
func (suf ServiceUsageFee) Determine(sp SalesPrice, sc SalesCommitionRate) *ServiceUsageFee {
    amount := int(sp.amount * sc.value)
    return NewServiceUsageFee(amount)
    // Output: 確定料金
}
```

<!--
Determineメソッドは、利用規約上のサービス利用料の定義と一致している
売買契約メソッドが、このDetemineを呼べば利用規約と同じフローを実現できる

ビジネスルールとクラス名が一致していると、正確に素早く仕様変更も可能
-->

---

<!--
class: lead invert
-->

# 名づけのリスク

---

<!--
class: top
-->

---

<!--
class: lead invert
-->

# 意味が伝わりにくい名前

---

<!--
class: top
-->

---
# 後で読みたい記事
https://logmi.jp/tech/articles/322952
https://dhbr.diamond.jp/articles/-/2447?page=4

# 参考文献
https://little-hands.hatenablog.com/entry/2017/11/28/bouded-context-concept
https://little-hands.hatenablog.com/entry/2017/09/27/014403
https://little-hands.hatenablog.com/entry/2017/12/07/bouded-context-implementation
https://e-words.jp/w/%E3%82%B3%E3%83%B3%E3%83%86%E3%82%AD%E3%82%B9%E3%83%88.html
https://amzn.asia/d/0dUMsrn
