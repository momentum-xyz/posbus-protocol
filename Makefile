# linux   use "bin/flatc"
# windows use "bin/flatc.exe"
gen-go:
	rm -rf flatbuff/go/api && flatbuff/bin/flatc_mac --go-namespace api --go -o "flatbuff/go/api" flatbuff/API.fbs --gen-onefile

gen-csharp:
	rm -rf flatbuff/csharp/api && flatbuff/bin/flatc_mac --csharp -o "flatbuff/csharp" flatbuff/API.fbs --gen-onefile

gen-rust:
	rm -rf flatbuff/rust/api && flatbuff/bin/flatc_mac --rust -o "flatbuff/rust" flatbuff/API.fbs --gen-onefile

.PHONY: gen-go gen-csharp gen-rust
