// cmd/nftmarketplacebuilderprotocolx/main.go
package main

import (
"flag"
"log"
"os"

"nftmarketplacebuilderprotocolx/internal/nftmarketplacebuilderprotocolx"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := nftmarketplacebuilderprotocolx.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
