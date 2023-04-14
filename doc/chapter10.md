---
marp: true
paginate: true
theme: gaia
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

# まとめてしまった名前
あるECサイトを作ると仮定する
巨大な「商品」クラスがある

---

<!--
class: top
-->

# まとめてしまった名前
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

<!--
class: lead invert
-->

# 目的ベースで名づけする 

---

<!--
class: top
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