package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func sample() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // 1行読み取り
	tmp := strings.Split(scanner.Text(), " ")
	fmt.Println(tmp)
}
