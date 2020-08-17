# qlikapp Entry point

## Directory structure:
```
cmd
│   ├── main.go
│   └── README.md

```

## Overview

The main.go file is an entry point for qlikapp application. It invokes application handler.
```
package main

import (
	"github.com/rahulsidpatil/qlikapp/pkg/handlers"
)

func main() {
	a := handlers.App{}
	a.Initialize()
	a.Run()
}

```
