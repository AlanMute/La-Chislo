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

		p, er := check(words[0])
		if p == -1 {
			c.JSON(http.StatusOK, gin.H{"result": er})
			return
		}
		n := len(words)
		if p == 0 && n != 1 {
			c.JSON(http.StatusOK, gin.H{"result": "Ноль не может быть с другими числами!"})
			return
		}
		if p != 0 {
			h, er := betw(p, n)
			if h == -1 {
				//fmt.Println("1")
				c.JSON(http.StatusOK, gin.H{"result": er})
				return
			}
			ans += h
			h, er = dix(p, words, n)
			if h == -1 {
				c.JSON(http.StatusOK, gin.H{"result": er})
				return
			}
			ans += h
			h, er = des(p, words, n)
			if h == -1 {
				c.JSON(http.StatusOK, gin.H{"result": er})
				return
			}
			ans += h
			h, er = four(p, words, n)
			if h == -1 {
				c.JSON(http.StatusOK, gin.H{"result": er})
				return
			}
			ans += h
			h, er = cent(p, words, n)
			if h == -1 {
				c.JSON(http.StatusOK, gin.H{"result": er})
				return
			}
			ans += h
			h, er = edin(p, words, n)
			if h == -1 {
				c.JSON(http.StatusOK, gin.H{"result": er})
				return
			}
			ans += h
		}

		c.JSON(http.StatusOK, gin.H{"result": ans})
	})

	router.Run(":8080")

}

func check(w string) (int, string) {
	p, ok := chesl[w]

	if !ok {
		return -1, "Ошибка в слове " + w
	}

	return p, ""
}

func betw(p int, n int) (int, string) {
	if p > 10 && p < 17 {
		n--
		if n == 0 {
			ans := p
			return ans, ""
		} else {
			log.Fatal("После числа от 11 до 16 не может идти что-то еще!")
		}
	}
	return 0, ""
}

func dix(p int, words []string, n int) (int, string) {
	if p == 10 {
		n--
		ans := p

		if n == 0 {
			return ans, ""
		}

		p, er := check(words[1])
		if p == -1 {
			return -1, er
		}
		n--

		if p > 6 && p < 10 {
			ans += p
			if n != 0 {
				p, er := check(words[2])
				if p == -1 {
					return -1, er
				}
				if p > 0 && p < 10 {
					return -1, "После единиц не могут идти единицы!"
				}
				if p >= 10 && p < 100 {
					return -1, "После единиц не могут идти десятки"
				}
				if p >= 100 {
					return -1, "После единиц не могут идти сотни"
				}
				if p == 0 {
					return -1, "После единиц не могут идти нули!"
				}

			}
		} else {
			if p >= 10 && p < 100 {
				return -1, "После десяток не могут идти десятки"
			}
			if p >= 100 {
				return -1, "После десяток не могут идти сотни"
			}
			if p == 0 {
				return -1, "После десяток не могут идти нули!"
			}
		}
		return ans, ""
	}
	return 0, ""
}

func des(p int, words []string, n int) (int, string) {
	if p >= 20 && p <= 50 {
		n--
		ans := p
		if n == 0 {
			return ans, ""
		}

		p, er := check(words[1])
		if p == -1 {
			return -1, er
		}

		n--
		if p > 0 && p < 10 {
			ans += p
			if n != 0 {
				p, er := check(words[2])
				if p == -1 {
					return -1, er
				}
				if p > 0 && p < 10 {
					return -1, "После единиц не могут идти единицы!"
				}
				if p >= 10 && p < 100 {
					return -1, "После единиц не могут идти десятки"
				}
				if p >= 100 {
					return -1, "После единиц не могут идти сотни"
				}
				if p == 0 {
					return -1, "После единиц не могут идти нули!"
				}
			}
		} else {
			if p >= 10 && p < 100 {
				return -1, "После десяток не могут идти десятки"
			}
			if p >= 100 {
				return -1, "После десяток не могут идти сотни"
			}
			if p == 0 {
				return -1, "После десяток не могут идти нули!"
			}
		}
		return ans, ""
	}
	if p == 60 {
		n--
		ans := p
		if n == 0 {
			return ans, ""
		}

		p, er := check(words[1])
		if p == -1 {
			return -1, er
		}

		n--

		if p > 0 && p < 10 {
			if n != 0 {
				p, er := check(words[2])
				if p == -1 {
					return -1, er
				}
				if p > 0 && p < 10 {
					return -1, "После единиц не могут идти единицы!"
				}
				if p >= 10 && p < 100 {
					return -1, "После единиц не могут идти десятки"
				}
				if p >= 100 {
					return -1, "После единиц не могут идти сотни"
				}
				if p == 0 {
					return -1, "После единиц не могут идти нули!"
				}
			}
			ans += p
		} else {
			n = len(words) - 1
			a, er := dix(p, words[1:], n)
			if a == -1 {
				return -1, er
			}
			ans += a
			a, er = betw(p, n)
			if a == -1 {
				return -1, er
			}
			ans += a
		}
		return ans, ""
	}

	return 0, ""
}

