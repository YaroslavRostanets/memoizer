package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"time"
)

var actions = []string{"logged in", "logged out", "created record", "deleted record", "updated account"}

type logItem struct {
	action    string
	timestamp time.Time
}

type User struct {
	id    int
	email string
	logs  []logItem
}

func (u User) getActivityInfo() string {
	output := fmt.Sprintf("UID: %d; Email: %s;\nActivity Log:\n", u.id, u.email)
	for index, item := range u.logs {
		output += fmt.Sprintf("%d. [%s] at %s\n", index, item.action, item.timestamp.Format(time.RFC3339))
	}

	return output
}

func worker(jobs <-chan int, results chan<- User) {
	for j := range jobs {
		results <- generateUser(j)
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	startTime := time.Now()

	fmt.Println("Starting...")

	const numJobs = 100
	jobs := make(chan int, numJobs)
	results := make(chan User, numJobs)

	for w := 1; w <= runtime.NumCPU(); w++ {
		go worker(jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	for a := 1; a <= numJobs; a++ {
		u := <-results
		go saveUserInfo(u)
	}
	close(jobs)

	fmt.Printf("DONE! Time Elapsed: %.2f seconds\n", time.Since(startTime).Seconds())
}

func saveUserInfo(user User) {
	fmt.Printf("WRITING FILE FOR UID %d\n", user.id)

	filename := fmt.Sprintf("users/uid%d.txt", user.id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString(user.getActivityInfo())
	time.Sleep(time.Second)
}

func generateUser(idx int) User {
	time.Sleep(time.Millisecond * 100)
	return User{
		id:    idx,
		email: fmt.Sprintf("user%d@company.com", idx),
		logs:  generateLogs(rand.Intn(1000)),
	}
}

func generateLogs(count int) []logItem {
	logs := make([]logItem, count)

	for i := 0; i < count; i++ {
		logs[i] = logItem{
			action:    actions[rand.Intn(len(actions)-1)],
			timestamp: time.Now(),
		}
	}

	return logs
}
