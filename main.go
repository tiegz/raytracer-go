package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	. "github.com/tiegz/raytracer-go/examples"
)

var examples map[string]func(int) = map[string]func(int){
	// Examples from the Raytracer Challenge
	"World":                       func(jobs int) { RunDrawWorld(jobs) },
	"WorldWithPlane":              func(jobs int) { RunDrawWorldWithPlane(jobs) },
	"WorldWithPatterns":           func(jobs int) { RunDrawWorldWithPatterns(jobs) },
	"WorldWithMultiplePatterns":   func(jobs int) { RunDrawWorldWithMultiplePatterns(jobs) },
	"WorldWithCube":               func(jobs int) { RunDrawWorldWithCube(jobs) },
	"WorldWithTable":              func(jobs int) { RunDrawWorldWithTable(jobs) },
	"WorldWithCylinderAndCone":    func(jobs int) { RunDrawWorldWithCylinderAndCone(jobs) },
	"WorldWithHexagonGroup":       func(jobs int) { RunDrawWorldWithHexagonGroup(jobs) },
	"WorldWithTriangles":          func(jobs int) { RunDrawWorldWithTriangles(jobs) },
	"WorldWithTeapot":             func(jobs int) { RunDrawWorldWithTeapot(jobs) },
	"WorldWithDice":               func(jobs int) { RunDrawWorldWithDice(jobs) },
	"WorldWithCubeOfSpheres":      func(jobs int) { RunDrawWorldWithCubeOfSpheres(jobs) },
	"WorldWithSphereAndAreaLight": func(jobs int) { RunDrawWorldWithSphereAndAreaLight(jobs) },
	"WorldWithSnowman":            func(jobs int) { RunDrawWorldWithSnowman(jobs) },
	"WorldWithUVPattern":          func(jobs int) { RunDrawWorldWithUVPattern(jobs) },
	"UVAlignCheck":                func(jobs int) { RunDrawUVAlignCheck(jobs) },
	"UVAlignCheckCubes":           func(jobs int) { RunDrawUVAlignCheckCubes(jobs) },
	"UVImage":                     func(jobs int) { RunDrawUVImage(jobs) },
	"Skybox":                      func(jobs int) { RunDrawSkybox(jobs) },
	"Cover":                       func(jobs int) { RunDrawCover(jobs) },

	// Other examples
	"Animation": func(jobs int) { RunAnimation(jobs) },
}

func main() {
	// Root command
	cmd := flag.NewFlagSet("raytracer", flag.ExitOnError)

	// Sub-commands
	exampleCmd := flag.NewFlagSet("example", flag.ExitOnError)
	exampleNamePtr := exampleCmd.String("name", "", "Example name. (Default: World)")
	exampleListPtr := exampleCmd.Bool("list", false, "List examples.")
	exampleJobsPtr := exampleCmd.Int("jobs", 1, "Run n jobs in parallel. (Default: 1)")
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
			if f, ok := examples[name]; ok {
				f(*exampleJobsPtr)
			} else {
				fmt.Printf("Example %s not found!\nRun 'raytracer-go example -list' to see examples.\n", name)
			}
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
	fmt.Println("raytracer-go is a tool for rendering 3d scenes using raytracing.")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("\traytracer <command> [arguments]")
	fmt.Println()
	fmt.Println("The commands are:")
	fmt.Println()
	fmt.Println("\texample\t\trender an example scene")
	fmt.Println("\tversion\t\tprint raytracer-go version")
	fmt.Println("\thelp   \t\tshow usage for a command (eg 'help example')")
	os.Exit(1)
}

func printUsageForSubcommand(name string, cmd, subCmd *flag.FlagSet) {
	fmt.Println("raytracer-go is a tool for rendering 3d scenes using raytracing.")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Printf("\traytracer %s [arguments]", name)
	fmt.Println()
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println()
	subCmd.PrintDefaults()
	os.Exit(1)
}
