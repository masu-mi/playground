# 要求確認メモ

Q. カードゲームのデッキはどんなカードを扱うか?
候補: トランプ, ドミニオン, ボーナンザ, UNO
A. トランプ

Q. スートはフレンチ?
A. OK(Clubs, Diamonds, Hearts, Spades)

Q. デッキでサポートする処理は?
A. シャッフル, ソート, Push, Pop, 中身を確認する

サポート外: カット

Q. カードにユニークなど制約はるか?
A. ない(1種類1枚が一般的だけど制約ではない) -> たしかにDeckの責任範囲ではなさそう

Q. Jokerは1種類?
A. ExtraJokerはない: 途中で増えるかも?


## 設計方針メモ

Q. Goで実装する際カードの定義にinterfaceを使うか?

すくなくとも最初は使わない。
まず、ゲーム内でカードの同値判断・表示など実際の値を参照することが多い。
また、いまのところ他のカードなどを考慮しないため。

Blackjack実装では使った。
カードは値として考えており同値性なども内容で対応したため構造体、クラスベースOOPであれば普通のクラスやdataクラスを選んだ。

いろいろなカードゲームがあるが、それらを組み合わせてチップの枚数を競うことを想定してをゲームをインタフェースとして定義した。
その際に、詳細には関心がないため各種ゲームの終了時にチップ変動を返すというインタフェースにした。

ゲームの生成で多態を実現したいためクラスベースOOPでのAbstractFactoryを使う。
Goでは関数スキーマに型名を与えられるためGameFactoryという型で関数型を定めた。
オブジェクトにしなかったのはゲームインスタンス生成に関わる状態が共通のプレイヤー以外にないため。
