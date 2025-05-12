package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
)

var mode string

func main() {
	flag.StringVar(&mode, "mode", "plain", "Ausgabemodus: plain oder structured")
	flag.Parse()

	var logger *slog.Logger
	if mode == "structured" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}

	// Startup-Header
	if mode == "plain" {
		fmt.Println("Initiating Startup Sequence\n")
	} else {
		logger.Info("Initiating Startup Sequence")
	}

	// Initialisiere Subsysteme
	err := initializeSubspaceEmitterField(logger)
	if err != nil {
		if mode == "plain" {
			fmt.Println("Failed to initialize subspace emitter field:", err)
		} else {
			logger.Error("Failed to initialize subspace emitter field", "error", err)
		}
		os.Exit(1)
	}

	err = initializeFasterThanLightCommunicationArray(logger)
	if err != nil {
		if mode == "plain" {
			fmt.Println("Failed to initialize faster-than-light communication array:", err)
		} else {
			logger.Error("Failed to initialize faster-than-light communication array", "error", err)
		}
		os.Exit(1)
	}

	err = initializeQuantumEntanglementProtocol(logger)
	if err != nil {
		if mode == "plain" {
			fmt.Println("Failed to initialize quantum entanglement protocol:", err)
		} else {
			logger.Error("Failed to initialize quantum entanglement protocol", "error", err)
		}
		os.Exit(1)
	}

	err = initializeDarkMatterContainmentField(logger)
	if err != nil {
		if mode == "plain" {
			fmt.Println("Failed to initialize dark matter containment field:", err)
		} else {
			logger.Error("Failed to initialize dark matter containment field", "error", err)
		}
		os.Exit(1)
	}

	err = initializeTachyonUplink(logger)
	if err != nil {
		if mode == "plain" {
			fmt.Println("Failed to initialize tachyon uplink:", err)
		} else {
			logger.Error("Failed to initialize tachyon uplink", "error", err)
		}
		os.Exit(1)
	}

	err = initializeGravitonRelayNodes(logger)
	if err != nil {
		if mode == "plain" {
			fmt.Println("Failed to initialize graviton relay nodes:", err)
		} else {
			logger.Error("Failed to initialize graviton relay nodes", "error", err)
		}
		os.Exit(1)
	}

	// Abschlussmeldung
	if mode == "plain" {
		fmt.Println("Faster-than-light communication array is now online and ready")
	} else {
		logger.Info("Faster-than-light communication array is now online and ready")
	}
}

func initializeGravitonRelayNodes(logger *slog.Logger) error {
	if mode == "plain" {
		fmt.Println("Activating graviton relay nodes")
		fmt.Println("Graviton relay nodes online")
		fmt.Println("Finalizing interstellar handshake protocol\n")
	} else {
		logger.Info("Activating graviton relay nodes")
		logger.Info("Graviton relay nodes online")
		logger.Info("Finalizing interstellar handshake protocol")
	}
	return nil
}

func initializeTachyonUplink(logger *slog.Logger) error {
	if mode == "plain" {
		fmt.Println("Establishing tachyon uplink")
		fmt.Println("Tachyon uplink established")
		fmt.Println("Verifying neutrino stream integrity")
		fmt.Println("Neutrino stream integrity confirmed\n")
	} else {
		logger.Info("Establishing tachyon uplink")
		logger.Info("Tachyon uplink established")
		logger.Info("Verifying neutrino stream integrity")
		logger.Info("Neutrino stream integrity confirmed")
	}
	return nil
}

func initializeDarkMatterContainmentField(logger *slog.Logger) error {
	if mode == "plain" {
		fmt.Println("Opening dark matter containment field")
		fmt.Println("Dark matter containment field is now open")
		fmt.Println("Aligning hyperspace frequency bands\n")
	} else {
		logger.Info("Opening dark matter containment field")
		logger.Info("Dark matter containment field is now open")
		logger.Info("Aligning hyperspace frequency bands")
	}
	return nil
}

func initializeQuantumEntanglementProtocol(logger *slog.Logger) error {
	if mode == "plain" {
		fmt.Println("Preparing particle for quantum entanglement")
		fmt.Println("Particle is entering entanglement chamber")
		fmt.Println("Initiating quantum entanglement protocol")
		fmt.Println("Quantum entanglement established\n")
	} else {
		logger.Info("Preparing particle for quantum entanglement")
		logger.Info("Particle is entering entanglement chamber")
		logger.Info("Initiating quantum entanglement protocol")
		logger.Info("Quantum entanglement established")
	}
	return nil
}

func initializeFasterThanLightCommunicationArray(logger *slog.Logger) error {
	if mode == "plain" {
		fmt.Println("Calibrating multidimensional antenna array")
		fmt.Println("Synchronizing phase with local spacetime")
		fmt.Println("Starting test sequence for subspace emitter field")
		fmt.Println("Test sequence for subspace emitter field completed successfully")
		fmt.Println("Subspace emitter field is now operational\n")
	} else {
		logger.Info("Calibrating multidimensional antenna array")
		logger.Info("Synchronizing phase with local spacetime")
		logger.Info("Starting test sequence for subspace emitter field")
		logger.Info("Test sequence for subspace emitter field completed successfully")
		logger.Info("Subspace emitter field is now operational")
	}
	return nil
}

func initializeSubspaceEmitterField(logger *slog.Logger) error {
	if mode == "plain" {
		fmt.Println("Powering up zero-point energy core")
		fmt.Println("Stabilizing quantum vacuum fluctuations")
		fmt.Println("Unfolding subspace emitter field")
		fmt.Println("Subspace emitter field fully unfolded")
		fmt.Println("Subspace emitter field is now active\n")
	} else {
		logger.Info("Powering up zero-point energy core")
		logger.Info("Stabilizing quantum vacuum fluctuations")
		logger.Info("Unfolding subspace emitter field")
		logger.Info("Subspace emitter field fully unfolded")
		logger.Info("Subspace emitter field is now active")
	}
	return nil
}
