package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cobra_simple_client_server",
	Short: "Example of simple client-server application using Cobra",
}

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Client for the simple client-server application",
	Run: func(cmd *cobra.Command, args []string) {
		err := client()
		if err != nil {
			fmt.Printf("Error running client: %v\n", err)
		}
	},
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server for the simple client-server application",
	Run: func(cmd *cobra.Command, args []string) {
		err := server()
		if err != nil {
			fmt.Printf("Error starting server: %v\n", err)
		}
	},
}

func main() {
	rootCmd.AddCommand(clientCmd)
	rootCmd.AddCommand(serverCmd)
	rootCmd.Execute()
}

func server() error {
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
		fmt.Printf("RemoteAddr=%s Path=%s\n", r.RemoteAddr, r.URL.Path)
	})
	fmt.Printf("Listening on 0.0.0.0:8000 see http://127.0.0.1:8000 and http://127.0.0.1:8000/api/hello\n")
	return http.ListenAndServe(":8000", nil)
}

func client() error {
	body, err := httpGet("http://127.0.0.1:8000/api/hello")
	if err != nil {
		return err
	}
	fmt.Print(body)
	return nil
}

func httpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
