package main

import (
	"fmt"
	"time"

	"github.com/asynkron/protoactor-go/actor"
)

func main() {
	// Initialize the actor system
	context := actor.NewActorSystem()

	// Create the Reddit Engine actor
	engineProps := actor.PropsFromProducer(func() actor.Actor {
		return NewRedditEngineActor()
	})
	enginePID := context.Root.Spawn(engineProps)

	// Ensure enginePID is valid
	if enginePID == nil {
		panic("Failed to spawn Reddit Engine actor!")
	}

	// Create the Simulator actor
	simulatorProps := actor.PropsFromProducer(func() actor.Actor {
		return NewSimulatorActor(enginePID, 100) // Adjust user count as needed
	})
	simulatorPID := context.Root.Spawn(simulatorProps)

	// Ensure simulatorPID is valid
	if simulatorPID == nil {
		panic("Failed to spawn Simulator actor!")
	}

	// Run the simulation
	fmt.Println("Simulation started. Waiting for completion...")

	// Wait dynamically for the simulation to run
	time.Sleep(time.Second * 3) // Replace this with dynamic coordination if possible

	context.Root.Stop(enginePID)

	// Shutdown the actor system after actors are stopped
	context.Shutdown()

	fmt.Println("\nSimulation ended.")
}
