# CLIFIG
## Install Instructions:

### 1. Shell (Mac/Linux)
Install the latest version:
```
curl -fsSL https://github.com/StatixLabs/clifig/raw/main/install.sh | bash
```

### Testing:

Run Test:
```
go test ./...
```

Setting up ginkgo for a package:
```
ginkgo bootstrap
```

Generate tests:
```
ginkgo generate <name of file>
```

### Build:
using [goreleaser](https://goreleaser.com/) for build.
```
brew install goreleaser/tap/goreleaser
```
