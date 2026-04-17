#!/bin/bash
# Build multi-platform binaries and package for npm
set -e

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
NPM_DIR="$PROJECT_DIR/npm"
DIST_DIR="$PROJECT_DIR/dist"

# Read version from npm/package.json
VERSION=$(node -p "require('$NPM_DIR/package.json').version")
MODULE="github.com/gitlink-org/gitlink-cli"
LDFLAGS="-s -w -X '${MODULE}/cmd.Version=${VERSION}'"
BINARY="gitlink-cli"

echo "=== Building gitlink-cli v${VERSION} ==="

# Clean
rm -rf "$DIST_DIR"
mkdir -p "$DIST_DIR"

# Platforms to build: GOOS/GOARCH
PLATFORMS=(
  "darwin/amd64"
  "darwin/arm64"
  "linux/amd64"
  "linux/arm64"
  "windows/amd64"
  "windows/arm64"
)

cd "$PROJECT_DIR"

for PLATFORM in "${PLATFORMS[@]}"; do
  GOOS="${PLATFORM%/*}"
  GOARCH="${PLATFORM#*/}"
  OUTPUT_NAME="${BINARY}"
  if [ "$GOOS" = "windows" ]; then
    OUTPUT_NAME="${BINARY}.exe"
  fi

  echo "Building ${GOOS}/${GOARCH}..."

  ARCHIVE_DIR="${DIST_DIR}/${BINARY}_${VERSION}_${GOOS}_${GOARCH}"
  mkdir -p "$ARCHIVE_DIR"

  CGO_ENABLED=0 GOOS="$GOOS" GOARCH="$GOARCH" \
    go build -ldflags "$LDFLAGS" -o "${ARCHIVE_DIR}/${OUTPUT_NAME}" .

  # Create archive (zip for Windows, tar.gz for others)
  if [ "$GOOS" = "windows" ]; then
    ARCHIVE_NAME="${BINARY}_${VERSION}_${GOOS}_${GOARCH}.zip"
    (cd "${ARCHIVE_DIR}" && zip -q "../${ARCHIVE_NAME}" "$OUTPUT_NAME")
  else
    ARCHIVE_NAME="${BINARY}_${VERSION}_${GOOS}_${GOARCH}.tar.gz"
    (cd "$DIST_DIR" && tar -czf "$ARCHIVE_NAME" -C "${ARCHIVE_DIR}" "$OUTPUT_NAME")
  fi

  echo "  -> dist/${ARCHIVE_NAME}"
  rm -rf "$ARCHIVE_DIR"
done

echo ""
echo "=== Build complete ==="
echo "Archives in dist/:"
ls -lh "$DIST_DIR"/*.tar.gz "$DIST_DIR"/*.zip 2>/dev/null

echo ""
echo "=== Packaging npm ==="

# Copy README and skills to npm dir
cp "$PROJECT_DIR/README.md" "$NPM_DIR/README.md"
rm -rf "$NPM_DIR/skills"
cp -r "$PROJECT_DIR/skills" "$NPM_DIR/skills"

# Ensure bin dir exists and wrapper is executable
chmod +x "$NPM_DIR/bin/gitlink-cli"

echo ""
echo "=== Done ==="
echo ""
echo "Next steps:"
echo "  1. Upload dist/*.tar.gz to GitLink Release v${VERSION}"
echo "     URL: https://www.gitlink.org.cn/Gitlink/gitlink-cli/releases"
echo ""
echo "  2. Publish npm package:"
echo "     cd npm && npm publish --access public"
echo ""
echo "  Or for local-packed npm (includes binary for current platform):"
echo "     ./scripts/pack-local.sh"
