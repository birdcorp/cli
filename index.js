const { execFileSync } = require('child_process');
const path = require('path');

const binaryPath = path.join(__dirname, 'dist', `birdcli-${process.platform}`);

try {
  const output = execFileSync(binaryPath, [], { encoding: 'utf8' });
  console.log(output);
} catch (error) {
  console.error('Error executing binary:', error);
}
