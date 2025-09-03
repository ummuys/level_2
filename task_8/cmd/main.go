package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	srv := flag.String("server", "time.google.com", "NTP server")
	flag.Parse()

	t, err := ntp.Time(*srv)
	if err != nil {
		fmt.Fprintln(os.Stderr, "NTP error: ", err)
		os.Exit(1)
	}

	fmt.Println(t.Format(time.RFC3339Nano)) // На разных OC разный вывод, поэтому приводим к одному стандарту
}
