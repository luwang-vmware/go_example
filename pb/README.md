1. how to generate pb.go
a. cd /Users/luwang/Desktop/GIT/TMC/go_examples;
b. protoc -I=. --go_out=./ ./pb/example.proto
2. how to run the code
a. cd /Users/luwang/Desktop/GIT/TMC/go_examples/pb_examples
b. go run add.go output
c. go run list.go output