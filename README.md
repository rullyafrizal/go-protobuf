## GolangProtocol Buffer

### Compile Protocol Buffer
- Compile Protobuf with relative directory (same directory with the .proto file)
```
protoc --go_out={output_path} --go_opt=paths=source_relative {proto_file_path}
```

- For example : 
```
protoc --go_out=. --go_opt=paths=source_relative ./user/*.proto
(it will generate whatever.pb.go in the same directory with the .proto file)
```