package bus

type Config struct {
	WorkerID string `json:"worker_id" yaml:"worker_id" env:"WORKER_ID"`
}
