package root

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func acquirePrompt(opts cliopts) (string, error) {

	// If prompt is passed through --prompt, that's it.
	if opts.prompt != "" {
		return opts.prompt, nil
	}

	// If --stdin is passed, prompt is supposed to be read from stdin.
	if opts.stdin {
		b, err := io.ReadAll(os.Stdin)
		if err != nil {
			return "", fmt.Errorf("error reading prompt from stdin: %v", err)
		}

		if len(b) == 0 {
			return "", fmt.Errorf("prompt from stdin has zero length")
		}

		return string(b), nil
	}

	// If no option is supplied, read prompt interactively.
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	prompt, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("error reading prompt from interactive input: %v", err)
	}

	if len(prompt) == 0 {
		return "", fmt.Errorf("prompt has zero length")
	}

	return prompt, nil
}
