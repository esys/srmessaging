package cmd

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/esys/srmessaging/internal/messaging"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringP("endpoint", "e", ":8080", "HTTP endpoint to listen to")
	runCmd.Flags().StringP("broker", "b", "localhost:9092", "Message broker endpoint")
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the server",
	Run: func(cmd *cobra.Command, args []string) {
		endpoint, _ := cmd.Flags().GetString("endpoint")
		broker, _ := cmd.Flags().GetString("broker")

		store := messaging.NewStore()

		consumer := messaging.NewConsumer(store, broker)
		defer consumer.Close()
		go consumer.ReadMessages()

		handler := messaging.NewWsHandler(store, broker)
		http.HandleFunc("/ws", handler.Handle)
		log.Printf("server started at %s\n", endpoint)

		log.Fatal(http.ListenAndServe(endpoint, nil))
	},
}
