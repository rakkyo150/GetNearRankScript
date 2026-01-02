package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: 15 * time.Second,
}

type SongKey struct {
	Hash       string
	Difficulty string
}

type basicPlayerResponse struct {
	CountryRank int `json:"countryRank"`
}

type playerScoresResponse struct {
	PlayerScores []struct {
		Leaderboard struct {
			SongHash   string `json:"songHash"`
			Difficulty struct {
				DifficultyRaw string `json:"difficultyRaw"`
			} `json:"difficulty"`
		} `json:"leaderboard"`
		Score struct {
			Pp float64 `json:"pp"`
		} `json:"score"`
	} `json:"playerScores"`
}

type countryRankResponse struct {
	Players []struct {
		CountryRank int    `json:"countryRank"`
		ID          string `json:"id"`
	} `json:"players"`
}

func GetYourCountryRank(cfg *Config) (int, error) {
	fmt.Println("Start GetYourCountryRankPageNumber")
	endpoint := fmt.Sprintf("https://scoresaber.com/api/player/%s/basic", cfg.YourId)
	var resp basicPlayerResponse
	if err := doGet(endpoint, &resp); err != nil {
		return 0, err
	}

	fmt.Printf("Your Local Rank %d\n", resp.CountryRank)
	return resp.CountryRank, nil
}

func GetLocalTargetedIDs(cfg *Config, yourCountryRank int) ([]string, error) {
	fmt.Println("Start GetLocalRankData")
	pageNumber := 1 + (yourCountryRank-1)/50
	baseEndpoint := fmt.Sprintf("https://scoresaber.com/api/players?page=%d&countries=jp", pageNumber)
	lowerEndpoint := fmt.Sprintf("https://scoresaber.com/api/players?page=%d&countries=jp", pageNumber+1)
	higherEndpoint := fmt.Sprintf("https://scoresaber.com/api/players?page=%d&countries=jp", pageNumber-1)

	result, err := GetCountryRankData(baseEndpoint)
	if err != nil {
		return nil, err
	}

	lowRank := yourCountryRank + cfg.RankRange
	highRank := yourCountryRank - cfg.RankRange
	if highRank <= 0 {
		highRank = 1
	}

	otherPage := false
	branchRank := 0
	for rank := highRank; rank < lowRank; rank++ {
		if rank%50 == 0 {
			otherPage = true
			branchRank = rank
		}
	}

	if otherPage {
		var second map[int]string
		if yourCountryRank > branchRank {
			second, err = GetCountryRankData(higherEndpoint)
		} else {
			second, err = GetCountryRankData(lowerEndpoint)
		}
		if err != nil {
			return nil, err
		}
		for rank, id := range second {
			result[rank] = id
		}
	}

	idSet := make(map[string]struct{})
	for i := 0; lowRank-i > yourCountryRank; i++ {
		if id, ok := result[lowRank-i]; ok {
			idSet[id] = struct{}{}
		}

		if id, ok := result[highRank+i]; ok {
			idSet[id] = struct{}{}
		}
	}

	if id, ok := result[yourCountryRank]; ok {
		delete(idSet, id)
	}

	if len(idSet) == 0 {
		return nil, fmt.Errorf("対象のIDが見つかりませんでした")
	}

	ids := make([]string, 0, len(idSet))
	for id := range idSet {
		ids = append(ids, id)
	}

	return ids, nil
}

func GetPlayResult(id string, pageRange int) (map[SongKey]float64, error) {
	playScores := make(map[SongKey]float64)
	if pageRange <= 0 {
		return playScores, nil
	}

	for page := 1; page <= pageRange; page++ {
		endpoint := fmt.Sprintf("https://scoresaber.com/api/player/%s/scores?page=%d", id, page)
		var resp playerScoresResponse
		if err := doGet(endpoint, &resp); err != nil {
			return nil, err
		}

		for _, score := range resp.PlayerScores {
			key := SongKey{
				Hash:       score.Leaderboard.SongHash,
				Difficulty: score.Leaderboard.Difficulty.DifficultyRaw,
			}
			playScores[key] = score.Score.Pp
		}
	}

	if len(playScores) == 0 {
		fmt.Printf("No %s's Play Scores\n", id)
	}

	return playScores, nil
}

func GetCountryRankData(endpoint string) (map[int]string, error) {
	rankMap := make(map[int]string)
	var resp countryRankResponse
	if err := doGet(endpoint, &resp); err != nil {
		return nil, err
	}

	for _, player := range resp.Players {
		rankMap[player.CountryRank] = player.ID
	}

	if len(rankMap) == 0 {
		fmt.Printf("No Country Rank and Id at %s\n", endpoint)
	}

	return rankMap, nil
}

func doGet(url string, target interface{}) error {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("endpoint %s returned status %d", url, resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}
