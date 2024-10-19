package gen

import (
	"sync/atomic"
)

var id atomic.Uint64

type Inc int

type Response struct {
	Id uint64
}

type Request struct{}

// Can't have a function with this signature
// func (i *Inc) Next() uint64 {
// 	return id.Add(1)
// }

// Required signature,

func (i *Inc) Next(_ Request, response *Response) error {
	response.Id = id.Add(1)

	return nil
}
