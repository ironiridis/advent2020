package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

const TestFileTemplate = `package main

import (
	"testing"

	"github.com/ironiridis/advent2020/scando"
)

func TestPart1Example(t *testing.T) {
	scando.Strap(part1func, "example data", "expected output", t)
}

func TestPart2Example(t *testing.T) {
	t.Fail()
}
`

const MainFileTemplate = `package main

import (
	"fmt"

	"github.com/ironiridis/advent2020/scando"
)

func part1func(in chan string) (string, error) {
	return "", fmt.Errorf("unimplemented")
}

func part2func(in chan string) (string, error) {
	return "", fmt.Errorf("unimplemented")
}

func main() {
	fmt.Println("Day {{.}}, part 1 - summary")
	ans, err := part1func(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 1 Answer: %q\n", ans)
	fmt.Println("Day {{.}}, part 2 - summary")
	ans, err = part2func(scando.Input())
	if err != nil {
		fmt.Printf("Cannot determine answer: %v\n", err)
		return
	}
	fmt.Printf("Part 2 Answer: %q\n", ans)
}
`

func dayfmt(d int) string {
	return fmt.Sprintf("day%02d", d)
}

func makeFolder(f string, d int) (err error) {
	testFile, err := template.New("testfile").Parse(TestFileTemplate)
	if err != nil {
		return
	}
	mainFile, err := template.New("mainfile").Parse(MainFileTemplate)
	if err != nil {
		return
	}

	err = os.Mkdir(f, os.ModeDir|os.ModePerm)
	if err != nil {
		return
	}

	testfp, err := os.Create(fmt.Sprintf("%s/day%02d_test.go", f, d))
	if err != nil {
		return
	}
	defer testfp.Close()
	err = testFile.Execute(testfp, d)
	if err != nil {
		return
	}

	mainfp, err := os.Create(fmt.Sprintf("%s/day%02d.go", f, d))
	if err != nil {
		return
	}
	defer mainfp.Close()
	err = mainFile.Execute(mainfp, d)
	if err != nil {
		return
	}

	return
}

func findNextDayFolder() int {
	panic("🤷‍♂️")
	return 0
}

func main() {
	var err error
	folder := flag.String("folder", "", "explicitly specify folder to create")
	day := flag.Int("day", 0, "specify day to create")
	flag.Parse()
	if *folder != "" {
		fmt.Printf("Making folder explicitly at %q\n", *folder)
		if *day == 0 {
			fmt.Println("Day is not set. Specify day number with -day")
			return
		}
		err = makeFolder(*folder, *day)
	} else if *day != 0 {
		fmt.Printf("Making folder at %q\n", dayfmt(*day))
		err = makeFolder(dayfmt(*day), *day)
	} else {
		fmt.Println("Searching for next non-existing folder")
		*day = findNextDayFolder()
		err = makeFolder(dayfmt(*day), *day)
	}
	if err != nil {
		fmt.Printf("failed: %v\n", err)
	}
}
