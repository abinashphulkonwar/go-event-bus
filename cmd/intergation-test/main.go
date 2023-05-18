package main

import (
	"bytes"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/abinashphulkonwar/go-event-bus/client"
)

const marks = 1000000
const newBufferSize = 1024 * 1024

func main() {
	println("start")
	c := client.NewClient("localhost:8080")

	sum, err := send(c)

	if err != nil {
		log.Fatal("send error: %\n", err)
	}

	println(sum)

	got, err := recive(c)

	if err != nil {
		log.Fatal("recive error: %\n", err)
	}

	if sum != got {
		log.Fatal("sum != got")
	}

	log.Printf("sum: %d, got: %d\n", sum, got)
}

func send(c *client.Client) (sum int64, err error) {
	var b bytes.Buffer

	for i := 0; i < marks; i++ {
		sum += int64(i)
		if b.Len() >= newBufferSize {
			err := c.Send(b.Bytes())
			if err != nil {
				return 0, err
			}
		}
	}

	if b.Len() != 0 {
		err := c.Send(b.Bytes())
		if err != nil {
			return 0, err
		}
	}
	return sum, nil
}

func recive(c *client.Client) (sum int64, err error) {
	buf := make([]byte, newBufferSize)
	for {
		res, err := c.Recive(buf)

		if err == io.EOF {
			return sum, err
		}
		if err != nil {
			return 0, err
		}

		inta := strings.Split(string(res), "\n")

		for _, str := range inta {
			if str == "" {
				continue
			}

			i, err := strconv.Atoi(str)

			if err != nil {
				return 0, err
			}

			sum += int64(i)
		}

	}
}
