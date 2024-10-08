package main

import (
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)

	// output, err := cmd.CombinedOutput()
	// if err != nil {
	// 	panic(fmt.Sprintf("Error running tern command: %v\n\nOutput: %s", err, string(output)))
	// }
	// fmt.Printf("Command output: %s\n", string(output))

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
