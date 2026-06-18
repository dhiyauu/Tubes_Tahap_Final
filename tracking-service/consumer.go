package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func StartConsumer() {

	if rabbitChannel == nil {
		log.Fatal("rabbitChannel masih nil")
	}

	msgs, err := rabbitChannel.Consume(
		"tracking_queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	go func() {

		for msg := range msgs {

			fmt.Println("===================================")
			fmt.Println("TRACKING SERVICE MENERIMA MESSAGE")
			fmt.Println(string(msg.Body))
			fmt.Println("===================================")

			// Decode JSON
			var event TrackingEvent

			err := json.Unmarshal(msg.Body, &event)
			if err != nil {
				log.Println("Gagal decode JSON:", err)
				continue
			}

			// Tampilkan isi event
			log.Printf(
				"Resi=%s | Lokasi=%s | Event=%s | Timestamp=%s",
				event.Resi,
				event.Lokasi,
				event.Event,
				event.Timestamp,
			)

			// Simpan ke database
			err = trackingRepo.Insert(event)

			if err != nil {
				log.Println("Gagal simpan ke MySQL:", err)
				continue
			}

			log.Println("Tracking berhasil disimpan ke database")

			// Kirim ke Report Service
			err = PublishReportEvent(event)

			if err != nil {
				log.Println("Gagal publish ke report_queue:", err)
			} else {
				log.Println("Berhasil publish ke report_queue")
			}
		}

	}()

}