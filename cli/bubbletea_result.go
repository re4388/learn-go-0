package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/bitfield/script"
	tea "github.com/charmbracelet/bubbletea"
)

const correctConfigPath = "/Users/re4388/.config/gcloud/configurations/config_default"

/*
Here, I already setup two path in the dir xxxx-personal and xxx-wemo
So, I simple copy to a tmp file and mv to the correct file
*/
func copyAndRename(configPath string) {
	// copy config_default_personal to config_default_personal-tmp
	copy_script := fmt.Sprintf("cp %s %s-tmp", configPath, configPath)
	// fmt.Println(cp_script) // for debug
	script.Exec(copy_script).Wait()

	// rename to config_default_personal-tmo to config_default
	rename_script := fmt.Sprintf("mv %s-tmp %s", configPath, correctConfigPath)
	// fmt.Println(cp_script)
	script.Exec(rename_script).Wait()

}

var choices = []string{"wemo", "personal"}

type ResultModel struct {
	cursorIdx int
	choice    string
}

func (m ResultModel) Init() tea.Cmd {
	return nil
}

func (model ResultModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return model, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			model.choice = choices[model.cursorIdx]
			return model, tea.Quit

		case "down", "j":
			model.cursorIdx++
			if model.cursorIdx >= len(choices) {
				model.cursorIdx = 0
			}

		case "up", "k":
			model.cursorIdx--
			if model.cursorIdx < 0 {
				model.cursorIdx = len(choices) - 1
			}
		}
	}

	return model, nil
}

func (model ResultModel) View() string {
	str := strings.Builder{}
	str.WriteString("Which GCP Account to use?\n\n")

	for idx := 0; idx < len(choices); idx++ {
		if model.cursorIdx == idx {
			str.WriteString("(•) ")
		} else {
			str.WriteString("( ) ")
		}
		str.WriteString(choices[idx])
		str.WriteString("\n")
	}
	str.WriteString("\n(press q to quit)\n")

	return str.String()
}

func RUN_bubbleTea_result() {
	program := tea.NewProgram(ResultModel{})

	// Run returns the model as a tea.Model.
	m, err := program.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(ResultModel); ok && m.choice != "" {
		fmt.Printf("\n---\nYou chose %s!\n", m.choice)

		switch m.choice {
		case "wemo":
			fmt.Printf("gcp a/c switch to %s...\n", m.choice)
			copyAndRename("/Users/re4388/.config/gcloud/configurations/config_default_wemo")
		case "personal":
			fmt.Printf("gcp a/c switch to %s...\n", m.choice)
			copyAndRename("/Users/re4388/.config/gcloud/configurations/config_default_personal")

		default:
			fmt.Printf("no this env, choose `wemo` or `personal`")
		}
	}

	fmt.Printf("done\n")
}
