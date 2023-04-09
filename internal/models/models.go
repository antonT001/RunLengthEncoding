package models

import "sync"

type Storage struct {
	M  map[int][]string
	Mu sync.Mutex
	Wg sync.WaitGroup
}

type Msg struct {
	Data []string `json:"data"`
}

type OutError struct {
	Success bool    `json:"success"`
	Error   *string `json:"error,omitempty"`
}

type RleResponse struct {
	Success bool     `json:"success"`
	Data    []string `json:"data"`
}
