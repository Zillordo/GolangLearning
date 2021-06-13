package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, "userId", 777)
	ctx = context.WithValue(ctx, "userName", "Bond")

	res, err := dbAccess(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}

	_, _ = fmt.Fprintln(w, res)
}

func dbAccess(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	ch := make(chan int)

	go pretend(ctx, ch)

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case i := <-ch:
		return i, nil
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Println(ctx)
	_, _ = fmt.Fprintln(w, ctx)
}

func pretend(ctx context.Context, ch chan int) {
	time.Sleep(2 * time.Second)
	uid := ctx.Value("userId").(int)

	if ctx.Err() != nil {
		return
	}
	ch <- uid
}
