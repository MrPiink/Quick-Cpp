# Quick C++

## Overview

Quick C++ is a GoLang CLI tool that automates setting up a new C++ project with the pitchfork layout and CMake. If you do not already have Msys2 it will automatically install it for you and add it to the path. This project aims to make developing in C++ faster and more convenient, especially for new C++ users on Windows that do not wish to use a heavy IDE like Visual Studio.

## Features

- **Command 1**: `qcpp create <project_name>`

  - Description: Creates a new project directory with a slim Pitchfork layout, adds CMakeLists and CMakePresets, and installs Msys2 along with the necessary Ming64-gw packages.

  - Flags:
    - `--full`: Includes the less commonly used directories from the Pitchfork layout.
    - `--skip-msys2`: Skips installing Msys2.
    - `--no-files`: Omits adding the CMake and example files.

- **Command 2**: `qcpp revert <project_name>`

  - Description: Deletes the specified project directory with the option to completely remove Msys2 and the added path in the PATH environment variable.

  - Flags:
    - `--msys2`: Uninstalls Msys2 along with reverting the users PATH environment variable.

## Installation

To install Quick C++, follow these steps:

1. Download the latest release from [releases page](https://github.com/MrPiink/Quick-Cpp/releases).

2. Extract the zip

3. Add Quick-Cpp to the user PATH environment variable.

## Usage

To use Quick C++, follow these steps:

1. In a terminal located where you want to create your project

   Enter the command:
   `qcpp create MyProject`

2. Add flags to customize your setup

   Enter the command:
   `qcpp create MyProject2 --skipmsys`

3. To reverse changes made by Quick C++

   Enter the command:
   `qcpp revert MyProject`

## License

This project is licensed under the [GNU General Public License v3] - see the [LICENSE](LICENSE) file for details.
