// package fileStore
/*===================================================================================*\



\*===================================================================================*/
package fileStore

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// func main() {
// 	cmdName := "ls"
// 	cmdArgs := []string{"-la", "."}
// 	ExecWithScanner(cmdName, cmdArgs)
// }
/*-----------------------------------------------------------------------------------*\
\*-----------------------------------------------------------------------------------*/
func ExecWithScanner(cmdName string, cmdArgs []string) string {
	log.Printf(">===>VidExeCmdPkg.ExecWithScanner(%s)", cmdName)
	//var stdoutbuf bytes.Buffer
	var stderrbuf bytes.Buffer

	cmd := exec.Command(cmdName, cmdArgs...)
	//cmd.Stdout = &stdoutbuf
	cmd.Stderr = &stderrbuf
	log.Printf("cmdName=%v\n", cmdName)
	log.Printf("cmdArgs=%v\n", cmdArgs)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "ExecWithScanner() Error at cmd.StdoutPipe()", err)
		fmt.Println(fmt.Sprint(err) + ": " + stderrbuf.String())
		os.Exit(1)
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		log.Printf("starting func()\n")
		cnt := 1
		for scanner.Scan() {
			fmt.Printf("%d.%s\n", cnt, scanner.Text())
			cnt = cnt + 1
		}
		fmt.Println("---- END of SCANNER ----")
	}()

	log.Printf("cmd.Start\n")
	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "ExecWirthScanner() Error at cmd.Start():err=", err)
		fmt.Println(stderrbuf.String())
		os.Exit(1)
	}

	log.Printf("cmd.Wait\n")
	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "ExecWithScanner() Error at cmd.Wait():err=", err)
		fmt.Println(stderrbuf.String())
		os.Exit(1)
	}
	return stderrbuf.String()
}