func four(p int, words []string, n int) (int, string) {
	if p == 4 {
		n--
		ans := p
		if n == 0 {
			return ans, ""
		}

		p, er := check(words[1])
		if p == -1 {
			return -1, er
		}
		n--
		if p == 20 {
			ans += 80 - 4
			if n == 0 {
				return ans, ""
			}

			p, er := check(words[2])
			if p == -1 {
				return -1, er
			}
			n--

			if p > 0 && p < 10 {
				if n != 0 {
					p, er := check(words[3])
					if p == -1 {
						return -1, er
					}
					if p > 0 && p < 10 {
						return -1, "После единиц не могут идти единицы!"
					}
					if p >= 10 && p < 100 {
						return -1, "После единиц не могут идти десятки"
					}
					if p >= 100 {
						return -1, "После единиц не могут идти сотни"
					}
					if p == 0 {
						return -1, "После единиц не могут идти нули!"
					}
				}
				ans += p
			} else {
				n = len(words) - 2
				a, er := dix(p, words[2:], n)
				if a == -1 {
					return -1, er
				}
				ans += a
				a, er = betw(p, n)
				if a == -1 {
					return -1, er
				}
				ans += a
			}
		}
		return ans, ""

	}
	return 0, ""
}

func cent(p int, words []string, n int) (int, string) {
	if p == 100 {
		ans := 100
		n--

		if n == 0 {
			return ans, ""
		}

		p, er := check(words[1])
		if p == -1 {
			return -1, er
		}
		n--
		if p == 100 {
			return -1, "После сотней не могут идти сотни"
		}
		if p > 0 && p < 10 && p != 4 {
			if n != 0 {
				p, er := check(words[2])
				if p == -1 {
					return -1, er
				}
				if p > 0 && p < 10 {
					return -1, "После единиц не могут идти единицы!"
				}
				if p >= 10 && p < 100 {
					return -1, "После единиц не могут идти десятки"
				}
				if p >= 100 {
					return -1, "После единиц не могут идти сотни"
				}
				if p == 0 {
					return -1, "После единиц не могут идти нули!"
				}
			}
			ans += p
		} else {
			n = len(words) - 1
			a, er := des(p, words[1:], n)
			if a == -1 {
				return -1, er
			}
			ans += a

			a, er = dix(p, words[1:], n)
			if a == -1 {
				return -1, er
			}
			ans += a

			a, er = betw(p, n)
			if a == -1 {
				return -1, er
			}
			ans += a

			a, er = four(p, words[1:], n)
			if a == -1 {
				return -1, er
			}
			ans += a
		}
		return ans, ""
	}
	return 0, ""
}

func edin(p int, words []string, n int) (int, string) {
	if p > 0 && p < 10 && p != 4 {
		n--
		ans := p
		if n == 0 {
			return ans, ""
		}
		p, er := check(words[1])
		if p == -1 {
			return -1, er
		}
		n--
		if p != 100 {
			return -1, "После первой единицы могут быть только сотни!"
		}
		n = len(words) - 1
		ans = 100 * ans

		a, er := cent(p, words[1:], n)
		if a == -1 {
			return -1, er
		}
		ans += a - 100

		return ans, ""
	}
	return 0, ""
}
