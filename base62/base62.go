package base62

var (
	Inst = NewBase62() // 单例
)

// Base62 base62编码
type Base62 struct {
	table []byte // base62表格
}

// Encode base62编码
func (object *Base62) Encode(input uint32) (output []byte) {
	indexes := make([]uint32, 0)
	for 0 < input {
		indexes = append(indexes, input%uint32(len(object.table)))
		input /= uint32(len(object.table))
	}
	output = make([]byte, 0, len(indexes))
	for i := len(indexes) - 1; i >= 0; i-- {
		output = append(output, object.table[indexes[i]])
	}
	return
}

// Decode Base62解码
func (object *Base62) Decode(input []byte) (output uint32) {
	v := uint32(1)
	for i := len(input) - 1; i >= 0; i-- {
		for j, e := range object.table {
			if e == input[i] {
				output = output + uint32(j)*v
				v *= 62
				break
			}
		}
	}
	return
}

// NewBase62 工厂方法
func NewBase62() *Base62 {
	return &Base62{
		table: []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"),
	}
}
