package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var chesl map[string]int = map[string]int{
	"zero":      0,
	"un":        1,
	"deux":      2,
	"trois":     3,
	"quatre":    4,
	"cinq":      5,
	"six":       6,
	"sept":      7,
	"huit":      8,
	"neuf":      9,
	"onze":      11,
	"douze":     12,
	"treize":    13,
	"quatorze":  14,
	"quinze":    15,
	"seize":     16,
	"vingt":     20,
	"trente":    30,
	"quarante":  40,
	"cinquante": 50,
	"soixante":  60,
	"dix":       10,
	"cent":      100,
}

func main() {
	router := gin.Default()

	router.Static("/static", "./static")

	router.GET("/convert", func(c *gin.Context) {
		inp := c.Query("number")

		words := strings.Fields(inp)
		ans := 0

		p := check(words[0])
		n := len(words)
		if p != 0 {
			ans += betw(p, n) + dix(p, words, n) + des(p, words, n) + four(p, words, n) + cent(p, words, n) + edin(p, words, n)
		}

		c.JSON(http.StatusOK, gin.H{"result": ans})
	})

	router.Run(":8080")

}

func check(w string) int {
	p, ok := chesl[w]

	if !ok {
		log.Fatal("Ошибка в синтаксисе!")
	}

	return p
}

func betw(p int, n int) int {
	if p > 10 && p < 17 {
		n--
		if n == 0 {
			ans := p
			return ans
		} else {
			log.Fatal("После числа от 11 до 16 не может идти что-то еще!")
		}
	}
	return 0
}

func dix(p int, words []string, n int) int {
	if p == 10 {
		n--
		ans := p

		if n == 0 {
			return ans
		}

		p = check(words[1])
		n--

		if p > 6 && p < 10 {
			ans += p
			if n != 0 {
				log.Fatal("После единиц не могут идти числа!")
			}
		} else {
			log.Fatal("После десятки следует только число от 7 до 9!")
		}
		return ans
	}
	return 0
}

func des(p int, words []string, n int) int {
	if p >= 20 && p <= 50 {
		n--
		ans := p
		if n == 0 {
			return ans
		}

		p = check(words[1])
		n--
		if p > 0 && p < 10 {
			ans += p
			if n != 0 {
				log.Fatal("После единиц не могут идти числа!")
			}
		} else {
			log.Fatal("После десяток идут только единицы!")
		}
		return ans
	}
	if p == 60 {
		n--
		ans := p
		if n == 0 {
			return ans
		}

		p = check(words[1])
		n--

		if p > 0 && p < 10 {
			if n != 0 {
				log.Fatal("После единиц не могут идти числа!")
			}
			ans += p
		} else {
			n = len(words) - 1
			ans += dix(p, words[1:], n) + betw(p, n)
		}
		return ans
	}

	return 0
}

func four(p int, words []string, n int) int {
	if p == 4 {
		n--
		ans := p
		if n == 0 {
			return ans
		}

		p = check(words[1])
		n--
		if p == 20 {
			ans += 80 - 4
			if n == 0 {
				return ans
			}

			p = check(words[2])
			n--

			if p > 0 && p < 10 {
				if n != 0 {
					log.Fatal("После единиц не могут идти числа!")
				}
				ans += p
			} else {
				n = len(words) - 2
				ans += dix(p, words[2:], n) + betw(p, n)
			}
		}
		return ans

	}
	return 0
}

func cent(p int, words []string, n int) int {
	if p == 100 {
		ans := 100
		n--

		if n == 0 {
			return ans
		}

		p = check(words[1])
		n--
		if p == 100 {
			log.Fatal("После сотней не могут идти сотни")
		}
		if p > 0 && p < 10 && p != 4 {
			if n != 0 {
				log.Fatal("После единиц не могут идти числа!")
			}
			ans += p
		} else {
			n = len(words) - 1
			ans += des(p, words[1:], n) + dix(p, words[1:], n) + betw(p, n) + four(p, words[1:], n)
		}
		return ans
	}
	return 0
}

func edin(p int, words []string, n int) int {
	if p > 0 && p < 10 && p != 4 {
		n--
		ans := p
		if n == 0 {
			return ans
		}
		p = check(words[1])
		n--
		if p != 100 {
			log.Fatal("После единиц могут быть только сотни!")
		}
		n = len(words) - 1
		ans = 100 * ans
		ans += cent(p, words[1:], n) - 100
		return ans
	}
	return 0
}
