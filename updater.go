package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func handleUpdate(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("git", "pull")

	output, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, "Git pull failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Git pull successful:\n%s", string(output))
}
