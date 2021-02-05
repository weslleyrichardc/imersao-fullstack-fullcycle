/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

	kafkaProducer "github.com/weslleyrichardc/imersao/desafio-01/application/kafka/producer"
	kafkaConsumer "github.com/weslleyrichardc/imersao/desafio-01/application/kafka/consumer"
)

// kafkaCmd represents the kafka command
var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Comunicação Kafka",
	Run: func(cmd *cobra.Command, args []string) {


		fmt.Println("Comunicação iniciada...")

		deliverChan := make(chan ckafka.Event)
		producer := kafkaProducer.NewKafkaProducer()

		kafkaProducer.Publish("Mensagem Teste", "Teste", producer, deliverChan)
		go kafkaProducer.DeliveryReport(deliverChan)

		consumer := kafkaConsumer.NewKafkaConsumer(producer, deliverChan)
		consumer.Consume()
	},
}

func init() {
	rootCmd.AddCommand(kafkaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kafkaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kafkaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
