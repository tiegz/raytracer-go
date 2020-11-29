package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	. "github.com/tiegz/raytracer-go/examples"
)

var examples map[string]func(bool, int) = map[string]func(bool, int){
	// Examples from the Raytracer Challenge
	"World":                       func(printProgress bool, jobs int) { RunDrawWorld(printProgress, jobs) },
	"WorldWithPlane":              func(printProgress bool, jobs int) { RunDrawWorldWithPlane(printProgress, jobs) },
	"WorldWithPatterns":           func(printProgress bool, jobs int) { RunDrawWorldWithPatterns(printProgress, jobs) },
	"WorldWithSphere":             func(printProgress bool, jobs int) { RunDrawWorldWithSphere(printProgress, jobs) },
	"WorldWithMultiplePatterns":   func(printProgress bool, jobs int) { RunDrawWorldWithMultiplePatterns(printProgress, jobs) },
	"WorldWithCube":               func(printProgress bool, jobs int) { RunDrawWorldWithCube(printProgress, jobs) },
	"WorldWithTable":              func(printProgress bool, jobs int) { RunDrawWorldWithTable(printProgress, jobs) },
	"WorldWithCylinderAndCone":    func(printProgress bool, jobs int) { RunDrawWorldWithCylinderAndCone(printProgress, jobs) },
	"WorldWithHexagonGroup":       func(printProgress bool, jobs int) { RunDrawWorldWithHexagonGroup(printProgress, jobs) },
	"WorldWithTriangles":          func(printProgress bool, jobs int) { RunDrawWorldWithTriangles(printProgress, jobs) },
	"WorldWithTeapot":             func(printProgress bool, jobs int) { RunDrawWorldWithTeapot(printProgress, jobs) },
	"WorldWithDice":               func(printProgress bool, jobs int) { RunDrawWorldWithDice(printProgress, jobs) },
	"WorldWithCubeOfSpheres":      func(printProgress bool, jobs int) { RunDrawWorldWithCubeOfSpheres(printProgress, jobs) },
	"WorldWithSphereAndAreaLight": func(printProgress bool, jobs int) { RunDrawWorldWithSphereAndAreaLight(printProgress, jobs) },
	"WorldWithSnowman":            func(printProgress bool, jobs int) { RunDrawWorldWithSnowman(printProgress, jobs) },
	"WorldWithUVPattern":          func(printProgress bool, jobs int) { RunDrawWorldWithUVPattern(printProgress, jobs) },
	"UVAlignCheck":                func(printProgress bool, jobs int) { RunDrawUVAlignCheck(printProgress, jobs) },
	"UVAlignCheckCubes":           func(printProgress bool, jobs int) { RunDrawUVAlignCheckCubes(printProgress, jobs) },
	"UVImage":                     func(printProgress bool, jobs int) { RunDrawUVImage(printProgress, jobs) },
	"Skybox":                      func(printProgress bool, jobs int) { RunDrawSkybox(printProgress, jobs) },
	"Cover":                       func(printProgress bool, jobs int) { RunDrawCover(printProgress, jobs) },

	// Other examples
	"Animation": func(printProgress bool, jobs int) { RunAnimation(printProgress, jobs) },
}

func main() {
	// Root command
	cmd := flag.NewFlagSet("raytracer", flag.ExitOnError)

	// --> example sub-command
	exampleCmd := flag.NewFlagSet("example", flag.ExitOnError)
	exampleNamePtr := exampleCmd.String("name", "", "Example name.")
	exampleListPtr := exampleCmd.Bool("list", false, "List examples.")
	examplePrintProgressPtr := exampleCmd.Bool("progress", true, "Write progress to stdout.")
	exampleJobsPtr := exampleCmd.Int("jobs", 1, "Run n jobs in parallel.")

	// --> version sub-command
	versionCmd := flag.NewFlagSet("version", flag.ExitOnError)

	// --> help sub-command
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
				f(*examplePrintProgressPtr, *exampleJobsPtr)
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
