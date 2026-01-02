# GetNearRankScript

## はじめに
[GetNearRankMod](https://github.com/rakkyo150/GetNearRankMod)([GetNearRank](https://github.com/culage/GetNearRank)を参考にした、順位が近くの人が自分よりPPを多く取ったランク曲を取得してプレイリストにするMOD)のツール版。<br>
Quest2などのスタンドアローン機のみでBeat SaberをプレイしていてPCModを使えない環境の方で、WinodwsとMacを使用している方はこちらを使ってください。<br>
GetNearRankModもGetNearRankScriptも動かすことができない環境の方は、Docker環境があれば使える[GetNearRankDocker](https://github.com/rakkyo150/GetNearRankDocker)をお使いください。<br>

日本ローカルランキングでのみ正常に動きます。<br>

## 使用方法について
お使いのosとcpuにあわせたファイルをダウンロードしてください。  
Macなら`darwin`とついているものになります。  
実行後にこの実行ファイルと同一のディレクトリに`Config.json`が生成されますので、ファイルを置くディレクトリには注意してください。  
MacやLinuxをご使用の方は、ターミナルなどを使用して、ダウンロードした実行ファイルがあるディレクトリで、以下のコマンドを実行してください。
```bash
chmod +x GetNearRankScript-os-cpu
```
その後、ターミナルで以下のコマンドを実行するかダウンロードした実行ファイルをダブルクリックすることにより実行が可能です。
```bash
./GetNearRankScript-os-cpu
```

## 設定について
初回実行時に設定を入力してもらう必要があります。<br>
入力してもらった設定は実行ファイルと同じ階層に`Config.json`として保存されます。<br>
設定は半角英数字で入力してください。<br>
二回目以降の実行で設定を変えたい場合は、`Config.json`を直接書き換えてください。<br>
ちなみにどこかでエラーが出る場合、エラーが無くなるまで設定入力を永遠にループします。<br>

### 設定項目
|項目|説明|
|:---|:---|
|`YourId`|自分のScoreSaberのID(https://scoresaber.com/u/1477469542377599 なら1477469542377599)|
|`RankRange`|自分の前後何位の人を対象とするか(ローカルランキング２ページ分まで情報取得可能) **おすすめ:3**|
|`PPFilter`|何PP差以上を対象とするか **おすすめ:20**|
|`YourPageRange`|自分のトップスコア何ページ目までの情報を取得するか **おすすめ:10**|
|`OthersPageRange`|ライバルのトップスコア何ページ目までの情報を取得するか **おすすめ:3**|

ログのPP差のところにMissingDataと出力される譜面は、YourPageRangeの範囲にリザルトが無いがOthersPageRangeの範囲にリザルトがある譜面です。
