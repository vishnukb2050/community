#!/bin/bash
set -e

echo "Starting Build and Deploy Process..."
FLUTTER_CMD="/home/vishnu/development/flutter/bin/flutter"

# 1. Clean and Build
echo "Cleaning project..."
$FLUTTER_CMD clean

echo "Building Release APK..."
$FLUTTER_CMD build apk --release

# 2. Verify APK exists
APK_PATH="build/app/outputs/flutter-apk/app-release.apk"
if [ ! -f "$APK_PATH" ]; then
    echo "Error: APK not found at $APK_PATH"
    exit 1
fi
echo "APK built successfully at $APK_PATH"

# 3. Detect Windows Downloads Path (WSL specific)
echo "Detecting Windows Downloads folder..."
if command -v cmd.exe &> /dev/null; then
    # Get Windows UserProfile (e.g., C:\Users\Vishnu)
    WIN_PROFILE=$(cmd.exe /c "echo %USERPROFILE%" 2>/dev/null | tr -d '\r')
    
    # Convert to WSL path (e.g., /mnt/c/Users/Vishnu)
    # Using wslpath if available, else manual conversion
    if command -v wslpath &> /dev/null; then
        WSL_PROFILE=$(wslpath -u "$WIN_PROFILE")
    else
        # Fallback for simple drive mapping
        WSL_PROFILE=$(echo "$WIN_PROFILE" | sed 's/\\/\//g' | sed 's/C:/\/mnt\/c/')
    fi
    
    DOWNLOADS_DIR="$WSL_PROFILE/Downloads"
    echo "Target Directory: $DOWNLOADS_DIR"
    
    # 4. Copy
    if [ -d "$DOWNLOADS_DIR" ]; then
        cp "$APK_PATH" "$DOWNLOADS_DIR/"
        echo "✅ Success! Copied app-release.apk to Windows Downloads."
    else
        echo "❌ Error: Could not verify Downloads directory at $DOWNLOADS_DIR"
        echo "You can manually copy the file from: $(pwd)/$APK_PATH"
    fi
else
    echo "cmd.exe not found. Assuming non-WSL environment."
    echo "Build complete. APK is located at: $(pwd)/$APK_PATH"
fi
