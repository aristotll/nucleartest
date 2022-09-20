package quickstart

import "testing"

func TestServer(t *testing.T) {
	ServerRun()

}

func TestClient(t *testing.T) {
	ClientRun()
}

func TestClientByAsync(t *testing.T) {
	ClientByAsync()
}
