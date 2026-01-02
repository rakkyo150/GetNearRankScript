package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	YourId          string `json:"YourId"`
	RankRange       int    `json:"RankRange"`
	PPFilter        int    `json:"PPFilter"`
	YourPageRange   int    `json:"YourPageRange"`
	OthersPageRange int    `json:"OthersPageRange"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("設定の読み込みに失敗しました: %w", err)
	}

	return &cfg, nil
}

func CreateInteractiveConfig(path string) (*Config, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Input Your ID")
	yourID, err := readLine(reader)
	if err != nil {
		return nil, err
	}

	rankRange, err := readInt(reader, "Input Rank Range")
	if err != nil {
		return nil, err
	}

	ppFilter, err := readInt(reader, "Input PP Filter")
	if err != nil {
		return nil, err
	}

	yourPageRange, err := readInt(reader, "Input Your Page Range")
	if err != nil {
		return nil, err
	}

	othersPageRange, err := readInt(reader, "Input Rivals' Page Range")
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		YourId:          yourID,
		RankRange:       rankRange,
		PPFilter:        ppFilter,
		YourPageRange:   yourPageRange,
		OthersPageRange: othersPageRange,
	}

	if err := writeConfig(path, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func readLine(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(line), nil
}

func readInt(reader *bufio.Reader, prompt string) (int, error) {
	for {
		fmt.Println(prompt)
		line, err := reader.ReadString('\n')
		if err != nil {
			return 0, err
		}

		value, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			fmt.Println("数値を入力してください")
			continue
		}

		return value, nil
	}
}

func writeConfig(path string, cfg *Config) error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0o666)
}
