
all: proto, wire

.PHONY: proto, wire

proto:
	@./build_proto.sh

wire:
	@./build_wire.sh