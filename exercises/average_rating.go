package exercises

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//  Question :
//	Suppose you are a class monitor and you have to take the performance rating of the class teacher from your classmates.
//	Write a program to take the ratings from all of your classmates and then print the average rating.
//					Expectation: use waitGroups
//
//	Assumptions
//		- There are 200 students in the class.
//		- Every student will take a random amount of time to respond.
//		- You can simulate the random response time of your classmates by using the math/rand package.

func AverageRating() {
	numberOfStudents := 200

	var wg sync.WaitGroup
	ratings := make(chan int, numberOfStudents)

	for i := 0; i < numberOfStudents; i++ {
		wg.Add(1)
		go collectRating(ratings, &wg)
	}

	wg.Wait()
	close(ratings)

	totalRating := 0
	var allRatings []int
	for rating := range ratings {
		totalRating += rating
		allRatings = append(allRatings, rating)
	}

	averageRating := float64(totalRating) / 200
	fmt.Println("All ratings : ", allRatings)
	fmt.Printf("Average rating of the class teacher: %.2f\n", averageRating)
}

func collectRating(ratings chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	responseTime := time.Duration(rand.Intn(5)) * time.Second
	time.Sleep(responseTime)

	rating := rand.Intn(11)
	ratings <- rating
}
