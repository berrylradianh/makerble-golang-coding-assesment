# Repogenerator Tool

This tool is designed to generate repository and service code for your Go application, streamlining the process of creating data access and business logic layers.

## Prerequisites

- Go (latest version recommended)
- Pre-generated models (e.g., using the `modelbuilder` tool)

## Usage

Run the `repogenerator` tool to generate repository and service code.

### Command Syntax

```bash
./repogenerator
```

No additional flags are required, as the tool uses the existing model definitions in your project.

## Building the Tool

The pre-built binary is for Linux 64-bit systems. If you are using a different operating system (e.g., Windows or macOS), build the tool manually:

```bash
go build -o repogenerator-youros{.exe for Windows} ./x_module/repogenerator/main.go
```

Replace `repogenerator-youros` with your desired binary name, and append `.exe` for Windows.

## Notes

- Ensure that the model files are correctly generated and available in the project before running the tool.
- The tool assumes a standard project structure; verify that the paths in your configuration align with the expected directory layout.
- For additional customization, review the generated repository and service code to ensure it meets your application's requirements.