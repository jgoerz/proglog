
Record: the data stored in our log.

Store: the file we store the records in.

Index: the file we store index entries in.

Segment: the abstraction that ties a store and an index together.

Log: the abstratction that ties all the segments together.


Dependencies or tools

- https://grpc.io/docs/protoc-installation/
    - https://github.com/protocolbuffers/protobuf/releases/tag/v21.9
- https://github.com/cloudflare/cfssl
    - go install github.com/cloudflare/cfssl/cmd/cfssl
    - go install github.com/cloudflare/cfssl/cmd/cfssljson

