package ipcbench

import (
	"encoding/json"
	"testing"

	"github.com/golang/protobuf/proto"
)

var (
	// write to this global to avoid otimizing out the results
	marshalRes []byte
	simpleRes  *Simple
)

func BenchmarkSimpleProtoMarshal(b *testing.B) {
	var r []byte
	s := &Simple{
		A: 42,
	}

	for n := 0; n < b.N; n++ {
		r, _ = proto.Marshal(s)
	}
	marshalRes = r
}

func BenchmarkSimpleJSONMarshal(b *testing.B) {
	var r []byte
	s := &Simple{
		A: 42,
	}

	for n := 0; n < b.N; n++ {
		r, _ = json.Marshal(s)
	}
	marshalRes = r
}

func BenchmarkSimpleProtoUnMarshal(b *testing.B) {
	s1 := &Simple{
		A: 42,
	}
	buf, _ := proto.Marshal(s1)

	var r Simple
	for n := 0; n < b.N; n++ {
		proto.Unmarshal(buf, &r)
	}
	simpleRes = &r
}

func BenchmarkSimpleJSONUnMarshal(b *testing.B) {
	s1 := &Simple{
		A: 42,
	}
	buf, _ := json.Marshal(s1)

	var r Simple
	for n := 0; n < b.N; n++ {
		json.Unmarshal(buf, &r)
	}
	simpleRes = &r
}
