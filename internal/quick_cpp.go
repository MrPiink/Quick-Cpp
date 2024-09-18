package internal

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/schollz/progressbar/v3"
)

func check(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func readFile(path string) string {
	content, err := os.ReadFile(path)
	check(err)
	return string(content)
}

const msys2DownloadURL = "https://github.com/msys2/msys2-installer/releases/download/2024-07-27/msys2-x86_64-20240727.exe"

var cmakeListsRootContent = readFile("../assets/CMakeListsRoot.txt")

var cmakePresetsContent = readFile("../assets/CMakePresets.json")

var readmeContent = readFile("../assets/README.md")

var mainCppContent = readFile("../assets/maincpp.txt")

var cmakeListsSrcContent = readFile("../assets/CMakeListsSrc.txt")

func add_files(projectName string) {
	os.Chdir(projectName)

	check(os.WriteFile("CMakeLists.txt", []byte(cmakeListsRootContent), 0644))
	check(os.WriteFile("CMakePresets.json", []byte(cmakePresetsContent), 0644))
	check(os.WriteFile("README.md", []byte(readmeContent), 0644))

	os.Chdir("src")
	
	check(os.WriteFile("main.cpp", []byte(mainCppContent), 0644))
	check(os.WriteFile("CMakeLists.txt", []byte(cmakeListsSrcContent), 0644))
	
	os.Chdir("../../")
}

func printProgressBar(done chan error) {
	bar := progressbar.Default(100)
	timeInterval := 1000 * time.Millisecond

	for i := 0; ; i++ {
		select {
		
		case <-done:
			bar.Add(100 - i)
			return
		
		default:
			if i != 0 && i % 99 == 0 {
				bar.Set(75)
				timeInterval = 500 * time.Millisecond

			} else {
				bar.Add(1)
			}

			time.Sleep(timeInterval)
		}
	}
}

func installMsys2(done chan error) {
	fmt.Println("Installing MSYS2...")
	cmd := exec.Command(".\\msys2-x86_64-latest.exe", "in", "--confirm-command", "--accept-messages", "--root", "C:/msys64")
	err := cmd.Run()
	check(err)

	done <- nil
}

func installToolchain() {
	fmt.Println("Installing Toolchain...")
	cmd := exec.Command("C:\\msys64\\usr\\bin\\bash.exe", "-l", "-c", "pacman -S --needed base-devel --noconfirm mingw-w64-ucrt-x86_64-toolchain")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	check(err)

	fmt.Println("Installing CMake...")
	cmd = exec.Command("C:\\msys64\\usr\\bin\\bash.exe", "-l", "-c", "pacman -S --noconfirm mingw-w64-ucrt-x86_64-cmake")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	check(err)
}

func msys2(projectName string) {
	os.Chdir(projectName)
	defer os.Chdir("../")
	
	out, err := os.Create("msys2-x86_64-latest.exe")
	check(err)

	resp, err := http.Get(msys2DownloadURL)
	check(err)

	_, err = io.Copy(out, resp.Body)
	check(err)

	out.Close()
	resp.Body.Close()

	msys2Done := make(chan error)

	go installMsys2(msys2Done)

	printProgressBar(msys2Done)

	defer os.Remove("msys2-x86_64-latest.exe")

	cmd := exec.Command("powershell", "[System.Environment]::GetEnvironmentVariable('PATH','User')")
	cmd.Stderr = os.Stderr
	userPath, err := cmd.Output()
	userPathString := string(userPath)
	check(err)

	if len(userPathString) > 0 && userPathString[len(userPathString)-1] != ';' {
		userPathString += ";"
	}

	msys2Path := "C:\\msys64\\ucrt64\\bin"
	cmd = exec.Command("setx", "PATH", userPathString + msys2Path)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	check(err)

	path := os.Getenv("PATH")
	
	if len(path) > 0 && path[len(path)-1] != ';' {
		path += ";"
	}
	
	os.Setenv("PATH", path + msys2Path)

	installToolchain()

	cmd = exec.Command("gcc", "--version")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	check(err)

	cmd = exec.Command("cmake", "--version")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	check(err)

	fmt.Println("The mingw-w64 toolchain and CMake are installed!")
}

func Create(projectName string, full bool, skipMsys bool, noFiles bool) {
	err := os.MkdirAll(projectName, 0755)
	check(err)

	os.Chdir(projectName)

	err = os.MkdirAll("build", 0755)
	check(err)
	err = os.MkdirAll("src", 0755)
	check(err)
	err = os.MkdirAll("tests", 0755)
	check(err)
	err = os.MkdirAll("data", 0755)
	check(err)
	err = os.MkdirAll("tools", 0755)
	check(err)
	err = os.MkdirAll("docs", 0755)
	check(err)

	if full {
		err = os.MkdirAll("include", 0755)
		check(err)
		err = os.MkdirAll("examples", 0755)
		check(err)
		err = os.MkdirAll("external", 0755)
		check(err)
		err = os.MkdirAll("libs", 0755)
		check(err)
		err = os.MkdirAll("extras", 0755)
		check(err)
	}

	os.Chdir("../")

	if !noFiles {
		add_files(projectName)
	}

	if !skipMsys {
		msys2(projectName)
	}
}

func Revert(projectName string, msys2 bool) {
	err := os.RemoveAll(projectName)
	check(err)

	if msys2 {
		cmd := exec.Command("powershell", "[System.Environment]::GetEnvironmentVariable('PATH','User')")
		cmd.Stderr = os.Stderr
		userPath, err := cmd.Output()
		userPathString := string(userPath)
		check(err)

		newPath := userPathString

		splitPath := strings.Split(userPathString, ";")

		for i, dir := range splitPath {
			dir = strings.TrimSpace(dir)

			if dir == "C:\\msys64\\ucrt64\\bin" {

				if i != 0 {
					dir = ";" + dir
				}

				fmt.Println(newPath)
				newPath = strings.Replace(userPathString, dir, "", 1)
			}
		}
		
		fmt.Println(newPath)
		cmd = exec.Command("setx", "PATH", newPath)
		err = cmd.Run()
		check(err)

		os.Setenv("PATH", newPath)

		cmd = exec.Command("C:\\msys64\\uninstall.exe", "pr", "--confirm-command")
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		check(err)
	}
}