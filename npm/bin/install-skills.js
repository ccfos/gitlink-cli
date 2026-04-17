#!/usr/bin/env node

"use strict";

const path = require("path");
const { execSync } = require("child_process");

const skillsDir = path.join(__dirname, "..", "skills");

console.log(`Installing gitlink-cli skills from ${skillsDir}...`);

try {
  execSync(`npx skills add "${skillsDir}" -y -g`, { stdio: "inherit" });
} catch (err) {
  if (err.status !== undefined) {
    process.exit(err.status);
  }
  console.error(`Failed to install skills: ${err.message}`);
  console.error(`\nYou can install manually:`);
  console.error(`  npx skills add "${skillsDir}" -y -g`);
  process.exit(1);
}
