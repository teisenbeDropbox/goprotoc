// Code generated by protoc-gen-dgo.
// source: example.proto
// DO NOT EDIT!

/*
Package test is a generated protocol buffer package.

It is generated from these files:
	example.proto

It has these top-level messages:
	A
	B
	C
	U
	E
*/
package test

import testing "testing"
import math_rand "math/rand"
import time "time"
import github_com_dropbox_goprotoc_proto "github.com/dropbox/goprotoc/proto"
import testing1 "testing"
import math_rand1 "math/rand"
import time1 "time"
import math_rand2 "math/rand"
import time2 "time"
import testing2 "testing"
import math_rand3 "math/rand"
import time3 "time"
import testing3 "testing"
import fmt1 "fmt"
import math_rand4 "math/rand"
import time4 "time"
import testing4 "testing"
import github_com_dropbox_goprotoc_proto1 "github.com/dropbox/goprotoc/proto"
import testing5 "testing"

func TestAProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedA(popr, false)
	data, err := github_com_dropbox_goprotoc_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &A{}
	if err := github_com_dropbox_goprotoc_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestAMarshalTo(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedA(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &A{}
	if err := github_com_dropbox_goprotoc_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkAProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*A, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedA(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := github_com_dropbox_goprotoc_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkAProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		data, err := github_com_dropbox_goprotoc_proto.Marshal(NewPopulatedA(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	msg := &A{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_dropbox_goprotoc_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestBProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedB(popr, false)
	data, err := github_com_dropbox_goprotoc_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &B{}
	if err := github_com_dropbox_goprotoc_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestBMarshalTo(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedB(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &B{}
	if err := github_com_dropbox_goprotoc_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkBProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*B, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedB(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := github_com_dropbox_goprotoc_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkBProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		data, err := github_com_dropbox_goprotoc_proto.Marshal(NewPopulatedB(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	msg := &B{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_dropbox_goprotoc_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestCProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedC(popr, false)
	data, err := github_com_dropbox_goprotoc_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &C{}
	if err := github_com_dropbox_goprotoc_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestCMarshalTo(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedC(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &C{}
	if err := github_com_dropbox_goprotoc_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkCProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*C, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedC(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := github_com_dropbox_goprotoc_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkCProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		data, err := github_com_dropbox_goprotoc_proto.Marshal(NewPopulatedC(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	msg := &C{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_dropbox_goprotoc_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestUProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedU(popr, false)
	data, err := github_com_dropbox_goprotoc_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &U{}
	if err := github_com_dropbox_goprotoc_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestUMarshalTo(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedU(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &U{}
	if err := github_com_dropbox_goprotoc_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkUProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*U, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedU(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := github_com_dropbox_goprotoc_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkUProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		data, err := github_com_dropbox_goprotoc_proto.Marshal(NewPopulatedU(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	msg := &U{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_dropbox_goprotoc_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestEProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedE(popr, false)
	data, err := github_com_dropbox_goprotoc_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &E{}
	if err := github_com_dropbox_goprotoc_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestEMarshalTo(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedE(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &E{}
	if err := github_com_dropbox_goprotoc_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkEProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*E, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedE(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := github_com_dropbox_goprotoc_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkEProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		data, err := github_com_dropbox_goprotoc_proto.Marshal(NewPopulatedE(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	msg := &E{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_dropbox_goprotoc_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestAAPI(t *testing1.T) {
	popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
	p := NewPopulatedA(popr, false)
	msg := &A{}
	if !apiEmptyA(msg, t) {
		t.Fatalf("A should be empty")
	}
	apiCopyA(msg, p, t)
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
	if apiEmptyA(p, t) != apiEmptyA(msg, t) {
		t.Fatalf("A should not be empty")
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
	msg.Clear()
	if !apiEmptyA(msg, t) {
		t.Fatalf("A should be empty")
	}
}

func apiCopyA(dst *A, src *A, t *testing1.T) {
	if dst == nil || src == nil {
		t.Fatalf("Cannot copy to(%v) or from(%v) nil message", dst, src)
	}
	if src.HasDescription() {
		dst.SetDescription(src.GetDescription())
	}
	if src.HasNumber() {
		dst.SetNumber(src.GetNumber())
	}
	t.Skip("Cannot copy costum field")
	src.XXX_unrecognized = dst.XXX_unrecognized
}

func apiEmptyA(msg *A, t *testing1.T) bool {
	if msg == nil {
		return true
	}
	if msg.HasDescription() {
		return false
	}
	if msg.HasNumber() {
		return false
	}
	t.Skip("Cannot check costum field")
	return true
}

func TestBAPI(t *testing1.T) {
	popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
	p := NewPopulatedB(popr, false)
	msg := &B{}
	if !apiEmptyB(msg, t) {
		t.Fatalf("B should be empty")
	}
	apiCopyB(msg, p, t)
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
	if apiEmptyB(p, t) != apiEmptyB(msg, t) {
		t.Fatalf("B should not be empty")
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
	msg.Clear()
	if !apiEmptyB(msg, t) {
		t.Fatalf("B should be empty")
	}
}

func apiCopyB(dst *B, src *B, t *testing1.T) {
	if dst == nil || src == nil {
		t.Fatalf("Cannot copy to(%v) or from(%v) nil message", dst, src)
	}
	t.Skip("Cannot copy embed field")
	src.XXX_unrecognized = dst.XXX_unrecognized
}

func apiEmptyB(msg *B, t *testing1.T) bool {
	if msg == nil {
		return true
	}
	t.Skip("Cannot check embed field")
	return true
}

func TestCAPI(t *testing1.T) {
	popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
	p := NewPopulatedC(popr, false)
	msg := &C{}
	if !apiEmptyC(msg, t) {
		t.Fatalf("C should be empty")
	}
	apiCopyC(msg, p, t)
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
	if apiEmptyC(p, t) != apiEmptyC(msg, t) {
		t.Fatalf("C should not be empty")
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
	msg.Clear()
	if !apiEmptyC(msg, t) {
		t.Fatalf("C should be empty")
	}
}

func apiCopyC(dst *C, src *C, t *testing1.T) {
	if dst == nil || src == nil {
		t.Fatalf("Cannot copy to(%v) or from(%v) nil message", dst, src)
	}
	if src.HasMySize() {
		dst.SetMySize(src.GetMySize())
	}
	src.XXX_unrecognized = dst.XXX_unrecognized
}

func apiEmptyC(msg *C, t *testing1.T) bool {
	if msg == nil {
		return true
	}
	if msg.HasMySize() {
		return false
	}
	return true
}

func TestUAPI(t *testing1.T) {
	popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
	p := NewPopulatedU(popr, false)
	msg := &U{}
	if !apiEmptyU(msg, t) {
		t.Fatalf("U should be empty")
	}
	apiCopyU(msg, p, t)
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
	if apiEmptyU(p, t) != apiEmptyU(msg, t) {
		t.Fatalf("U should not be empty")
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
	msg.Clear()
	if !apiEmptyU(msg, t) {
		t.Fatalf("U should be empty")
	}
}

func apiCopyU(dst *U, src *U, t *testing1.T) {
	if dst == nil || src == nil {
		t.Fatalf("Cannot copy to(%v) or from(%v) nil message", dst, src)
	}
	if src.HasA() {
		srcA := src.GetA()
		dstA, _ := dst.MutateA()
		apiCopyA(dstA, srcA, t)
	}
	if src.HasB() {
		srcB := src.GetB()
		dstB, _ := dst.MutateB()
		apiCopyB(dstB, srcB, t)
	}
	src.XXX_unrecognized = dst.XXX_unrecognized
}

func apiEmptyU(msg *U, t *testing1.T) bool {
	if msg == nil {
		return true
	}
	if msg.HasA() {
		return false
	}
	if msg.HasB() {
		return false
	}
	return true
}

func TestEAPI(t *testing1.T) {
	popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
	p := NewPopulatedE(popr, false)
	msg := &E{}
	if !apiEmptyE(msg, t) {
		t.Fatalf("E should be empty")
	}
	apiCopyE(msg, p, t)
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
	if apiEmptyE(p, t) != apiEmptyE(msg, t) {
		t.Fatalf("E should not be empty")
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
	msg.Clear()
	if !apiEmptyE(msg, t) {
		t.Fatalf("E should be empty")
	}
}

func apiCopyE(dst *E, src *E, t *testing1.T) {
	if dst == nil || src == nil {
		t.Fatalf("Cannot copy to(%v) or from(%v) nil message", dst, src)
	}
	src.XXX_unrecognized = dst.XXX_unrecognized
	src.XXX_extensions = dst.XXX_extensions
}

func apiEmptyE(msg *E, t *testing1.T) bool {
	if msg == nil {
		return true
	}
	return true
}

func TestUOnlyOne(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedU(popr, true)
	v := p.GetValue()
	msg := &U{}
	if !msg.SetValue(v) {
		t.Fatalf("OnlyOne: Could not set Value")
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !OnlyOne Equal %#v", msg, p)
	}
}
func TestAStringer(t *testing3.T) {
	popr := math_rand3.New(math_rand3.NewSource(time3.Now().UnixNano()))
	p := NewPopulatedA(popr, false)
	s1 := p.String()
	s2 := fmt1.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestBStringer(t *testing3.T) {
	popr := math_rand3.New(math_rand3.NewSource(time3.Now().UnixNano()))
	p := NewPopulatedB(popr, false)
	s1 := p.String()
	s2 := fmt1.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestCStringer(t *testing3.T) {
	popr := math_rand3.New(math_rand3.NewSource(time3.Now().UnixNano()))
	p := NewPopulatedC(popr, false)
	s1 := p.String()
	s2 := fmt1.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestUStringer(t *testing3.T) {
	popr := math_rand3.New(math_rand3.NewSource(time3.Now().UnixNano()))
	p := NewPopulatedU(popr, false)
	s1 := p.String()
	s2 := fmt1.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestEStringer(t *testing3.T) {
	popr := math_rand3.New(math_rand3.NewSource(time3.Now().UnixNano()))
	p := NewPopulatedE(popr, false)
	s1 := p.String()
	s2 := fmt1.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestAVerboseEqual(t *testing4.T) {
	popr := math_rand4.New(math_rand4.NewSource(time4.Now().UnixNano()))
	p := NewPopulatedA(popr, false)
	data, err := github_com_dropbox_goprotoc_proto1.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &A{}
	if err := github_com_dropbox_goprotoc_proto1.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestBVerboseEqual(t *testing4.T) {
	popr := math_rand4.New(math_rand4.NewSource(time4.Now().UnixNano()))
	p := NewPopulatedB(popr, false)
	data, err := github_com_dropbox_goprotoc_proto1.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &B{}
	if err := github_com_dropbox_goprotoc_proto1.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestCVerboseEqual(t *testing4.T) {
	popr := math_rand4.New(math_rand4.NewSource(time4.Now().UnixNano()))
	p := NewPopulatedC(popr, false)
	data, err := github_com_dropbox_goprotoc_proto1.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &C{}
	if err := github_com_dropbox_goprotoc_proto1.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestUVerboseEqual(t *testing4.T) {
	popr := math_rand4.New(math_rand4.NewSource(time4.Now().UnixNano()))
	p := NewPopulatedU(popr, false)
	data, err := github_com_dropbox_goprotoc_proto1.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &U{}
	if err := github_com_dropbox_goprotoc_proto1.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestEVerboseEqual(t *testing4.T) {
	popr := math_rand4.New(math_rand4.NewSource(time4.Now().UnixNano()))
	p := NewPopulatedE(popr, false)
	data, err := github_com_dropbox_goprotoc_proto1.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &E{}
	if err := github_com_dropbox_goprotoc_proto1.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestExampleDescription(t *testing5.T) {
	ExampleDescription()
}

//These tests are generated by github.com/dropbox/goprotoc/plugin/testgen
