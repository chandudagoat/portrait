generate:
	python -m grpc_tools.protoc --proto_path=proto proto/*.proto --python_out=pb/ --grpc_python_out=pb/