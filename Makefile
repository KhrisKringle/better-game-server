

backend:
	protoc --plugin=protoc-gen-ts_proto=./node_modules/.bin/protoc-gen-ts_proto \
    --ts_proto_out=frontend/src websocket.proto

frontend:
	protoc --plugin=protoc-gen-ts_proto=/opt/homebrew/bin/protoc-gen-ts_proto \
       --ts_proto_out=frontend/src \
       websocket.proto

all: backend frontend