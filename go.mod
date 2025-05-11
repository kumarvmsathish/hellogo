module github.com/kumarvmsathish/hellogo

go 1.24.1

// Usually we don't use replace in production, we always keep the functions/library in servers like github
replace github.com/kumarvmsathish/mystrings v0.0.0 => ../mystrings

require github.com/kumarvmsathish/mystrings v0.0.0

require github.com/wagslane/go-tinytime v0.0.2 // indirect -- added using 'go get github.com/wagslane/go-tinytime'

