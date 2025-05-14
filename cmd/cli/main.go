package main

import (
	"flag"
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"time"
)

var mode string

func splayTime(sleep int) {
	rnd := rand.Intn(sleep)
	time.Sleep(time.Duration(rnd) * time.Second)
}

func main() {
	flag.StringVar(&mode, "mode", "plain", "Ausgabemodus: plain oder structured")
	flag.Parse()

	var logger *slog.Logger
	if mode == "structured" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}

	// Startup-Header
	if mode == "plain" {
		fmt.Println("Initiating Startup Sequence")
		fmt.Println("")
	} else {
		logger.Info("Initiating Startup Sequence")
	}

	//initializeSubspaceEmitterField
	outputStatus(logger, []string{
		"Powering up zero-point energy core",
		"Stabilizing quantum vacuum fluctuations",
		"Unfolding subspace emitter field",
		"Subspace emitter field fully unfolded",
		"Subspace emitter field is now active",
	})

	//initializeFasterThanLightCommunicationArray
	outputStatus(logger, []string{
		"Calibrating multidimensional antenna array",
		"Synchronizing phase with local spacetime",
		"Starting test sequence for subspace emitter field",
		"Test sequence for subspace emitter field completed successfully",
		"Subspace emitter field is now operational",
	})

	//initializeQuantumEntanglementProtocol
	outputStatus(logger, []string{
		"Preparing particle for quantum entanglement",
		"Particle is entering entanglement chamber",
		"Initiating quantum entanglement protocol",
		"Quantum entanglement established",
	})

	//initializeDarkMatterContainmentField
	outputStatus(logger, []string{
		"Opening dark matter containment field",
		"Dark matter containment field is now open",
		"Aligning hyperspace frequency bands",
	})

	//initializeTachyonUplink
	outputStatus(logger, []string{
		"Establishing tachyon uplink",
		"Tachyon uplink established",
		"Verifying neutrino stream integrity",
		"Neutrino stream integrity confirmed",
	})

	//initializeGravitonRelayNodes
	outputStatus(logger, []string{
		"Activating graviton relay nodes",
		"Graviton relay nodes online",
		"Finalizing interstellar handshake protocol",
	})

	// Abschlussmeldung
	if mode == "plain" {
		fmt.Println("Faster-than-light communication array is now online and ready")
	} else {
		logger.Info("Faster-than-light communication array is now online and ready")
	}
}

func outputStatus(logger *slog.Logger, field []string) {
	sleepValue := len(field)
	for _, output := range field {
		if mode == "plain" {
			fmt.Println(output)
			splayTime(sleepValue)
		} else {
			logger.Info(output)
			splayTime(sleepValue)
		}
	}
	fmt.Println("")
}
