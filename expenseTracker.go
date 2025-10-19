package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type Expense struct {
	Description string
	Category string
	Amount float64
}

