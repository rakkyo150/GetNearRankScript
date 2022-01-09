# GetNearRankScript
[GetNearRankMod](https://github.com/rakkyo150/GetNearRankMod)(GetNearRankを参考にした、順位が近くの人が自分よりPPを多く取ったランク曲を取得してプレイリストにするMOD)のツール版。<br>
QUESTなどのスタンドアローン機のみでBeat Saberをしている人など、PCModを使えない環境の人はこちらを使ってください。<br>
WindowsとMacで利用できます。<br>
日本ローカルランキングでのみ正常に動きます。<br>

## 環境について
Windowsで利用する場合は.NET 6が必要です。<br>
.NET 6がない場合はエラーがでますが、出てくる指示に従ってインストールすれば使えるようになるはずです。<br>
Macについては、.NET 6との相性が良いとはいえないので、.NET 6のインストールが必要なものとそうでないものの両方を用意しました。<br>
.NET 6が必要なものは無印、不要なものは”self-contained”と名付けてます<br>
環境汚染を避けたい場合は後者でも構いませんが、少し容量が大きいです。<br>

## 設定について
初回実行時に設定を入力してもらう必要があります。<br>
入力してもらった設定は実行ファイルと同じ階層にConfig.jsonとして保存されます。<br>
設定は半角英数字で入力してください。<br>
二回目以降の実行で設定を変えたい場合は、Config.jsonを直接書き換えてください。<br>
ちなみにどこかでエラーが出る場合、エラーが無くなるまで設定入力を永遠にループします。<br>

### 設定項目
|項目|説明|
|:---|:---|
|`YourId`|自分のScoreSaberのID(https://scoresaber.com/u/1477469542377599なら1477469542377599)|
|`RankRange`|自分の前後何位の人を対象とするか(ローカルランキング２ページ分まで情報取得可能)|
|`PPFilter`|何PP差を対象とするか|
|`YourPageRange`|自分のトップスコア何ページ目までの情報を取得するか|
|`OthersPageRange`|ライバルのトップスコア何ページ目までの情報を取得するか|

ログのPP差のところにMissingDataと出力される譜面は、YourPageRangeの範囲にリザルトが無いがOthersPageRangeの範囲にリザルトがある譜面です。
