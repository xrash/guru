package root

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func run(opts cliopts, args []string) int {
	promptFromUser, err := acquirePrompt(opts)
	if err != nil {
		panic(err)
	}

	if len(args) == 0 {
		response, err := ollama(promptFromUser, opts.url, opts.model)
		if err != nil {
			panic(err)
		}
		fmt.Println(response)
		return 0
	}

	promptFiles := buildInputFiles(args)
	promptQuestion := "Given all files above, "
	prompt := promptFiles + promptQuestion + promptFromUser
	response, err := ollama(prompt, opts.url, opts.model)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
	return 0
}

func buildInputFiles(filenames []string) string {
	contents := make([]string, 0)

	for _, filename := range filenames {
		fi, err := os.Stat(filename)
		if err != nil {
			panic(err)
		}

		if fi.IsDir() {
			continue
		}

		f, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		b, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}

		content := fmt.Sprintf("This is file %s:\n\n%s\n", filename, b)

		contents = append(contents, content)
	}

	return strings.Join(contents, "\n\n\n")
}
