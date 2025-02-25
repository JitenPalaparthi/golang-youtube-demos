module demo

go 1.24

toolchain go1.24.0

require (
	github.com/JitenPalaparthi/shapes1 v0.0.0-20250225031134-cae8f5fa5b81
	github.com/google/uuid v1.6.0
)

replace github.com/JitenPalaparthi/shapes1 v0.0.0-20250225031134-cae8f5fa5b81 => ../shapes1
