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
