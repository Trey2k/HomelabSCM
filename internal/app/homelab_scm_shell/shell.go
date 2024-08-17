package homelab_scm_shell

import (
	"fmt"
	"os"
	"os/exec"
)

func Run() error {
    // Get the SSH_ORIGINAL_COMMAND environment variable
    sshOriginalCommand := os.Getenv("SSH_ORIGINAL_COMMAND")
    if sshOriginalCommand == "" {
        return fmt.Errorf("SSH_ORIGINAL_COMMAND is not set")
    }
    
    repos_path := "/var/opt/homelab-scm/git-data"
    os.Chdir(repos_path)
    fullArgs := []string{"-c", sshOriginalCommand}

    git_shell_path, err := exec.LookPath("git-shell")
    if err != nil {
        return err
    }

    cmd := exec.Command(git_shell_path, fullArgs...)
    cmd.Dir = repos_path
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}