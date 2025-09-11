package main

import (
	"fmt"

	"github.com/KOPs-ai/proto/client/go/common"
)

func main() {
	//load env
	common.LoadENV()
	fmt.Println(common.GetEndpoint("micros.base.user"))
}
