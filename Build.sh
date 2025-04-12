#!/bin/bash
echo "Building DVPocketBase for Linux..."

# Create builds directory if it doesn't exist
mkdir -p builds

# Build the Linux executable
go build -v -o builds/dvpocketbase-linux ./main.go

if [ $? -eq 0 ]; then
    # Make the binary executable
    chmod +x builds/dvpocketbase-linux
    echo "Build completed successfully! Executable saved to builds/dvpocketbase-linux"
else
    echo "Build failed with error code $?"
fi
