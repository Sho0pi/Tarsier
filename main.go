package main


import (
	"os/exec"
	"time"
)

const (
	XEYES_PATH 		=	"/usr/bin/xeyes"
)

var (

)

func main()	{

}

/**
 * @brief      This function will execute the xeyes and will stop it after random time
 *
 * @param 	   timeout the amount of time to execute
 *
 * @return     an error if accure
 */
func execute(timeout time.Duration) error{

	cmd := exec.Command(XEYES_PATH)
	if err := cmd.Start(); err !=nil {
		return err
	}

	done := make(chan error, 1)
	go func(){
		done <- cmd.Wait()
	}()

	select {
	case <-time.After(timeout * time.Second):
		if err := cmd.Process.Kill(); err != nil {
			return err
		}
	case err := <-done :
		if err != nil {
			return err
		}
	}
	return nil

}