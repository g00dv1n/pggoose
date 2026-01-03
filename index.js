#!/usr/bin/env node
import { spawn } from "node:child_process";
import os from "node:os";
import path from "node:path";
import { fileURLToPath } from "node:url";

const ARCH_MAP = {
  x64: "amd64",
  arm64: "arm64",
};

const suppotedOS = ["darwin", "linux"];

const platform = os.platform();
if (!suppotedOS.includes(platform)) {
  console.error(`Unsupported platform: ${platform}`);
  console.error(`Supported: ${suppotedOS.join(", ")}`);
  process.exit(1);
}
const arch = os.arch();
if (!ARCH_MAP[arch]) {
  console.error(`Unsupported ARCH: ${arch}`);
  console.error(`Supported: ${Object.keys(ARCH_MAP).join(", ")}`);
  process.exit(1);
}

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const binName = `pggoose-${platform}-${ARCH_MAP[arch]}`;

const binaryPath = path.join(__dirname, "dist", binName);

const child = spawn(binaryPath, process.argv.slice(2), {
  stdio: "inherit",
});

child.on("exit", (code) => process.exit(code || 0));
child.on("error", (err) => {
  console.error("Failed to start pggoose:", err.message);
  process.exit(1);
});
