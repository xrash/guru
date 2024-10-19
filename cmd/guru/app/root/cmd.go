package root

import (
	"os"

	"github.com/spf13/cobra"
)

type cliopts struct {
	model  string
	stdin  bool
	prompt string
	url    string
}

type rootCommand struct {
	options cliopts
}

func (c *rootCommand) run(cmd *cobra.Command, args []string) {
	os.Exit(run(c.options, args))
}

func CreateCmd() *cobra.Command {
	c := &rootCommand{}

	cmd := &cobra.Command{
		Use:   "root",
		Short: "Use ollama models query local code",
		Long:  `Use ollama models query local code`,
		Run:   c.run,
	}

	cmd.Flags().BoolVarP(
		&c.options.stdin,
		"stdin",
		"",
		false,
		"pass prompt through stdin",
	)

	cmd.Flags().StringVarP(
		&c.options.prompt,
		"prompt",
		"",
		"",
		"pass prompt through this option",
	)

	cmd.Flags().StringVarP(
		&c.options.url,
		"url",
		"",
		"http://localhost:11434/api/generate",
		"url of ollama api",
	)

	cmd.Flags().StringVarP(
		&c.options.model,
		"model",
		"",
		"codellama:7b",
		"name of the ollama model to call",
	)

	return cmd
}
