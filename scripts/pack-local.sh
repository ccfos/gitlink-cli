#!/bin/bash
# Pack npm package with the binary for the current platform pre-included.
# This way npm install does NOT need to download anything.
set -e

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
NPM_DIR="$PROJECT_DIR/npm"

MODULE="github.com/gitlink-org/gitlink-cli"
VERSION=$(node -p "require('$NPM_DIR/package.json').version")
LDFLAGS="-s -w -X '${MODULE}/cmd.Version=${VERSION}'"
BINARY="gitlink-cli"

echo "=== Building gitlink-cli for current platform ==="

cd "$PROJECT_DIR"

# Build for current platform
CGO_ENABLED=0 go build -ldflags "$LDFLAGS" -o "$NPM_DIR/bin/${BINARY}" .

chmod +x "$NPM_DIR/bin/${BINARY}"
cp "$PROJECT_DIR/README.md" "$NPM_DIR/README.md"

# Copy skills directory
rm -rf "$NPM_DIR/skills"
cp -r "$PROJECT_DIR/skills" "$NPM_DIR/skills"

echo "Binary size: $(du -h "$NPM_DIR/bin/${BINARY}" | cut -f1)"

echo ""
echo "=== Creating npm tarball ==="
cd "$NPM_DIR"
npm pack

echo ""
echo "=== Done ==="
echo "To install locally: npm install -g gitlink-ai-cli-${VERSION}.tgz"
echo "To publish: cd npm && npm publish --access public"
