package main


import (
	"time"
	"math/rand"
	"os/exec"
)

const (
	XEYES_PATH 			=	"/usr/bin/xeyes"

	EXECUTION_INTERVAL 	= 	5
)

var (

)

func main()	{

}

/**
 * @brief      For a given maxTimeout will return a random time.Duration
 *
 * @param      maxTimeout  The maximum timeout
 *
 * @return     a time within that range (0,maxTimeout]
 */
func generateRandomTimeout(maxTimeout int) time.Duration {
	rand.Seed(generateRandomSeed())
	n := rand.Intn(maxTimeout)
	return time.Duration(maxTimeout - n)
}

/**
 * @brief      This function will generate random seed from the ephoktime
 *
 * @return     the current UNIX time
 */
func generateRandomSeed() int64{
	return time.Now().Unix()
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

/**
 * @brief      Determines if for execution.
 *
 * @return     True if for execution, False otherwise.
 */
func isForExecution() bool {
	return generateRandomSeed() % EXECUTION_INTERVAL == 0
}