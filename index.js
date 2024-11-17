#!/usr/bin/env node

const { execSync } = require('child_process');
const { platform } = require('os');
const path = require('path');

const binaries = {
  darwin: 'birdcli-macos',
  win32: 'birdcli-windows.exe',
  linux: 'birdcli-linux',
};

const currentPlatform = platform();
const binaryName = binaries[currentPlatform];

if (!binaryName) {
  console.error(`Unsupported platform: ${currentPlatform}`);
  process.exit(1);
}

// Use process.cwd() to get the correct path for the global install
const binaryPath = path.join(process.cwd(), 'dist', binaryName);

try {
  execSync(`${binaryPath} ${process.argv.slice(2).join(' ')}`, { stdio: 'inherit' });
} catch (error) {
  console.error(`Failed to execute binary: ${error.message}`);
  process.exit(1);
}
