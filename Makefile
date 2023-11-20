
# Set the output directory
OUT_DIR := out

# Set the Go compiler
GO := go

# Set the build flags
BUILD_FLAGS := -v

# Set the source files
SRC := $(wildcard *.go)

# Set the binary name
BIN := productwarehousing

.PHONY: all clean

all: $(OUT_DIR)/$(BIN)

$(OUT_DIR)/$(BIN): $(SRC)
	@mkdir -p $(OUT_DIR)
	$(GO) build $(BUILD_FLAGS) -o $@ $<

clean:
	@rm -rf $(OUT_DIR)

docker:
	docker compose -f docker-compose.yaml up -d