package protobufs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yashwant-shrivastava/tech-session/protobufs/proto/customer/customernew"
	"github.com/yashwant-shrivastava/tech-session/protobufs/proto/customer/customerold"
	"google.golang.org/protobuf/proto"
	"testing"
)

func Test_IsProtoBufSame(t *testing.T) {
	c1 := &customernew.Customer{
		CustomerId: "12",
		Name:       "Yashwant",
	}

	c2 := &customerold.Customer{
		CustomerId: "12",
		Name:       "Yashwant",
	}

	//assert.Equal(t, c1, c2)

	marshalledProtoC1, _ := proto.Marshal(c1)
	marshalledProtoC2, _ := proto.Marshal(c2)
	assert.Equal(t, marshalledProtoC1, marshalledProtoC2)
}

func Test_BackwardCompatibility(t *testing.T) {
	c1 := &customernew.Customer{
		CustomerId: "12",
		Name:       "Yashwant",
		Email:      "yashwant@abc.com",
	}

	c2 := &customerold.Customer{}
	marshalledProtoC1, _ := proto.Marshal(c1)
	proto.Unmarshal(marshalledProtoC1, c2)
	fmt.Printf("c1: %+v\nc2: %+v\n", c1, c2)
}

func Test_ForwardCompatibility(t *testing.T) {
	c1 := &customerold.Customer{
		CustomerId: "12",
		Name:       "Yashwant",
	}

	c2 := &customernew.Customer{}
	marshalledProtoC1, _ := proto.Marshal(c1)
	proto.Unmarshal(marshalledProtoC1, c2)
	fmt.Printf("c1: %+v\nc2: %+v\n", c1, c2)
}
