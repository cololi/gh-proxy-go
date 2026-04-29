#!/bin/bash
set -e

REPO="cololi/Hub-Proxy-Go"
BINARY_NAME="hub-proxy-go"

# Check dependencies
if ! command -v curl >/dev/null 2>&1; then
    echo "Error: curl is not installed."
    exit 1
fi

if ! command -v sudo >/dev/null 2>&1; then
    echo "Error: sudo is not installed. This script requires root privileges to set up systemd."
    exit 1
fi

# Detect architecture
ARCH=$(uname -m)
case $ARCH in
    x86_64)
        FILE_ARCH="linux-amd64"
        ;;
    aarch64|arm64)
        FILE_ARCH="linux-arm64"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

echo "Detecting latest version..."
LATEST_TAG=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_TAG" ]; then
    echo "Error: Could not find latest release."
    exit 1
fi

DOWNLOAD_URL="https://github.com/$REPO/releases/download/$LATEST_TAG/${BINARY_NAME}-${FILE_ARCH}"

echo "Downloading ${BINARY_NAME} ${LATEST_TAG} for ${FILE_ARCH}..."
curl -L -o /tmp/${BINARY_NAME} "$DOWNLOAD_URL"
chmod +x /tmp/${BINARY_NAME}

echo "Installing binary to /usr/local/bin/..."
sudo mv /tmp/${BINARY_NAME} /usr/local/bin/${BINARY_NAME}

echo "Creating systemd service..."
sudo tee /etc/systemd/system/${BINARY_NAME}.service > /dev/null <<EOF
[Unit]
Description=Hub-Proxy-Go Service
After=network.target

[Service]
ExecStart=/usr/local/bin/${BINARY_NAME}
Restart=always
User=root
Environment=LISTEN=:8080

[Install]
WantedBy=multi-user.target
EOF

echo "Starting service..."
sudo systemctl daemon-reload
sudo systemctl enable --now ${BINARY_NAME}

echo "Successfully installed ${BINARY_NAME}!"
echo "Status: "
systemctl status ${BINARY_NAME} --no-pager
