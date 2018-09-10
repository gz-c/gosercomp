// Code generated by codecgen - DO NOT EDIT.

package gosercomp

import (
	"errors"
	codec1978 "github.com/ugorji/go/codec"
	"runtime"
	"strconv"
)

const (
	// ----- content types ----
	codecSelferCcUTF84480 = 1
	codecSelferCcRAW4480  = 0
	// ----- value types used ----
	codecSelferValueTypeArray4480  = 10
	codecSelferValueTypeMap4480    = 9
	codecSelferValueTypeString4480 = 6
	codecSelferValueTypeInt4480    = 2
	codecSelferValueTypeUint4480   = 3
	codecSelferValueTypeFloat4480  = 4
	codecSelferBitsize4480         = uint8(32 << (^uint(0) >> 63))
)

var (
	errCodecSelferOnlyMapOrArrayEncodeToStruct4480 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer4480 struct{}

func init() {
	if codec1978.GenVersion != 8 {
		_, file, _, _ := runtime.Caller(0)
		panic("codecgen version mismatch: current: 8, need " + strconv.FormatInt(int64(codec1978.GenVersion), 10) + ". Re-generate file: " + file)
	}
	if false { // reference the types, but skip this branch at build/run time
	}
}

func (x *ColorGroup) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer4480
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false // struct tag has 'toArray'
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(3)
			} else {
				r.WriteMapStart(3)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeInt(int64(x.Id))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"id\"")
				} else {
					r.EncodeString(codecSelferCcUTF84480, `id`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeInt(int64(x.Id))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeString(codecSelferCcUTF84480, string(x.Name))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"name\"")
				} else {
					r.EncodeString(codecSelferCcUTF84480, `name`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeString(codecSelferCcUTF84480, string(x.Name))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.Colors == nil {
					r.EncodeNil()
				} else {
					if false {
					} else {
						z.F.EncSliceStringV(x.Colors, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"colors\"")
				} else {
					r.EncodeString(codecSelferCcUTF84480, `colors`)
				}
				r.WriteMapElemValue()
				if x.Colors == nil {
					r.EncodeNil()
				} else {
					if false {
					} else {
						z.F.EncSliceStringV(x.Colors, e)
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayEnd()
			} else {
				r.WriteMapEnd()
			}
		}
	}
}

func (x *ColorGroup) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer4480
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap4480 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray4480 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(errCodecSelferOnlyMapOrArrayEncodeToStruct4480)
		}
	}
}

func (x *ColorGroup) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer4480
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		r.ReadMapElemKey()
		yys3 := z.StringView(r.DecodeStringAsBytes())
		r.ReadMapElemValue()
		switch yys3 {
		case "id":
			if r.TryDecodeAsNil() {
				x.Id = 0
			} else {
				x.Id = (int)(z.C.IntV(r.DecodeInt64(), codecSelferBitsize4480))
			}
		case "name":
			if r.TryDecodeAsNil() {
				x.Name = ""
			} else {
				x.Name = (string)(r.DecodeString())
			}
		case "colors":
			if r.TryDecodeAsNil() {
				x.Colors = nil
			} else {
				if false {
				} else {
					z.F.DecSliceStringX(&x.Colors, d)
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	r.ReadMapEnd()
}

func (x *ColorGroup) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer4480
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj8 int
	var yyb8 bool
	var yyhl8 bool = l >= 0
	yyj8++
	if yyhl8 {
		yyb8 = yyj8 > l
	} else {
		yyb8 = r.CheckBreak()
	}
	if yyb8 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Id = 0
	} else {
		x.Id = (int)(z.C.IntV(r.DecodeInt64(), codecSelferBitsize4480))
	}
	yyj8++
	if yyhl8 {
		yyb8 = yyj8 > l
	} else {
		yyb8 = r.CheckBreak()
	}
	if yyb8 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Name = ""
	} else {
		x.Name = (string)(r.DecodeString())
	}
	yyj8++
	if yyhl8 {
		yyb8 = yyj8 > l
	} else {
		yyb8 = r.CheckBreak()
	}
	if yyb8 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Colors = nil
	} else {
		if false {
		} else {
			z.F.DecSliceStringX(&x.Colors, d)
		}
	}
	for {
		yyj8++
		if yyhl8 {
			yyb8 = yyj8 > l
		} else {
			yyb8 = r.CheckBreak()
		}
		if yyb8 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj8-1, "")
	}
	r.ReadArrayEnd()
}
