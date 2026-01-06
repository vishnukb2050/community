#!/bin/bash
export PATH="$HOME/development/flutter/bin:$HOME/development/jdk-17/bin:$HOME/development/android-sdk/cmdline-tools/latest/bin:$HOME/development/android-sdk/platform-tools:$PATH"
export JAVA_HOME="$HOME/development/jdk-17"
export ANDROID_HOME="$HOME/development/android-sdk"
export ANDROID_SDK_ROOT="$HOME/development/android-sdk"

echo "Configuring Flutter..."
flutter config --android-sdk "$ANDROID_HOME"
flutter config --jdk-dir "$JAVA_HOME"

cd "$HOME/socwhiz/community/android-app"

# Ensure local.properties exists
echo "sdk.dir=$ANDROID_HOME" > android/local.properties
echo "flutter.sdk=$HOME/development/flutter" >> android/local.properties

echo "Fetching dependencies..."
flutter pub get

echo "Accepting licenses..."
yes | flutter doctor --android-licenses

echo "Starting build..."
flutter build apk --debug --verbose
