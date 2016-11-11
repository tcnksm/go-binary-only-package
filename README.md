# Sample Binary-Only Packages

Go1.7 introduces [Binary-Only Packages](https://github.com/golang/proposal/blob/master/design/2775-binary-only-packages.md). This is sample project (named `hello`) to create a Binary-Only Pacakage and distribute it as `zip` file. 
## How to create?

Actual source code is `hello.go` in the root directory of this repository. We will not distribute this code.

To create Binary-Only packages, we need prepare 2 files. The one is a special source code file and the other is package binary.

This special source code file is described in doc,

> the package must be distributed with a source file not excluded by build constraints and containing a "//go:binary-only-package" comment.

[https://tip.golang.org/pkg/go/build/#hdr-Binary_Only_Packages](https://tip.golang.org/pkg/go/build/#hdr-Binary_Only_Packages)

The example of this source code file is in [src/github.com/tcnksm/hello/](src/github.com/tcnksm/hello/).

Then, build pacakge binary and place it in `pkg/darwin_amd64/github.com/tcnksm/` directory.

```bash
$ go build -o pkg/darwin_amd64/github.com/tcnksm/hello.a -x
```

Finally, zip `src` and `pkg` directory,

```bash
$ zip -r hello.zip src/* pkg/*
```

## How to distribute?

Share above `zip` file which doesn't have actual source code. 

To use this binary pacakge, just unzip it in `$GOPATH`,

```bash
$ unzip hello.zip -d $GOPATH/
```

## How to use?
Please have a look at binaryPackageUsageExample.go, just use the binary package as a real package. 

But you should make sure that the .a file in $GOPATH/pkg directory's struct is the same as your fake minimal binary package in $GOPATH/src.

In the example here, pkg directory is

```
pkg/darwin_amd64/github.com/tcnksm
```

while src directory is

```
src/darwin_amd64/github.com/tcnksm
```
