#!/usr/bin/env node

"use strict";

const os = require("os");
const path = require("path");
const fs = require("fs");
const https = require("https");
const http = require("http");
const { execSync } = require("child_process");

const PACKAGE = require("../package.json");
const VERSION = PACKAGE.version;
const BINARY_NAME = "gitlink-cli";

// GitLink release download base URL
// Format: https://www.gitlink.org.cn/Gitlink/gitlink-cli/releases
// Attachment download: https://www.gitlink.org.cn/api/attachments/{attachment_id}
const RELEASE_BASE = "https://www.gitlink.org.cn";
const REPO_OWNER = "Gitlink";
const REPO_NAME = "gitlink-cli";

function getPlatformInfo() {
  const platform = os.platform();
  const arch = os.arch();

  const platformMap = {
    darwin: "darwin",
    linux: "linux",
    win32: "windows",
  };

  const archMap = {
    x64: "amd64",
    arm64: "arm64",
  };

  const goPlatform = platformMap[platform];
  const goArch = archMap[arch];

  if (!goPlatform || !goArch) {
    throw new Error(
      `Unsupported platform: ${platform}-${arch}. ` +
        `Supported: darwin-x64, darwin-arm64, linux-x64, linux-arm64, win32-x64, win32-arm64`
    );
  }

  return { platform: goPlatform, arch: goArch, isWindows: platform === "win32" };
}

function getBinaryName(platform) {
  return platform === "windows" ? BINARY_NAME + ".exe" : BINARY_NAME;
}

function getArchiveName(platform, arch) {
  const ext = platform === "windows" ? ".zip" : ".tar.gz";
  return `${BINARY_NAME}_${VERSION}_${platform}_${arch}${ext}`;
}

function fetch(url, options = {}) {
  return new Promise((resolve, reject) => {
    const maxRedirects = options.maxRedirects || 5;
    let redirectCount = 0;

    function doRequest(currentUrl) {
      const mod = currentUrl.startsWith("https") ? https : http;
      const req = mod.get(currentUrl, (res) => {
        // Follow redirects
        if (
          (res.statusCode === 301 ||
            res.statusCode === 302 ||
            res.statusCode === 307 ||
            res.statusCode === 308) &&
          res.headers.location
        ) {
          redirectCount++;
          if (redirectCount > maxRedirects) {
            reject(new Error(`Too many redirects (max ${maxRedirects})`));
            return;
          }
          let redirectUrl = res.headers.location;
          if (redirectUrl.startsWith("/")) {
            const parsed = new URL(currentUrl);
            redirectUrl = `${parsed.protocol}//${parsed.host}${redirectUrl}`;
          }
          doRequest(redirectUrl);
          return;
        }

        if (res.statusCode !== 200) {
          reject(
            new Error(`HTTP ${res.statusCode} when downloading ${currentUrl}`)
          );
          return;
        }

        if (options.json) {
          let body = "";
          res.on("data", (chunk) => (body += chunk));
          res.on("end", () => {
            try {
              resolve(JSON.parse(body));
            } catch (e) {
              reject(new Error(`Failed to parse JSON: ${e.message}`));
            }
          });
        } else {
          const chunks = [];
          res.on("data", (chunk) => chunks.push(chunk));
          res.on("end", () => resolve(Buffer.concat(chunks)));
        }
      });
      req.on("error", reject);
      req.setTimeout(60000, () => {
        req.destroy();
        reject(new Error("Request timed out"));
      });
    }

    doRequest(url);
  });
}

