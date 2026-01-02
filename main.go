package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	baseDir := baseDirectory()
	configPath := filepath.Join(baseDir, "Config.json")

	cfg, err := ensureConfig(configPath)
	if err != nil {
		log.Fatalf("設定を読み込めません: %v", err)
	}

	log.Println("Start")
	if err := generatePlaylist(cfg, baseDir); err != nil {
		log.Fatalf("プレイリスト生成に失敗しました: %v", err)
	}
}

func generatePlaylist(cfg *Config, baseDir string) error {
	log.Println("Getting Your Local Rank")
	yourRank, err := GetYourCountryRank(cfg)
	if err != nil {
		return err
	}

	log.Println("Getting Rivals' ID")
	targetedIDs, err := GetLocalTargetedIDs(cfg, yourRank)
	if err != nil {
		return err
	}

	log.Println("Getting Your Play Results")
	yourScores, err := GetPlayResult(cfg.YourId, cfg.YourPageRange)
	if err != nil {
		return err
	}

	log.Println("Getting Rivals' Play Results")
	var others []map[SongKey]float64
	for _, rivalID := range targetedIDs {
		log.Printf("Targeted Id %s\n", rivalID)
		rivalScores, err := GetPlayResult(rivalID, cfg.OthersPageRange)
		if err != nil {
			return fmt.Errorf("ライバル%sのスコア取得に失敗しました: %w", rivalID, err)
		}
		others = append(others, rivalScores)
	}

	log.Println("Making Lower PP Map List")
	list := MakeLowerPPMapList(others, yourScores, cfg)

	log.Println("Making Playlist")
	if err := MakePlaylist(list, cfg, baseDir); err != nil {
		return err
	}

	log.Println("Success!")
	return nil
}

func ensureConfig(path string) (*Config, error) {
	cfg, err := LoadConfig(path)
	if err == nil {
		return cfg, nil
	}

	if errors.Is(err, os.ErrNotExist) {
		return CreateInteractiveConfig(path)
	}

	fmt.Println("Configファイルの読み込みでエラーが発生したため、再作成します。")
	return CreateInteractiveConfig(path)
}

func baseDirectory() string {
	if exePath, err := os.Executable(); err == nil {
		if dir := filepath.Dir(exePath); dir != "" {
			return dir
		}
	}

	if cwd, err := os.Getwd(); err == nil {
		return cwd
	}

	return "."
}
