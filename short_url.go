package shor_url

import (
	"errors"
	"fmt"
	"math"
)

type UrlEncoder struct {
	alphabet  []rune
	blockSize int
	mask      int
	mapping   []int
}

const (
	DEFAULT_ALPHABET   = "mn6j2c4rv8bpygw95z7hsdaetxuk3fq"
	DEFAULT_BLOCK_SIZE = 24
	MIN_LENGTH         = 5
)

func NewUrlEncoder(alphabet string, blockSize int) (*UrlEncoder, error) {
	self := new(UrlEncoder)
	if len(alphabet) == 0 {
		alphabet = DEFAULT_ALPHABET
	}
	if len(alphabet) < 2 {
		return nil, errors.New("alphabet has to contain at least 2 characters")
	}
	self.alphabet = []rune(alphabet)
	if blockSize <= 0 {
		blockSize = DEFAULT_BLOCK_SIZE
	}
	self.blockSize = blockSize
	self.mask = (1 << blockSize) - 1
	self.mapping = make([]int, 0, 0)
	for i := 0; i < blockSize; i++ {
		self.mapping = append(self.mapping, i)
	}
	return self, nil
}

func (c *UrlEncoder) EncodeUrl(n, minLen int) string {
	if minLen <= 0 {
		minLen = MIN_LENGTH
	}
	return c.EnBase(c.Encode(n), minLen)
}

func (c *UrlEncoder) Encode(n int) int {
	return (n & (^c.mask + 1)) | c.encode(n&c.mask)
}

func (c *UrlEncoder) encode(n int) int {
	tmp := reverseInts(c.mapping)
	var result = 0
	for i, b := range tmp {
		if n&(1<<i) > 0 {
			result |= 1 << b
		}
	}
	return result
}

func (c *UrlEncoder) DecodeUrl(n []rune) int {
	return c.Decode(c.DeBase(n))
}

func (c *UrlEncoder) Decode(n int) int {
	return (n & (^c.mask + 1)) | c.decode(n&c.mask)
}

func (c *UrlEncoder) decode(n int) int {
	var result = 0
	for i, b := range reverseInts(c.mapping) {
		if n&(1<<b) > 0 {
			result |= 1 << i
		}
	}
	return result
}

func (c *UrlEncoder) EnBase(x, minLen int) string {
	if minLen <= 0 {
		minLen = MIN_LENGTH
	}
	result := c.enBase(x)
	padding := make([]rune, 0, 0)
	for i := 0; i < minLen-len(result); i++ {
		padding = append(padding, c.alphabet[0])
	}
	return fmt.Sprintf("%s%s", string(padding), string(result))
}

func (c *UrlEncoder) enBase(x int) []rune {
	n := len(c.alphabet)
	if x < n {
		return []rune{c.alphabet[x]}
	}
	out := c.enBase(int(math.Floor(float64(x) / float64(n))))
	out = append(out, c.alphabet[x%n])
	return out
}

func (c *UrlEncoder) DeBase(x []rune) int {
	n := len(c.alphabet)
	result := 0
	ints := make([]int, 0, 0)
	for _, v := range x {
		ints = append(ints, int(v))
	}
	for i, j := range reverseInts(ints) {
		val := int(math.Pow(float64(n), float64(i)))
		for ii, k := range c.alphabet {
			if j == int(k) {
				result += ii * val
			}
		}
	}
	return result
}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}
