package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	. "github.com/tiegz/raytracer-go/examples"
)

var examples map[string]func() = map[string]func(){
	// Examples from the Raytracer Challenge
	"World":                       func() { RunDrawWorld() },
	"WorldWithPlane":              func() { RunDrawWorldWithPlane() },
	"WorldWithPatterns":           func() { RunDrawWorldWithPatterns() },
	"WorldWithMultiplePatterns":   func() { RunDrawWorldWithMultiplePatterns() },
	"WorldWithCube":               func() { RunDrawWorldWithCube() },
	"WorldWithTable":              func() { RunDrawWorldWithTable() },
	"WorldWithCylinderAndCone":    func() { RunDrawWorldWithCylinderAndCone() },
	"WorldWithHexagonGroup":       func() { RunDrawWorldWithHexagonGroup() },
	"WorldWithTriangles":          func() { RunDrawWorldWithTriangles() },
	"WorldWithTeapot":             func() { RunDrawWorldWithTeapot() },
	"WorldWithDice":               func() { RunDrawWorldWithDice() },
	"WorldWithCubeOfSpheres":      func() { RunDrawWorldWithCubeOfSpheres() },
	"WorldWithSphereAndAreaLight": func() { RunDrawWorldWithSphereAndAreaLight() },
	"WorldWithSnowman":            func() { RunDrawWorldWithSnowman() },
	"WorldWithUVPattern":          func() { RunDrawWorldWithUVPattern() },
	"UVAlignCheck":                func() { RunDrawUVAlignCheck() },
	"UVAlignCheckCubes":           func() { RunDrawUVAlignCheckCubes() },
	"UVImage":                     func() { RunDrawUVImage() },
	"Skybox":                      func() { RunDrawSkybox() },
	"Cover":                       func() { RunDrawCover() },

	// Other examples
	"Animation": func() { RunAnimation() },
}

func main() {
	// Root command
	cmd := flag.NewFlagSet("raytracer", flag.ExitOnError)

	// Sub-commands
	exampleCmd := flag.NewFlagSet("example", flag.ExitOnError)
	exampleNamePtr := exampleCmd.String("name", "", "Example name. (Default: World)")
	exampleListPtr := exampleCmd.Bool("list", false, "List examples.")
	versionCmd := flag.NewFlagSet("version", flag.ExitOnError)
	helpCmd := flag.NewFlagSet("help", flag.ExitOnError)

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "example":
		exampleCmd.Parse(os.Args[2:])
	case "version":
		versionCmd.Parse(os.Args[2:])
	case "help":
		helpCmd.Parse(os.Args[2:])
	default:
		cmd.Parse(os.Args[1:])
		printUsage()
		os.Exit(1)
	}

	if exampleCmd.Parsed() {
		if *exampleListPtr == true {
			fmt.Println("Listing examples: ")
			for k, _ := range examples {
				fmt.Printf("\t%s\n", k)
			}
		} else if len(*exampleNamePtr) > 0 {
			name := *exampleNamePtr
			fmt.Printf("Rendering example: %s\n", name)
			examples[name]()
		} else {
			printUsageForSubcommand("example", cmd, exampleCmd)
		}
	}

	if versionCmd.Parsed() {
		branch, err := exec.Command("git", "describe", "--tags").Output()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
		fmt.Printf("raytracer-go version: %s\n", strings.TrimSpace(string(branch)))
	}

	if helpCmd.Parsed() {
		if len(os.Args) < 3 {
			printUsage()
			os.Exit(1)
		}
		switch os.Args[2] {
		case "example":
			printUsageForSubcommand("example", cmd, exampleCmd)
		default:
			printUsage()
		}
	}
}

func printUsage() {
	fmt.Println("raytracer-go is a tool for rendering 3d scenes using raytracing.\n")
	fmt.Println("Usage:\n")
	fmt.Println("\traytracer <command> [arguments]\n")
	fmt.Println("The commands are:\n")
	fmt.Println("\texample\t\trender an example scene")
	fmt.Println("\tversion\t\tprint raytracer-go version")
	fmt.Println("\thelp   \t\tshow usage for a command (eg 'help example')")
	os.Exit(1)
}

func printUsageForSubcommand(name string, cmd, subCmd *flag.FlagSet) {
	fmt.Println("raytracer-go is a tool for rendering 3d scenes using raytracing.\n")
	fmt.Println("Usage:\n")
	fmt.Printf("\traytracer %s [arguments]\n", name)
	fmt.Println()
	fmt.Println("Flags:\n")
	subCmd.PrintDefaults()
	os.Exit(1)
}
