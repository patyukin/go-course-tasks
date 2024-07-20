package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	permissions = 0o644
)

func main() {
	rootCmd := &cobra.Command{Use: "filetool"}

	createCmd := &cobra.Command{
		Use:   "create [filename]",
		Short: "Create a new file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := os.WriteFile(args[0], []byte("Example file content"), permissions)
			if err != nil {
				fmt.Println("Error creating file:", err)
			} else {
				fmt.Println("File created:", args[0])
			}
		},
	}

	readCmd := &cobra.Command{
		Use:   "read [filename]",
		Short: "Read file contents",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			data, err := os.ReadFile(args[0])
			if err != nil {
				fmt.Println("Error reading file:", err)
			} else {
				fmt.Println("File contents of", args[0], ":\n", string(data))
			}
		},
	}

	deleteCmd := &cobra.Command{
		Use:   "delete [filename]",
		Short: "Delete a file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := os.Remove(args[0])
			if err != nil {
				fmt.Println("Error deleting file:", err)
			} else {
				fmt.Println("File deleted:", args[0])
			}
		},
	}

	rootCmd.AddCommand(createCmd, readCmd, deleteCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
