package ipcbench

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/golang/protobuf/proto"
)

var nestedRes *Top

func randSeq(n int, r *rand.Rand) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

func generateNested(r *rand.Rand) *Nested {
	scalars := make([]*Scalars, r.Int31n(15))
	for i := range scalars {
		scalars[i] = &Scalars{
			D: r.Float64(),
			I: r.Int63(),
			B: r.Intn(2) == 1,
			S: randSeq(r.Intn(20), r),
		}
	}

	return &Nested{
		Id:   r.Int31(),
		Name: randSeq(r.Intn(20), r),
		Subtype: &Nested_SubType{
			Name: randSeq(r.Intn(20), r),
			Type: Nested_Type(r.Int31n(4)),
		},
		Scalars: scalars,
	}
}

func generateTop(n int, r *rand.Rand) *Top {
	nested := make([]*Nested, n)
	for i := range nested {
		nested[i] = generateNested(r)
	}
	return &Top{Nested: nested}
}

/*
 * MARSHAL
 */

func benchProtoMarshalNested(size int, b *testing.B) {
	var res []byte

	r := rand.New(rand.NewSource(42))
	top := generateTop(size, r)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		res, _ = proto.Marshal(top)
	}
	marshalRes = res
}

func BenchmarkNestedProtoMarshal1(b *testing.B)     { benchProtoMarshalNested(1, b) }
func BenchmarkNestedProtoMarshal10(b *testing.B)    { benchProtoMarshalNested(10, b) }
func BenchmarkNestedProtoMarshal100(b *testing.B)   { benchProtoMarshalNested(100, b) }
func BenchmarkNestedProtoMarshal1000(b *testing.B)  { benchProtoMarshalNested(1000, b) }
func BenchmarkNestedProtoMarshal5000(b *testing.B)  { benchProtoMarshalNested(5000, b) }
func BenchmarkNestedProtoMarshal10000(b *testing.B) { benchProtoMarshalNested(10000, b) }

func benchJSONMarshalNested(size int, b *testing.B) {
	var res []byte

	r := rand.New(rand.NewSource(42))
	top := generateTop(size, r)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		res, _ = json.Marshal(top)
	}
	marshalRes = res
}

func BenchmarkNestedJSONMarshal1(b *testing.B)     { benchJSONMarshalNested(1, b) }
func BenchmarkNestedJSONMarshal10(b *testing.B)    { benchJSONMarshalNested(10, b) }
func BenchmarkNestedJSONMarshal100(b *testing.B)   { benchJSONMarshalNested(100, b) }
func BenchmarkNestedJSONMarshal1000(b *testing.B)  { benchJSONMarshalNested(1000, b) }
func BenchmarkNestedJSONMarshal5000(b *testing.B)  { benchJSONMarshalNested(5000, b) }
func BenchmarkNestedJSONMarshal10000(b *testing.B) { benchJSONMarshalNested(10000, b) }

/*
 * UNMARSHAL
 */

func benchProtoUnMarshalNested(size int, b *testing.B) {
	r := rand.New(rand.NewSource(42))
	top := generateTop(size, r)
	buf, _ := proto.Marshal(top)

	b.ResetTimer()

	var res Top
	for n := 0; n < b.N; n++ {
		proto.Unmarshal(buf, &res)
	}
	nestedRes = &res
}

func BenchmarkNestedProtoUnMarshal1(b *testing.B)     { benchProtoUnMarshalNested(1, b) }
func BenchmarkNestedProtoUnMarshal10(b *testing.B)    { benchProtoUnMarshalNested(10, b) }
func BenchmarkNestedProtoUnMarshal100(b *testing.B)   { benchProtoUnMarshalNested(100, b) }
func BenchmarkNestedProtoUnMarshal1000(b *testing.B)  { benchProtoUnMarshalNested(1000, b) }
func BenchmarkNestedProtoUnMarshal5000(b *testing.B)  { benchProtoUnMarshalNested(5000, b) }
func BenchmarkNestedProtoUnMarshal10000(b *testing.B) { benchProtoUnMarshalNested(10000, b) }

func benchJSONUnMarshalNested(size int, b *testing.B) {
	r := rand.New(rand.NewSource(42))
	top := generateTop(size, r)
	buf, _ := json.Marshal(top)

	b.ResetTimer()

	var res Top
	for n := 0; n < b.N; n++ {
		json.Unmarshal(buf, &res)
	}
	nestedRes = &res
}

func BenchmarkNestedJSONUnMarshal1(b *testing.B)     { benchJSONUnMarshalNested(1, b) }
func BenchmarkNestedJSONUnMarshal10(b *testing.B)    { benchJSONUnMarshalNested(10, b) }
func BenchmarkNestedJSONUnMarshal100(b *testing.B)   { benchJSONUnMarshalNested(100, b) }
func BenchmarkNestedJSONUnMarshal1000(b *testing.B)  { benchJSONUnMarshalNested(1000, b) }
func BenchmarkNestedJSONUnMarshal5000(b *testing.B)  { benchJSONUnMarshalNested(5000, b) }
func BenchmarkNestedJSONUnMarshal10000(b *testing.B) { benchJSONUnMarshalNested(10000, b) }
