package main


import (
	"time"
	"math/rand"
)

const (
	XEYES_PATH 		=	"/usr/bin/xeyes"
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