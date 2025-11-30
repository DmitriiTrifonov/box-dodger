define WEBGAME
<!DOCTYPE html>
<script src="wasm_exec.js"></script>
<script>
const go = new Go();
WebAssembly.instantiateStreaming(fetch("box-dodger.wasm"), go.importObject).then(result => {
    go.run(result.instance);
});
</script>
endef


.PHONY: build build-web

bin/:
	mkdir bin

build: bin/
	go build -o bin/box-dodger cmd/box-dodger/main.go

build-web: wasm_exec.js
	GOOS=js GOARCH=wasm go build -o bin/box-dodger.wasm cmd/box-dodger/main.go

wasm_exec.js: bin/
	cp $(GOROOT)/lib/wasm/wasm_exec.js bin/

index.html:
	@echo "$(WEBGAME)" > index.html