#!/usr/bin/env node

"use strict";

const path = require("path");
const { execFileSync } = require("child_process");

const ext = process.platform === "win32" ? ".exe" : "";
const binaryPath = path.join(__dirname, "gitlink-cli" + ext);

try {
  execFileSync(binaryPath, process.argv.slice(2), { stdio: "inherit" });
} catch (err) {
  if (err.status !== undefined) {
    process.exit(err.status);
  }
  console.error(`Failed to run gitlink-cli: ${err.message}`);
  console.error(
    "Binary may not be installed. Try reinstalling: npm install -g @gitlink-ai/cli"
  );
  process.exit(1);
}
