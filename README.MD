# DVPocketBase

A backend database and API server for DV built with [PocketBase](https://pocketbase.io/), providing user management, inventory tracking, and other features for the DV ecosystem.

## Features

- User management
- Item inventory system
- Economy with balance tracking
- Infraction management system
- Automatic daily rewards system
- Cross-platform (Windows & Linux)

## Installation (for Devs)

### Prerequisites

- Go 1.23.4 or higher
- Git

### Installing

1. Clone the repository:

   ``` bash
   git clone https://github.com/WOFTNW/DVPocketBase.git
   cd DVPocketBase
   ```

2. Build for your platform:

   **Windows:**

   ``` powershell
   ./Build.bat
   ```

   **Linux:**

   ``` bash
   ./Build.sh
   ```

## Usage

### Running the server

**Windows:**

  ``` powershell
  ./run.bat serve
  ```

**Linux:**

  ``` bash
  ./run.sh serve
```

By default, the server will start on <http://localhost:8090> with the admin dashboard available at <http://localhost:8090/_/>

### Server Configuration

after running the server head to the dashboard to create a superuser account
that will allow you to manage the dashboard

## Database Schema

to load the Schema head to the settings tab in the Dashboard
then under Sync go to import collections
hit the load From JSON file button
pick the json file under /schemea/pb_schema.json

The database includes the following collections:

- **users**: Stores user information including Discord ID, Minecraft username, balance, etc.
- **items**: Configurable items with prices, descriptions, and actions
- **user_inventory**: Many-to-many relationship between users and items
- **infractions**: User warning and moderation records

## Development

### Project Structure

- `/main.go` - Application entry point
- `/builds/` - Compiled executables (ignored by Git)
- `/pb_data/` - Database and storage files
- `/pb_public/` - Public static files
- `/scheme/` - Database schema definitions

### Building for Development

Use the build scripts to compile for your platform:
Windows

``` powershell
./Build.bat
```

Linux

``` bash
./Build.sh
```

### GitHub Actions

The project includes CI/CD workflows that:

1. Build the application for Windows and Linux
2. Run tests
