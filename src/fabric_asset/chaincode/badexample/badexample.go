package main

import (
	"bytes"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"math/rand"
	"strconv"
	"time"
)

type BadExampleCC struct {
}

//每一个链码必须实现2个方法Init()，Invok()

//链码的初始化
func (c *BadExampleCC) Init(stub shim.ChaincodeStubInterface) pb.Response {

	//直接返回成功
	return shim.Success(nil)
}

//链码交互的入口
func (c *BadExampleCC) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	//直接返回一个随机数结果
	return shim.Success(bytes.NewBufferString(strconv.Itoa(int(rand.Int63n(time.Now().Unix())))).Bytes())
}

func main() {
	err := shim.Start(new(BadExampleCC))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}

}
