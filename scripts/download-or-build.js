const { execSync } = require('child_process');
const os = require('os');
const path = require('path');
const fs = require('fs');

const platform = os.platform();
const arch = os.arch();
let binaryPath = '';

switch (platform) {
  case 'win32':
    binaryPath = path.join(__dirname, '..', 'dist', 'birdcli-windows.exe');
    break;
  case 'darwin':
    binaryPath = path.join(__dirname, '..', 'dist', 'birdcli-macos');
    break;
  case 'linux':
    binaryPath = path.join(__dirname, '..', 'dist', 'birdcli-linux');
    break;
  default:
    console.error(`Unsupported platform: ${platform}`);
    process.exit(1);
}

if (!fs.existsSync(binaryPath)) {
    console.error(`Binary does not exist at ${binaryPath}`);
    process.exit(1);
} else {
  console.log(`Binary already exists at ${binaryPath}`);
}

// Make the binary executable (for UNIX systems)
if (platform !== 'win32') {
  fs.chmodSync(binaryPath, 0o755);
}