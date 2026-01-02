package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Playlist struct {
	PlaylistTitle  string `json:"playlistTitle"`
	Songs          []Song `json:"songs"`
	PlaylistAuthor string `json:"playlistAuthor"`
}

type Song struct {
	Difficulties []Difficulty `json:"difficulties"`
	Hash         string       `json:"hash"`
}

type Difficulty struct {
	Name           string `json:"name"`
	Characteristic string `json:"characteristic"`
}

func MakeLowerPPMapList(others []map[SongKey]float64, yours map[SongKey]float64, cfg *Config) []SongKey {
	seen := make(map[SongKey]struct{})
	var list []SongKey

	for _, other := range others {
		for key, otherPP := range other {
			if _, ok := seen[key]; ok {
				continue
			}

			if yourPP, has := yours[key]; has {
				if otherPP-yourPP >= float64(cfg.PPFilter) {
					seen[key] = struct{}{}
					list = append(list, key)
					fmt.Printf("%s,%s,%.2fPP\n", key.Hash, key.Difficulty, otherPP-yourPP)
				}
				continue
			}

			seen[key] = struct{}{}
			list = append(list, key)
			fmt.Printf("%s,%s, MissingData\n", key.Hash, key.Difficulty)
		}
	}

	return list
}

func MakePlaylist(list []SongKey, cfg *Config, baseDir string) error {
	fileDate := time.Now().Format("20060102")
	fileName := fmt.Sprintf("%s-RR%d-PF%d-YPR%d-OPR%d", fileDate, cfg.RankRange, cfg.PPFilter, cfg.YourPageRange, cfg.OthersPageRange)
	playlistPath := filepath.Join(baseDir, fmt.Sprintf("%s.bplist", fileName))

	playlist := Playlist{
		PlaylistTitle:  fileName,
		PlaylistAuthor: "HOGE_PLAYLIST_AUTHOR",
	}

	for _, entry := range list {
		song := Song{
			Hash: entry.Hash,
			Difficulties: []Difficulty{
				{
					Name:           difficultyName(entry.Difficulty),
					Characteristic: characteristicName(entry.Difficulty),
				},
			},
		}
		playlist.Songs = append(playlist.Songs, song)
	}

	data, err := json.MarshalIndent(playlist, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(playlistPath, data, 0o666)
}

func difficultyName(raw string) string {
	switch {
	case strings.Contains(raw, "ExpertPlus"):
		return "expertPlus"
	case strings.Contains(raw, "Expert"):
		return "expert"
	case strings.Contains(raw, "Hard"):
		return "hard"
	case strings.Contains(raw, "Normal"):
		return "normal"
	case strings.Contains(raw, "Easy"):
		return "easy"
	default:
		return ""
	}
}

func characteristicName(raw string) string {
	switch {
	case strings.Contains(raw, "Standard"):
		return "Standard"
	case strings.Contains(raw, "NoArrow"):
		return "NoArrow"
	case strings.Contains(raw, "SingleSaber"):
		return "SingleSaber"
	default:
		return ""
	}
}
