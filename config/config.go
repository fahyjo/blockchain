package config

type Config struct {
	Network
	Store
}

type Network struct {
	ListenAddr string
	Peers      []string
}

type Store struct {
	Block       string
	Transaction string
	Utxo        string
}
