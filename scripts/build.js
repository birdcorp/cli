const { execSync } = require('child_process');
const path = require('path');
const fs = require('fs');

const binaries = {
  darwin: 'birdcli-macos',
  win32: 'birdcli-windows.exe',
  linux: 'birdcli-linux',
};

const distDir = path.join(__dirname, '..', 'dist');

// Ensure the dist directory exists
if (!fs.existsSync(distDir)) {
  fs.mkdirSync(distDir, { recursive: true });
}

// Platforms and architectures to build
const targets = [
  { os: 'darwin', arch: 'amd64', output: binaries.darwin },
  { os: 'linux', arch: 'amd64', output: binaries.linux },
  { os: 'windows', arch: 'amd64', output: binaries.win32 },
];

// Build binaries
targets.forEach((target) => {
  const outputPath = path.join(distDir, target.output);
  console.log(`Building for ${target.os}-${target.arch}...`);

  try {
    execSync(
      `GOOS=${target.os} GOARCH=${target.arch} go build -o ${outputPath}`,
      { stdio: 'inherit', env: { ...process.env } }
    );
    console.log(`Built: ${outputPath}`);
  } catch (error) {
    console.error(`Failed to build for ${target.os}-${target.arch}: ${error.message}`);
  }
});

console.log('Build completed successfully');