package mecab

// #cgo CFLAGS: -I/usr/local/include
// #cgo LDFLAGS: -L/usr/local/lib -lmecab -lstdc++
// #include <mecab.h>
import "C"

type model struct {
	m *C.mecab_model_t
}

type tagger struct {
	t *C.mecab_t
}

type Node struct {
	Surface  string
	Feature  []string
	Length   uint32
	RLength  uint32
	Id       uint32
	RcAttr   uint16
	LcAttr   uint16
	PosID    uint16
	CharType uint8
	Stat     uint8
	IsBest   uint8
	Alpha    float32
	Beta     float32
	WCost    int16
	Cost     int64
}

func NewModel(arg string) (res *model) {
	m := C.mecab_model_new2(C.CString(arg))
	if m != nil {
		res = &model{m: m}
	}
	return
}

func (m *model) NewTagger() (res *tagger) {
	t := C.mecab_model_new_tagger(m.m)
	if t != nil {
		res = &tagger{t: t}
	}
	return
}

func (m *model) Close() {
	C.mecab_model_destroy(m.m)
}

func (t *tagger) Close() {
	C.mecab_destroy(t.t)
}

func (t *tagger) Parse(str string) string {
	return t.ParseToString(str)
}

func (t *tagger) ParseToString(str string) string {
	res := C.mecab_sparse_tostr(t.t, C.CString(str))
	return C.GoString(res)
}

// TODO: func (t *tagger) ParseToNode(str string) (res []Node)

// TODO: func (t *tagger) ParseNBest(str string) (res []string)
