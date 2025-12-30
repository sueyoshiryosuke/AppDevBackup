# AppDevBackup
軽量・高速・手軽なスナップショット型バックアップツール

## 概要
AppDevBackup は、**フォルダを丸ごと日時付きでバックアップするだけ**の  
超シンプルなバックアップツールです。  
Git のような大げさな管理は不要だけど、最低限の履歴は残したい、  
そんな個人開発者向けに作られています。

- バックアップ先は `Backup_Folder/YYYYMMDD_HHMMSS/`
- 変更履歴を一言メモとして保存可能
- 除外したいファイル・フォルダは `Excluded.txt` に記述
- 進捗バー付き
- 設定不要・実行するだけの軽量ツール

## 使い方

### 1. AppDevBackup.exe を任意のフォルダに置く
このフォルダがバックアップ対象になります。

### 2. AppDevBackup.exe を実行
起動すると、変更履歴の入力を求められます。

- 入力した場合 → `Change-history.txt` がバックアップ先に作成されます  
- 空の場合 → 履歴ファイルは作成されません

### 3. バックアップ完了
`Backup_Folder/YYYYMMDD_HHMMSS/` に  
対象フォルダの内容が丸ごとコピーされます。
  
進捗はバーで表示されます。

## 除外設定

### 自動で除外されるもの
以下は Excluded.txt に書かなくても常に除外されます。

- `AppDevBackup.exe`
- `Backup_Folder`
- `Excluded.txt`

### 任意で除外したい場合
同じ階層に `Excluded.txt` を置き、  
除外したいファイル名・フォルダ名を1行ずつ書きます。

例:

```
node_modules
temp
debug.log
```

## バイナリのビルド方法（Go）

```
go build -ldflags="-s -w" -o AppDevBackup.exe main.go
```

`-s -w` によりデバッグ情報を削除し、バイナリを軽量化します。

## ライセンス
MIT License  

## 開発者
SUEYOSHI Ryosuke

## プロジェクトページ
https://github.com/sueyoshiryosuke/AppDevBackup

## 変更履歴
### Ver.20251231
- 新規作成

以上です。
