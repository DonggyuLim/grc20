package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type BalanceRequest struct {
	TokenName string `json:"tokenName"`
	Account   string `json:"account"`
}

func BalanceOf(c *gin.Context) {
	r := BalanceRequest{}
	err := c.ShouldBindJSON(&r)
	fmt.Println("r==========", r)
	if err != nil {
		c.String(400, err.Error())
	}
	t := GetToken(c, r.TokenName)

	balance := t.BalanceOf(r.Account)

	c.JSON(200, gin.H{
		"account": r.Account,
		"balance": balance,
	})
}

// 토큰 정보
func TokenInfo(c *gin.Context) {
	tn := c.Param("name")
	fmt.Println(tn)
	t := GetToken(c, tn)

	c.JSON(200, gin.H{
		"tokenName":   t.GetName(),
		"symbol":      t.GetSymbol(),
		"decimal":     t.GetDecimal(),
		"totalSupply": t.GetTotalSupply(),
	})
}

type allowanceRequest struct {
	TokenName string `json:"tokenName"`
	Owner     string `json:"owner"`
	Spender   string `json:"spender"`
}

func Allowance(c *gin.Context) {
	r := allowanceRequest{}
	err := c.ShouldBindJSON(&r)
	fmt.Println("r==========", r)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	t := GetToken(c, r.TokenName)
	t.Allowance(r.Owner, r.Spender)

	c.JSON(200, gin.H{
		"allowance": t.Allowance(r.Owner, r.Spender),
	})
}
