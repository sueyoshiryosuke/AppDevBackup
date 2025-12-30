// AppDevBackup
// 軽量・高速・手軽なスナップショット型バックアップツール
// 対象フォルダを日時付きで丸ごとコピーし、必要なら変更履歴を保存する。
// 除外設定は Excluded.txt または自動除外リストで管理する。
// 2025 SUEYOSHI Ryosuke

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "path/filepath"
    "strings"
    "time"
)

func main() {
    // 実行ファイルのパスと基準ディレクトリを取得
    exePath, _ := os.Executable()
    baseDir := filepath.Dir(exePath)
    bakDir := filepath.Join(baseDir, "Backup_Folder")

    os.MkdirAll(bakDir, 0755)

    // --- 変更履歴の入力 ---
    fmt.Print("変更履歴を入力してください（空ならスキップ）: ")
    reader := bufio.NewReader(os.Stdin)
    history, _ := reader.ReadString('\n')
    history = strings.TrimSpace(history)

    // --- バックアップ先フォルダ作成 ---
    timestamp := time.Now().Format("20060102_150405")
    destDir := filepath.Join(bakDir, timestamp)
    os.MkdirAll(destDir, 0755)

    // --- 除外リスト読み込み ---
    excluded := loadExcluded(filepath.Join(baseDir, "Excluded.txt"))

    // 自動除外（必須）
    excluded["AppDevBackup.exe"] = true
    excluded["Backup_Folder"] = true
    excluded["Excluded.txt"] = true

    // --- コピー対象一覧取得 ---
    entries, _ := os.ReadDir(baseDir)
    targets := []os.DirEntry{}

    for _, e := range entries {
        if excluded[e.Name()] {
            continue
        }
        targets = append(targets, e)
    }

    total := len(targets)
    copied := 0

    fmt.Println("バックアップ開始...")

    // --- GA風進捗バー ---
    barLen := 30

    for _, entry := range targets {
        name := entry.Name()
        src := filepath.Join(baseDir, name)
        dst := filepath.Join(destDir, name)

        if entry.IsDir() {
            copyDir(src, dst)
        } else {
            copyFile(src, dst)
        }

        copied++
        progress := float64(copied) / float64(total)
        filled := int(progress * float64(barLen))
        bar := strings.Repeat("=", filled) + strings.Repeat("-", barLen-filled)

        fmt.Printf("\r[%s] %.1f%% (%d/%d)", bar, progress*100, copied, total)
    }

    // --- Change-history.txt 作成 ---
    if history != "" {
        f, _ := os.Create(filepath.Join(destDir, "Change-history.txt"))
        defer f.Close()
        f.WriteString(history + "\n")
    }

    fmt.Printf("\n完了しました！ → %s\n", destDir)
    fmt.Println("Enter を押して終了...")
    fmt.Scanln()
}

// loadExcluded
// Excluded.txt を読み込み、除外対象を map[string]bool で返す。
// ファイルが無い場合は空のマップを返す。
func loadExcluded(path string) map[string]bool {
    m := make(map[string]bool)
    f, err := os.Open(path)
    if err != nil {
        return m
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line != "" {
            m[line] = true
        }
    }
    return m
}

// copyFile
// 単一ファイルをコピーする。
// エラーは無視して静かにスキップする（軽量ツールのため）。
func copyFile(src, dst string) {
    in, err := os.Open(src)
    if err != nil {
        return
    }
    defer in.Close()

    out, err := os.Create(dst)
    if err != nil {
        return
    }
    defer out.Close()

    io.Copy(out, in)
}

// copyDir
// ディレクトリを再帰的にコピーする。
// Walk を使ってフォルダ構造をそのまま再現する。
func copyDir(src, dst string) {
    filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        rel, _ := filepath.Rel(src, path)
        target := filepath.Join(dst, rel)

        if info.IsDir() {
            os.MkdirAll(target, 0755)
        } else {
            copyFile(path, target)
        }
        return nil
    })
}
