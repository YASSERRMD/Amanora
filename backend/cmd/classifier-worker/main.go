package main

import "github.com/YASSERRMD/Amanora/backend/internal/worker"

func main() {
	worker.RunUntilSignal("classifier-worker")
}
