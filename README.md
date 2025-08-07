# changefeedpb

Protobuf definitions for CockroachDB changefeeds.

This package defines the protocol buffer schemas used to encode changefeed messages for CockroachDB's CDC system.


If any changes are made to changefeed.proto, it needs to be regenerated using the protoc compiler with the gogoroach plugin. To do this follow the steps: 

1. First make sure you have the protoc-gen-gogoroach plugin installed and available in your PATH.
2. Then run the following command from the root of the repository:
```
ERR_DIR=$(go list -m -f '{{.Dir}}' github.com/cockroachdb/errors)/errorspb GOGO_TYPES=$(go list -m -f '{{.Dir}}' github.com/gogo/protobuf)/typesGOGO_DESC=$(go list -m -f '{{.Dir}}' github.com/gogo/protobuf)/protoc-gen-gogo/descriptor

protoc -I . \
  -I "$ERR_DIR" \
  -I "$GOGO_TYPES" \
  -I "$GOGO_DESC" \                                        
  --plugin=protoc-gen-gogoroach=$(which protoc-gen-gogoroach) \
  --gogoroach_out=paths=source_relative,plugins=grpc,\        
Merrorspb/errors.proto=github.com/cockroachdb/errors/errorspb,\                     
Mgoogle/api/annotations.proto=google.golang.org/genproto/googleapis/api/annotations,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor,\               
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types:. \
  changefeed.proto
```


