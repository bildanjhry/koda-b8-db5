package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Io(tags string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s", tags)
	item, _ := reader.ReadString('\n')
	item = strings.ToLower(strings.Trim(item, "\n"))
	return item, nil
}
