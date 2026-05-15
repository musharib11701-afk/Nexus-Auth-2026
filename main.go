package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("--- Nexus-Auth-2026: Database Registration ---")

	fmt.Print("Create Username: ")
	user, _ := reader.ReadString('\n')
	user = strings.TrimSpace(user)

	fmt.Print("Create Password: ")
	pass, _ := reader.ReadString('\n')
	pass = strings.TrimSpace(pass)

	err := rdb.Set(ctx, user, pass, 0).Err()
	if err != nil {
		fmt.Println("\n❌ Error: Is the Redis 'Box' window open?")
		return
	}

	fmt.Printf("\n✅ SUCCESS! User '%s' is now in the database.\n", user)
}