package config

type Config struct {
	Network   *Network  `json:"network"`
	Store     *Store    `json:"store"`
	Crypto    *Crypto   `json:"crypto"`
	Consensus Consensus `json:"consensus"`
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

type Crypto struct {
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
}

type Consensus struct {
	AmValidator bool     `json:"amValidator"`
	Validators  []string `json:"validators"`
}
