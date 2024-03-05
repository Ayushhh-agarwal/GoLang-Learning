package exercises

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//  Question :
//	Write a program that starts with a single bank account with a starting balance of Rs.500. It should be
//	possible to deposit and withdraw money concurrently. However, the balance update should happen only once
//	at any point in time. If the withdrawal action should fail if the balance is reaching negative.
//
//																	Expectations: Use Goroutines, Mutex

func withdrawAmount(balance *int, withdrawArray []int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	for _, withdrawAmt := range withdrawArray {
		mutex.Lock()
		if *balance < withdrawAmt {
			fmt.Printf("Insufficient Funds to Withdraw %d as Balance is %d\n", withdrawAmt, *balance)
			mutex.Unlock()
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			continue
		}
		*balance = *balance - withdrawAmt
		fmt.Printf("Amount Withdrawn : %d | Balance Left : %d\n", withdrawAmt, *balance)
		mutex.Unlock()
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	}
}

func depositAmount(balance *int, depositArray []int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	for _, depositedAmt := range depositArray {
		mutex.Lock()
		*balance = *balance + depositedAmt
		fmt.Printf("Amount Deposited : %d | Updated Balance : %d\n", depositedAmt, *balance)
		mutex.Unlock()
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	}
}

func BankAccount() {
	bankAccountBalance := 500
	fmt.Println("Starting Balance: ", bankAccountBalance)

	withdrawArray := []int{200, 150, 300, 750, 350}
	depositArray := []int{100, 250, 300, 50, 400}

	wg := sync.WaitGroup{}

	mutex := sync.Mutex{}
	wg.Add(2)
	go withdrawAmount(&bankAccountBalance, withdrawArray, &wg, &mutex)
	go depositAmount(&bankAccountBalance, depositArray, &wg, &mutex)

	wg.Wait()

	fmt.Println("Closing Balance: ", bankAccountBalance)
}
