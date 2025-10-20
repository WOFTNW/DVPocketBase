#!/bin/bash
echo "Running DVPocketBase for Linux..."

# Check if the executable exists, if not build it
if [ ! -f "builds/dvpocketbase-linux" ]; then
    echo "Executable not found. Building first..."
    bash ./Build.sh
fi

# Make sure it's executable
chmod +x builds/dvpocketbase-linux

# Run the application
echo "Starting DVPocketBase..."
./builds/dvpocketbase-linux "$@"

echo "Application exited with code $?"
