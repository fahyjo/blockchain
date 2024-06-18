package config

type Config struct {
	Network `json:"network"`
	Store   `json:"store"`
}

type Network struct {
	ListenAddr string   `json:"listenAddr"`
	Peers      []string `json:"peers"`
}

type Store struct {
	Block       string `json:"block"`
	Transaction string `json:"transaction"`
	Utxo        string `json:"utxo"`
}
