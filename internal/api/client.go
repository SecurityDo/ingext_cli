package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
)

// IngextAppAPI defines the contract for interacting with the backend
type IngextAppAPI interface {
	Init(cluster, namespace, kubeCtx string) error
	Call(functionName string, functionArgs json.RawMessage) error

	// Domain specific methods
	AddStreamSource() error
	AddStreamSink() error
	// ... add other methods here
}

// Client is the concrete implementation of IngextAppAPI
type Client struct {
	Logger    *slog.Logger
	Cluster   string
	Namespace string
	// Add k8s clients, rest clients, etc here

	// Embed the K8s helper
	k8s *K8sClusterClient
}

// Option 1: Constructor injection (Recommended)
func NewClient(logger *slog.Logger) *Client {
	// Fallback: If caller passes nil, use a "No-Op" or Default logger
	// so the code doesn't panic.
	if logger == nil {
		// Using os.Stderr by default is safe for libraries
		logger = slog.New(slog.NewTextHandler(os.Stderr, nil))
	}
	return &Client{
		Logger: logger,
		k8s:    NewK8sClient(), // Initialize the helper
	}
}

// Ensure Client implements the interface
var _ IngextAppAPI = (*Client)(nil)

func (c *Client) Init(cluster, namespace string, kubeContext string) error {
	c.Cluster = cluster
	c.Namespace = namespace

	c.Logger.Debug("initializing k8s client", "context", kubeContext)

	// 1. Connect to Kubernetes
	if err := c.k8s.Connect(kubeContext); err != nil {
		return err
	}
	c.Logger.Info("connected to kubernetes cluster", "context", kubeContext)

	// TODO: Perform actual login / connection logic here
	fmt.Printf("DEBUG: Connecting to cluster '%s' in namespace '%s'...\n", cluster, namespace)
	return nil
}

func (c *Client) Call(functionName string, functionArgs json.RawMessage) error {
	// TODO: Implementation
	return nil
}

func (c *Client) AddStreamSource() error {
	// TODO: Implementation
	//fmt.Println("Action: Adding Stream Source via API...")

	// Use structured logging
	c.Logger.Info("adding stream source",
		"status", "pending",
		"retry", 1,
	)

	// Warn or Error
	//c.Logger.Debug("validating payload size...")
	return nil
}

func (c *Client) AddStreamSink() error {
	return nil
}
