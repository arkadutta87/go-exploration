package main

import (
	"fmt"
	"go-arka-practice/chapterone"
	"go-arka-practice/mathutility"
	redis "go-arka-practice/redisclusterutil"
	"go-arka-practice/xmlparsing"
	"math"
	"math/cmplx"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

const (
	secretCodeLength int = 6
)

// var c, python, java bool
// var i, j = 1, 2

var (
	//ToBe ... boolean variable
	ToBe              = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println("Hello , Arka Dutta ! Start with go-lang")
	fmt.Println("My favorite number is ", rand.Intn(100))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(chapterone.Add(20, 90))

	a, b := chapterone.Swap("Arka", "Dutta")
	fmt.Println(a, b)

	fmt.Println(chapterone.Split(17))

	// var c, python, java = true, false, "no!"
	// var i, j = 1, 2
	// k := 3
	// c, python, java := true, false, "no!"
	// fmt.Println(i, j, k, c, python, java)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var i int
	var f float64
	var be bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, be, s)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	fmt.Println(chapterone.SumXNaturalNumbers(10))

	chapterone.PointerTesting()

	chapterone.PrintVertex(5, 9)
	chapterone.ManipulateVertex(9, 9)

	chapterone.InterfaceOne(chapterone.T{S: "arka"})

	elements := []int{10, 2, 30, 40, 5, 6, -1, -9, 12, 11, 13, 11}

	chapterone.MergeSort(elements)

	fmt.Println(elements)

	secretCode := mathutility.GenerateRandomSecretCode(secretCodeLength)
	xmlPostBody := xmlparsing.SendMailXMLPreparation(secretCode)

	fmt.Println("SMS body created : ", xmlPostBody)
	// restclient.SendSms(xmlPostBody)
	redis.InitV2()

	redisKey := "Arka_Life"
	err := redis.SetV2(redisKey, "Saiesha-Titli-Jhumur")
	if err != nil {
		fmt.Println("Redis Connection not set properly")
	}

	val, err := redis.GetV2(redisKey)
	if err != nil {
		fmt.Println("The Value for the key couldnot be read - ", err.Error())
	} else {
		fmt.Println("Read from Redis Cluster : ", val)
	}

	// err = redis.Close()
	// if err != nil {
	// 	fmt.Println("Error Closing the Connection to the Redis Cluster - ", err.Error())
	// }

	uniqueValue := "ArkaDutta0504"
	key := "referral_client_123"

	for i := 0; i < 10; i++ {

		go func() {
			//generate a unique UID
			uuid1, err := uuid.NewUUID()
			if err != nil {
				fmt.Println("could not create UUID: ", err)
			}

			isLocked, errV2 := redis.Lock(key, uuid1.String(), 15000)
			if errV2 != nil {
				fmt.Println("Error while trying to acquire the lock - ", errV2.Error(), isLocked)
				return
			}
			if isLocked {
				fmt.Println("Acquiring the lock succeded: Hurray")
				fmt.Println("Sleeping for 10 seconds")
				time.Sleep(10000 * time.Millisecond)
				redis.Unlock(key, uniqueValue)
			} else {
				fmt.Printf("Go Routine with UUID - %v couldnot acquire lock\n", uuid1)
			}

			// try acquiring a lock
			// If lock acquired sleep for random time : If not log that
			// release the lock if you reach here
		}()

	}

	time.Sleep(30000 * time.Millisecond)

}
