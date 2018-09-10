## Golang Serialization Benchmark

### Serializers

This project test the below go serializers, which compares with go standard _json_ and _xml_.

- [skycoin/src/cipher/encoder](https://github.com/skycoin/skycoin/tree/develop/src/cipher/encoder)
- [encoding/json](http://golang.org/pkg/encoding/json/)
- [encoding/xml](http://golang.org/pkg/encoding/xml/)
- [github.com/tinylib/msgp](http://github.com/tinylib/msgp)
- [github.com/golang/protobuf](http://github.com/golang/protobuf)
- [github.com/gogo/protobuf](http://github.com/gogo/protobuf)
- [github.com/google/flatbuffers](http://github.com/google/flatbuffers)
- [Apache/Thrift](https://github.com/apache/thrift/tree/master/lib/go)
- [Apache/Avro](https://github.com/linkedin/goavro)
- [andyleap/gencode](https://github.com/andyleap/gencode)
- [ugorji/go/codec](https://github.com/ugorji/go/tree/master/codec)
- [go-memdump](https://github.com/alexflint/go-memdump)
- [colfer](https://github.com/pascaldekloe/colfer)
- [zebrapack](https://github.com/glycerine/zebrapack)
- [gotiny](https://github.com/niubaoshu/gotiny)
- [github.com/ugorji/go/codec](http://github.com/ugorji/go/codec)
- [hprose-golang](https://github.com/hprose/hprose-golang/tree/master/io)
- [vmihailenco/msgpack.v2](https://github.com/vmihailenco/msgpack)
- [Sereal](https://github.com/Sereal/Sereal)
- [ffjson](https://github.com/pquerna/ffjson)
- [easyjson](https://github.com/mailru/easyjson)
- [jsoniter](https://github.com/json-iterator/go)

###  Excluded Serializers

Given existed [benchmark](https://github.com/alecthomas/go_serialization_benchmarks) by alecthomas，the below serializers are excluded from this test because of their poor performance.

- [encoding/gob](http://golang.org/pkg/encoding/gob/)
- [github.com/alecthomas/binary](http://github.com/alecthomas/binary)
- [github.com/davecgh/go-xdr/xdr](http://github.com/davecgh/go-xdr/xdr)
- [labix.org/v2/mgo/bson](http://labix.org/v2/mgo/bson)
- [github.com/DeDiS/protobuf](http://github.com/DeDiS/protobuf)
- [gopkg.in/vmihailenco/msgpack.v2](http://gopkg.in/vmihailenco/msgpack.v2)
- [bson](http://github.com/micro/go-bson)

### Test Environment
go version: **1.11**

You'll need to install some tools.  On OS X:

```sh
brew install thrift flatbuffers
```


- For `MessagePack`，you need install the tool and use `go generate` to generate code:

  ```go
  go get github.com/tinylib/msgp
  go generate
  ```

- For `ProtoBuf`, you need to install [protoc](https://github.com/google/protobuf/releases)，protobuf lib and generate code：

  ```go
  go get github.com/golang/protobuf
  go generate
  ```

- For `gogo/protobuf`, use the below commands：

  ```go
  go get github.com/gogo/protobuf/gogoproto
  go get -u github.com/gogo/protobuf/protoc-gen-gogofaster
  go generate
  ```

- For `flatbuffers`, you need to install [flatbuffers compiler](https://github.com/google/flatbuffers/releases,  and flatbuffers lib：

  ```go
  go get github.com/google/flatbuffers/go
  go generate
  ```

- For `thrift`, you need to install [thrift compiler](https://thrift.apache.org/download), and thrift lib：

  ```go
  go get git.apache.org/thrift.git/lib/go/thrift
  go generate
  ```

- For `Avro`, you need to install goavro and checkout `v1.0.5`：

    ```go
    go get github.com/linkedin/goavro
    cd $GOPATH/src/github.com/linkedin/goavro
    git checkout v1.0.5
    cd -
    go generate
    ```

- For `gencode`, you need to install gencode, and geneate code by gencode：

  ```go
  go get github.com/andyleap/gencode
  gencode go -schema=gencode.schema -package gosercomp
  ```


- For `easyjson`, you need to install easyjson:

  ```go
  go get github.com/mailru/easyjson
  go generate
  ```

- For `zebraPack `, you need to install zebraPack, and generate code：

  ```go
  go get github.com/glycerine/zebrapack
  go generate zebrapack_data.go
  ```

- For `ugorji/go/codec` you need to install codecgen and `codec` lib:

```go
  go get -tags=unsafe  -u github.com/ugorji/go/codec/codecgen
  go get -tags=unsafe -u github.com/ugorji/go/codec

  codecgen -o data_codec.go data.go
```

- For `niubaoshu/gotiny` you need to checkout `v0.0.2`:

```go
	go get github.com/niubaoshu/gotiny
	cd $GOPATH/src/github.com/niubaoshu/gotiny
	git checkout v0.0.2
```

`ugorji/go/codec` supports msgpack、cbor、binc、json, and this project test its  cbor and msgpack.

> Actually，you can use `go generate` to generate code.

**Test:**

```
go test -bench=. -benchmem
```

### Test Data Model

All tests are using the same data model as below:

```go
type ColorGroup struct {
    ID     int `json:"id" xml:"id,attr""`
    Name   string `json:"name" xml:"name"`
    Colors []string `json:"colors" xml:"colors"`
}
`
```

### Benchmark

Results generated on Mid-2015 Macbook Pro base model:

```
BenchmarkMarshalByJson-8                       	 2000000	       688 ns/op	     128 B/op	       2 allocs/op
BenchmarkUnmarshalByJson-8                     	  500000	      2212 ns/op	     216 B/op	       8 allocs/op
BenchmarkMarshalByXml-8                        	  500000	      3726 ns/op	    4800 B/op	      11 allocs/op
BenchmarkUnmarshalByXml-8                      	  100000	     13710 ns/op	    3027 B/op	      73 allocs/op
BenchmarkMarshalByMsgp-8                       	20000000	       122 ns/op	      80 B/op	       1 allocs/op
BenchmarkUnmarshalByMsgp-8                     	10000000	       222 ns/op	      32 B/op	       5 allocs/op
BenchmarkMarshalByProtoBuf-8                   	10000000	       233 ns/op	      48 B/op	       1 allocs/op
BenchmarkUnmarshalByProtoBuf-8                 	 3000000	       606 ns/op	     160 B/op	      10 allocs/op
BenchmarkMarshalByGogoProtoBuf-8               	10000000	       132 ns/op	      48 B/op	       1 allocs/op
BenchmarkUnmarshalByGogoProtoBuf-8             	 3000000	       387 ns/op	     144 B/op	       8 allocs/op
BenchmarkMarshalByFlatBuffers-8                	 5000000	       361 ns/op	      16 B/op	       1 allocs/op
BenchmarkUnmarshalByFlatBuffers-8              	2000000000	         0.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnmarshalByFlatBuffers_withFields-8   	10000000	       126 ns/op	       0 B/op	       0 allocs/op
BenchmarkMarshalByThrift-8                     	 3000000	       446 ns/op	      64 B/op	       1 allocs/op
BenchmarkUnmarshalByThrift-8                   	 1000000	      1300 ns/op	     656 B/op	      11 allocs/op
BenchmarkMarshalByAvro-8                       	 3000000	       541 ns/op	      48 B/op	       6 allocs/op
BenchmarkUnmarshalByAvro-8                     	  500000	      3257 ns/op	    1672 B/op	      62 allocs/op
BenchmarkMarshalByGencode-8                    	30000000	        44.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnmarshalByGencode-8                  	10000000	       121 ns/op	      32 B/op	       5 allocs/op
BenchmarkMarshalByUgorjiCodecAndCbor-8         	 2000000	       727 ns/op	     112 B/op	       3 allocs/op
BenchmarkUnmarshalByUgorjiCodecAndCbor-8       	 2000000	       673 ns/op	      48 B/op	       6 allocs/op
BenchmarkMarshalByUgorjiCodecAndMsgp-8         	 2000000	       692 ns/op	     112 B/op	       3 allocs/op
BenchmarkUnmarshalByUgorjiCodecAndMsgp-8       	 2000000	       722 ns/op	      48 B/op	       6 allocs/op
BenchmarkMarshalByUgorjiCodecAndBinc-8         	 2000000	       807 ns/op	     112 B/op	       3 allocs/op
BenchmarkUnmarshalByUgorjiCodecAndBinc-8       	 1000000	      1107 ns/op	     824 B/op	      10 allocs/op
BenchmarkMarshalByUgorjiCodecAndJson-8         	 2000000	      1046 ns/op	     112 B/op	       3 allocs/op
BenchmarkUnmarshalByUgorjiCodecAndJson-8       	 2000000	       813 ns/op	      48 B/op	       6 allocs/op
BenchmarkMarshalByEasyjson-8                   	 5000000	       337 ns/op	     128 B/op	       1 allocs/op
BenchmarkUnmarshalByEasyjson-8                 	 3000000	       535 ns/op	      32 B/op	       5 allocs/op
BenchmarkMarshalByFfjson-8                     	 1000000	      1064 ns/op	     424 B/op	       9 allocs/op
BenchmarkUnmarshalByFfjson-8                   	 1000000	      1535 ns/op	     480 B/op	      13 allocs/op
BenchmarkMarshalByJsoniter-8                   	 2000000	       659 ns/op	      96 B/op	       2 allocs/op
BenchmarkUnmarshalByJsoniter-8                 	 2000000	       647 ns/op	      32 B/op	       5 allocs/op
BenchmarkUnmarshalByGJSON-8                    	 1000000	      1885 ns/op	     624 B/op	       7 allocs/op
BenchmarkMarshalByGoMemdump-8                  	  300000	      5328 ns/op	    1032 B/op	      30 allocs/op
BenchmarkUnmarshalByGoMemdump-8                	 1000000	      1534 ns/op	    2400 B/op	      12 allocs/op
BenchmarkMarshalBySky-8                        	 1000000	      1140 ns/op	     208 B/op	      13 allocs/op
BenchmarkUnmarshalBySky-8                      	  500000	      3175 ns/op	     883 B/op	      31 allocs/op
BenchmarkMarshalByColfer-8                     	50000000	        33.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnmarshalByColfer-8                   	 5000000	       212 ns/op	      96 B/op	       6 allocs/op
BenchmarkMarshalByZebrapack-8                  	10000000	       240 ns/op	     182 B/op	       0 allocs/op
BenchmarkUnmarshalByZebrapack-8                	 5000000	       303 ns/op	      32 B/op	       5 allocs/op
BenchmarkMarshalByGotiny-8                     	 3000000	       465 ns/op	     184 B/op	       6 allocs/op
BenchmarkUnmarshalByGotiny-8                   	 5000000	       261 ns/op	      88 B/op	       2 allocs/op
BenchmarkMarshalByHprose-8                     	 3000000	       462 ns/op	     210 B/op	       1 allocs/op
BenchmarkUnmarshalByHprose-8                   	 2000000	       643 ns/op	     288 B/op	       9 allocs/op
BenchmarkMarshalBySereal-8                     	 1000000	      2288 ns/op	     792 B/op	      22 allocs/op
BenchmarkUnmarshalBySereal-8                   	 2000000	       755 ns/op	      80 B/op	       6 allocs/op
BenchmarkMarshalByMsgpackV2-8                  	 1000000	      2124 ns/op	     192 B/op	       4 allocs/op
BenchmarkUnmarshalByMsgpackv2-8                	 1000000	      1657 ns/op	     232 B/op	      11 allocs/op
BenchmarkMarshalMsgZColorGroup-8               	20000000	        81.4 ns/op	      80 B/op	       1 allocs/op
BenchmarkAppendMsgZColorGroup-8                	30000000	        44.7 ns/op	 313.29 MB/s	       0 B/op	       0 allocs/op
BenchmarkUnmarshalZColorGroup-8                	20000000	        86.9 ns/op	 161.02 MB/s	       0 B/op	       0 allocs/op
BenchmarkEncodeZColorGroup-8                   	20000000	        75.8 ns/op	 184.78 MB/s	      16 B/op	       1 allocs/op
BenchmarkDecodeZColorGroup-8                   	10000000	       136 ns/op	 102.57 MB/s	       0 B/op	       0 allocs/op
```


## Size of marshalled results


```
	gosercomp_test.go:91: json:				 65 bytes
	gosercomp_test.go:94: xml:				 137 bytes
	gosercomp_test.go:97: msgp:				 47 bytes
	gosercomp_test.go:100: protobuf:				 36 bytes
	gosercomp_test.go:103: gogoprotobuf:			 36 bytes
	gosercomp_test.go:107: flatbuffers:			 108 bytes
	gosercomp_test.go:113: thrift:				 63 bytes
	gosercomp_test.go:127: avro:				 32 bytes
	gosercomp_test.go:136: gencode:				 34 bytes
	gosercomp_test.go:142: UgorjiCodec_Cbor:		 47 bytes
	gosercomp_test.go:148: UgorjiCodec_Msgp:		 47 bytes
	gosercomp_test.go:154: UgorjiCodec_Bin:			 53 bytes
	gosercomp_test.go:156: UgorjiCodec_Json:		 91 bytes
	gosercomp_test.go:159: easyjson:			 65 bytes
	gosercomp_test.go:162: ffjson:				 65 bytes
	gosercomp_test.go:165: jsoniter:			 65 bytes
	gosercomp_test.go:169: memdump:				 200 bytes
	gosercomp_test.go:172: colfer:				 35 bytes
	gosercomp_test.go:175: zebrapack:			 35 bytes
	gosercomp_test.go:178: gotiny:				 32 bytes
	gosercomp_test.go:183: hprose:				 32 bytes
	gosercomp_test.go:187: sereal:				 76 bytes
	gosercomp_test.go:190: msgpackv2:			 47 bytes
```
