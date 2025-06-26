# Modelbuilder Tool

This tool is designed to generate and clean up Go structs created by the GORM `gentool` for database models.

## Prerequisites

- Go (latest version recommended)
- GORM `gentool` installed:
  ```bash
  go install gorm.io/gen/tools/gentool@latest
  ```

## Usage

Run the `modelbuilder` tool with the appropriate flags to generate structs from your database tables.

### Command Syntax

```bash
./modelbuilder [-flag value]
```

### Available Flags

| Flag | Description                          | Default Value                                  |
|------|--------------------------------------|------------------------------------------------|
| `-u` | Database user for connection         | `user`                                         |
| `-p` | Database password                    | (empty)                                        | 
| `-d` | Database name                        | `clinic-portal-makerble-golang-test`           |
| `-t` | Comma-separated list of table names  | `yourtables`                                   |
| `-db`| Database host IP or hostname         | `localhost`                                    |

### Example

To generate structs for `newtable1` and `newtable2` in the `clinic-portal-makerble-golang-test` database:

```bash
./modelbuilder -u username -p password -d clinic-portal-makerble-golang-test -t newtable1,newtable2 -db 127.0.0.1
```

## Building the Tool

The pre-built binary is for Linux 64-bit systems. If you are using a different operating system (e.g., Windows or macOS), build the tool manually:

```bash
go build -o modelbuilder-youros{.exe for Windows} ./x_module/modelbuilder/main.go
```

Replace `modelbuilder-youros` with your desired binary name, and append `.exe` for Windows.

## Notes

- Ensure the database is accessible and the provided credentials are correct before running the tool.
- The `-t` flag requires table names to be comma-separated without spaces (e.g., `table1,table2`).
- For additional details on `gentool`, refer to the [GORM documentation](https://gorm.io/gen/).