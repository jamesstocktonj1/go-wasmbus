package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/jamesstocktonj1/go-wasmbus"
)

func main() {
	bus, err := wasmbus.New(&wasmbus.Config{
		NatsUrl: "samba-share:4222",
	})
	if err != nil {
		log.Fatal(err)
	}

	bus.OnComponentScaled(func(evt wasmbus.BaseEvent, data wasmbus.ComponentScaled) {
		slog.Info("handle OnComponentScaled", "source", evt.Source, "data", data)
	})
	bus.OnComponentScaledFailed(func(evt wasmbus.BaseEvent, data wasmbus.ComponentScaledFailed) {
		slog.Info("handle OnComponentScaledFailed", "source", evt.Source, "data", data)
	})
	bus.OnLinkdefSet(func(evt wasmbus.BaseEvent, data wasmbus.LinkdefSet) {
		slog.Info("handle OnLinkdefSet", "source", evt.Source, "data", data)
	})
	bus.OnLinkdefDeleted(func(evt wasmbus.BaseEvent, data wasmbus.LinkdefDeleted) {
		slog.Info("handle OnLinkdefDeleted", "source", evt.Source, "data", data)
	})
	bus.OnProviderStarted(func(evt wasmbus.BaseEvent, data wasmbus.ProviderStarted) {
		slog.Info("handle OnProviderStarted", "source", evt.Source, "data", data)
	})
	bus.OnProviderStartFailed(func(evt wasmbus.BaseEvent, data wasmbus.ProviderStartFailed) {
		slog.Info("handle OnProviderStartFailed", "source", evt.Source, "data", data)
	})
	bus.OnProviderStopped(func(evt wasmbus.BaseEvent, data wasmbus.ProviderStopped) {
		slog.Info("handle OnProviderStopped", "source", evt.Source, "data", data)
	})
	bus.OnHealthCheckPassed(func(evt wasmbus.BaseEvent, data wasmbus.HealthCheckPassed) {
		slog.Info("handle OnHealthCheckPassed", "source", evt.Source, "data", data)
	})
	bus.OnHealthCheckFailed(func(evt wasmbus.BaseEvent, data wasmbus.HealthCheckFailed) {
		slog.Info("handle OnHealthCheckFailed", "source", evt.Source, "data", data)
	})
	bus.OnHealthCheckStatus(func(evt wasmbus.BaseEvent, data wasmbus.HealthCheckStatus) {
		slog.Info("handle OnHealthCheckStatus", "source", evt.Source, "data", data)
	})
	bus.OnConfigSet(func(evt wasmbus.BaseEvent, data wasmbus.ConfigSet) {
		slog.Info("handle OnConfigSet", "source", evt.Source, "data", data)
	})
	bus.OnConfigDeleted(func(evt wasmbus.BaseEvent, data wasmbus.ConfigDeleted) {
		slog.Info("handle OnConfigDeleted", "source", evt.Source, "data", data)
	})
	bus.OnLabelsChanged(func(evt wasmbus.BaseEvent, data wasmbus.LabelsChanged) {
		slog.Info("handle OnLabelsChanged", "source", evt.Source, "data", data)
	})
	bus.OnHostHeartbeat(func(evt wasmbus.BaseEvent, data wasmbus.HostHeartbeat) {
		slog.Info("handle OnHostHeartbeat", "source", evt.Source, "data", data)
	})
	bus.OnHostStarted(func(evt wasmbus.BaseEvent, data wasmbus.HostStarted) {
		slog.Info("handle OnHostStarted", "source", evt.Source, "data", data)
	})
	bus.OnHostStopped(func(evt wasmbus.BaseEvent, data wasmbus.HostStopped) {
		slog.Info("handle OnHostStopped", "source", evt.Source, "data", data)
	})

	err = bus.Start()
	if err != nil {
		log.Fatal(err)
	}
	defer bus.Close()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT)

	<-signalCh
}
