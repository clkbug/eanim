# eanim

## Install

```bash
go install github.com/clkbug/eanim
```

## Usage

```bash
eanim a.png b.png c.png # 直接指定
eanim ./img/            # ディレクトリ指定（ディレクトリ内の画像ファイルすべてが対象）
eanim                   # 無指定（カレントディレクトリ内の画像ファイル全てが対象）
```

キー操作

* スペース：再生・停止
* 右・左：次の画像・前の画像
* 上・下：再生速度増加・減少