async function findReleaseAsset(platform, arch) {
  const archiveName = getArchiveName(platform, arch);

  // Try fetching release info from GitLink API
  const apiUrl = `${RELEASE_BASE}/api/${REPO_OWNER}/${REPO_NAME}/releases.json`;
  console.log(`Fetching release info from ${apiUrl}`);

  try {
    const releases = await fetch(apiUrl, { json: true });

    // Find the release matching our version
    let release = null;
    const tagName = `v${VERSION}`;

    if (Array.isArray(releases)) {
      release = releases.find(
        (r) => r.tag_name === tagName || r.tag_name === VERSION
      );
      if (!release && releases.length > 0) {
        // Fall back to latest release
        release = releases[0];
      }
    } else if (releases && releases.releases) {
      const list = releases.releases;
      release = list.find(
        (r) => r.tag_name === tagName || r.tag_name === VERSION
      );
      if (!release && list.length > 0) {
        release = list[0];
      }
    }

    if (release && release.attachments) {
      // Try exact match first
      let asset = release.attachments.find(
        (a) => a.title === archiveName || a.filename === archiveName
      );
      // If no exact match (version mismatch on fallback release), match by platform/arch pattern
      if (!asset) {
        const ext = platform === "windows" ? ".zip" : ".tar.gz";
        const pattern = `_${platform}_${arch}${ext}`;
        asset = release.attachments.find(
          (a) => (a.title || a.filename || "").endsWith(pattern)
        );
      }
      if (asset) {
        // Return the download URL for this attachment
        let downloadUrl = asset.url || `${RELEASE_BASE}/api/attachments/${asset.id}`;
        // Ensure absolute URL
        if (downloadUrl.startsWith("/")) {
          downloadUrl = RELEASE_BASE + downloadUrl;
        }
        return downloadUrl;
      }
    }
  } catch (e) {
    console.log(`Warning: Could not fetch release info: ${e.message}`);
  }

  // Fallback: try direct download URL pattern
  return `${RELEASE_BASE}/api/${REPO_OWNER}/${REPO_NAME}/releases/${tagName}/assets/${archiveName}`;
}

async function downloadAndExtract(url, destDir, platform) {
  console.log(`Downloading ${BINARY_NAME} from ${url}...`);

  const data = await fetch(url);
  const isWindows = platform === "windows";
  const archiveExt = isWindows ? "download.zip" : "download.tar.gz";
  const archivePath = path.join(destDir, archiveExt);

  fs.writeFileSync(archivePath, data);
  console.log(`Downloaded ${(data.length / 1024 / 1024).toFixed(1)} MB`);

  // Extract
  if (isWindows) {
    execSync(
      `powershell -NoProfile -Command "Expand-Archive -Force -Path '${archivePath}' -DestinationPath '${destDir}'"`,
      { stdio: "pipe" }
    );
  } else {
    execSync(`tar -xzf "${archivePath}" -C "${destDir}"`, { stdio: "pipe" });
  }
  fs.unlinkSync(archivePath);

  // Find the binary in extracted files
  const binaryName = getBinaryName(platform);
  const binaryPath = path.join(destDir, binaryName);

  if (!fs.existsSync(binaryPath)) {
    // It might be in a subdirectory
    const files = fs.readdirSync(destDir);
    for (const file of files) {
      const subPath = path.join(destDir, file, binaryName);
      if (fs.existsSync(subPath)) {
        fs.renameSync(subPath, binaryPath);
        break;
      }
    }
  }

  if (!fs.existsSync(binaryPath)) {
    throw new Error(`Binary "${binaryName}" not found after extraction`);
  }

  // Make executable (not needed on Windows)
  if (!isWindows) {
    fs.chmodSync(binaryPath, 0o755);
  }
  console.log(`Installed ${BINARY_NAME} to ${binaryPath}`);
}

async function main() {
  try {
    const { platform, arch } = getPlatformInfo();
    console.log(`Platform: ${platform}-${arch}`);

    const binDir = path.join(__dirname, "..", "bin");
    if (!fs.existsSync(binDir)) {
      fs.mkdirSync(binDir, { recursive: true });
    }

    const binaryPath = path.join(binDir, getBinaryName(platform));

    // If binary already exists, check version matches
    if (fs.existsSync(binaryPath)) {
      try {
        const output = execSync(`"${binaryPath}" version`, { encoding: "utf-8", stdio: "pipe", timeout: 5000 });
        if (output.includes(VERSION)) {
          console.log(`${BINARY_NAME} v${VERSION} already installed, skipping download.`);
          return;
        }
        console.log(`${BINARY_NAME} version mismatch (got: ${output.trim()}, want: ${VERSION}), updating...`);
      } catch (e) {
        console.log(`${BINARY_NAME} binary exists but is not compatible, re-downloading...`);
      }
      fs.unlinkSync(binaryPath);
    }

    const downloadUrl = await findReleaseAsset(platform, arch);
    await downloadAndExtract(downloadUrl, binDir, platform);
  } catch (err) {
    console.error(`\nFailed to install ${BINARY_NAME}: ${err.message}`);
    console.error(
      `\nYou can install manually:\n` +
        `  1. Download from https://www.gitlink.org.cn/${REPO_OWNER}/${REPO_NAME}/releases\n` +
        `  2. Extract and place the binary in your PATH\n` +
        `  3. Or build from source: git clone && make build\n`
    );
    process.exit(1);
  }
}

main();
