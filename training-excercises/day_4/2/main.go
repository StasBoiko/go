//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
// Hint: time.Ticker can be used
// Hint 2: to calculate timediff for Advanced lvl use:
//
//  start := time.Now()
//	// your work
// t := time.Now()
// elapsed := t.Sub(start) // 1s or whatever time has passed

package main

import (
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {

	c := make(chan *User)

	go func() {
		defer func() {
			recover()
		}()
		start := time.Now()
		process()
		t := time.Now()
		u.TimeUsed += int64(t.Sub(start))
		c <- u
	}()

	select {
	case <-c:
		return true
	case <-time.After(10 - time.Duration(u.TimeUsed)*time.Second):
		return false
	}

}

//Второе решение:
// func HandleRequest(process func(), u *User) bool {
// 	c := make(chan *User)

// 	go func() {
// 		process()
// 		c <- u
// 	}()

// 	for {
// 		select {
// 		case <-c:
// 			return true
// 		case <-time.Tick(time.Second * 1):
// 			u.TimeUsed++
// 			if u.TimeUsed > 10 {
// 				return false
// 			}
// 		}
// 	}
// }

func main() {
	RunMockServer()
}