// package fileStore
/*===================================================================================*\



\*===================================================================================*/
package fileStore

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func r() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from error")
	} else {
		fmt.Println("Could NOT Recover from error")
	}
}

// SignalWithPid ...
/*-----------------------------------------------------------------------------------*\
\*-----------------------------------------------------------------------------------*/
func SignalWithPid(pid int) error {
	log.Printf(">===>VidExeCmdPkg.SignalWithPid(%d)", pid)
	proc, err := os.FindProcess(pid)

	if err != nil {
		log.Printf("failed to find process: reason=%v\n", err)
		return err
	}

	err = proc.Signal(syscall.SIGINT)

	if err != nil {
		log.Printf("failed to signal process: reason=", err)
		return err
	}
	return nil
}

// ExecReturnPid ...
/*-----------------------------------------------------------------------------------*\
\*-----------------------------------------------------------------------------------*/
func ExecReturnPid(cmdName string, cmdArgs []string) (int, error) {
	log.Printf(">===>VidExeCmdPkg.ExecReturnPid(%s) \n", cmdName)
	cmd := exec.Command(cmdName, cmdArgs...)

	if err := cmd.Start(); err != nil {
		log.Printf("Error ExecReturnPid %v\n", err)
		return 0, err
	}

	//log.Printf("cmd=%v\n", cmd)
	log.Printf("cmd.Process.Pid=%v\n", cmd.Process.Pid)
	return cmd.Process.Pid, nil
}

// KillWithPid ...
/*-----------------------------------------------------------------------------------*\
\*-----------------------------------------------------------------------------------*/
func KillWithPid(pid int) error {
	log.Printf(">===>VidExeCmdPkg.KillWithPid(%d)", pid)
	proc, err := os.FindProcess(pid)

	if err != nil {
		log.Printf("failed to find process: reason=%v\n", err)
		return err
	}

	err = proc.Kill()

	if err != nil {
		log.Printf("failed to kill process: reason=%v\n", err)
		return err
	}
	return nil
}

// ExecThenKill ...
/*-----------------------------------------------------------------------------------*\
\*-----------------------------------------------------------------------------------*/
func ExecThenKill(cmdName string, cmdArgs []string) error {
	log.Printf(">===>VidExeCmdPkg.ExecThenKill(%s)", cmdName)
	// Start a process:
	cmd := exec.Command(cmdName, cmdArgs...)

	if err := cmd.Start(); err != nil {
		log.Printf("ERROR ExecThenKill %v\n", err)
		return err
	}

	log.Printf("cmd=%v\n", cmd)
	log.Printf("cmd.Process.Pid=%v\n", cmd.Process.Pid)

	// Wait for the process to finish or kill it after a timeout:
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	select {
	case <-time.After(10 * time.Second):
		if err := cmd.Process.Kill(); err != nil {
			log.Printf("failed to kill process: %v\n", err)
			return err
		}
		log.Println("process killed as timeout reached")
	case err := <-done:
		if err != nil {
			log.Printf("process finished with error =  %v\n", err)
			return err
		}
		log.Print("process finished successfully")
	}
	return nil
}
