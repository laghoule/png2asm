# png2asm

`png2asm` is a lightweight Go tool designed to convert 256-color paletted PNG images into assembly-compatible `.inc` files. This is particularly useful for retro game development, embedded systems, or any scenario where you need to integrate graphical assets directly into assembly code.

The tool extracts the color index for each pixel from the input PNG and outputs it as a byte array, ready to be included in your assembly projects.

## Features

*   Converts 256-color paletted PNG images.
*   Generates an assembly `.inc` file containing pixel color indices.
*   Supports custom input PNG file names.
*   Includes comments in the generated `.inc` file for clarity.

## Requirements

*   **Go Language**: You need to have Go installed on your system (if you want to build from source).
*   **256-color Paletted PNG Image**: The input image *must* be a PNG with a 256-color palette (indexed color mode). The tool will terminate with an error if a non-paletted image is provided.

## Usage

1.  **Build the Executable**:
    Navigate to the project directory and build the executable:
    ```bash
    go build -o png2asm main.go
    ```

2.  **Prepare Your Image**:
    Place your 256-color paletted PNG image in the same directory as the `png2asm` executable, or specify its path.

3.  **Run the Tool**:
    *   **Using the default input file (`sprite.png`)**:
        If your input image is named `sprite.png`:
        ```bash
        ./png2asm
        ```
    *   **Specifying a custom input file**:
        Use the `-src` flag to provide the path to your PNG image:
        ```bash
        ./png2asm -src path/to/your_image.png
        ```

    The tool will generate a file named `sprite.inc` in the current directory.

## Example Output (`sprite.inc`)

Given an input `sprite.png` (for example, an 8x8 sprite), the generated `sprite.inc` file might look something like this:

```assembly
; Automatically generated sprite
my_sprite LABEL BYTE
    DB 0,1,2,3,4,5,6,7
    DB 8,9,10,11,12,13,14,15
    DB 16,17,18,19,20,21,22,23
    DB 24,25,26,27,28,29,30,31
    DB 32,33,34,35,36,37,38,39
    DB 40,41,42,43,44,45,46,47
    DB 48,49,50,51,52,53,54,55
    DB 56,57,58,59,60,61,62,63
```
