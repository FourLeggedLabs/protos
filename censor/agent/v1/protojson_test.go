package agentv1_test

import (
	"strings"
	"testing"
	"time"

	agentv1 "github.com/FourLeggedLabs/protos/gen/go/censor/agent/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestAgentEventProtoJSONRoundTrip(t *testing.T) {
	in := &agentv1.AgentEvent{
		Ts:      timestamppb.New(time.Date(2026, 9, 12, 8, 0, 0, 123456789, time.UTC)),
		Type:    agentv1.EventType_EVENT_TYPE_CONNECT,
		Action:  agentv1.Action_ACTION_MONITOR_DENY,
		Pid:     123,
		Comm:    "curl",
		Dst:     "1.2.3.4",
		DstPort: 443,
		Proto:   agentv1.Protocol_PROTOCOL_TCP,
		Host:    "example.com",
		Rule:    "allowedHosts",
	}

	b, err := protojson.Marshal(in)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	line := string(b)
	if !strings.Contains(line, `"type":"EVENT_TYPE_CONNECT"`) {
		t.Fatalf("expected enum name in JSON, got %s", line)
	}

	out := &agentv1.AgentEvent{}
	if err := protojson.Unmarshal(b, out); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if out.GetPid() != 123 || out.GetHost() != "example.com" || out.GetAction() != agentv1.Action_ACTION_MONITOR_DENY {
		t.Fatalf("round-trip mismatch: %+v", out)
	}
}

func TestAgentPolicyProtoJSONRoundTrip(t *testing.T) {
	in := &agentv1.AgentPolicy{
		CorrelationId:         "corr-1",
		Mode:                  agentv1.Mode_MODE_ENFORCE,
		AllowedHosts:          []string{"github.com", "**.npmjs.org"},
		DeniedHosts:           []string{"evil.example"},
		WatchSudo:             true,
		DisableSudo:           false,
		AutoAllowGithubHosts:  true,
	}
	b, err := protojson.Marshal(in)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	out := &agentv1.AgentPolicy{}
	if err := protojson.Unmarshal(b, out); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if out.GetMode() != agentv1.Mode_MODE_ENFORCE || len(out.GetAllowedHosts()) != 2 {
		t.Fatalf("round-trip mismatch: %+v", out)
	}
}

func TestAgentLogUploadWithMatrix(t *testing.T) {
	matrix, err := structpb.NewStruct(map[string]any{"os": "ubuntu-latest", "go": "1.24"})
	if err != nil {
		t.Fatalf("struct: %v", err)
	}
	in := &agentv1.AgentLogUpload{
		CorrelationId: "corr-1",
		Mode:          agentv1.Mode_MODE_MONITOR,
		AgentVersion:  "0.1.0",
		Status:        "success",
		Job: &agentv1.JobContext{
			RepoOwner:  "FourLeggedLabs",
			RepoName:   "censor",
			RunId:      "123",
			RunAttempt: "1",
			Job:        "build",
			JobId:      "build:abcd1234efgh",
			Matrix:     matrix,
		},
	}
	b, err := protojson.Marshal(in)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	out := &agentv1.AgentLogUpload{}
	if err := protojson.Unmarshal(b, out); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if out.GetJob().GetJobId() != "build:abcd1234efgh" {
		t.Fatalf("job_id: %q", out.GetJob().GetJobId())
	}
	if out.GetJob().GetMatrix().GetFields()["os"].GetStringValue() != "ubuntu-latest" {
		t.Fatalf("matrix: %+v", out.GetJob().GetMatrix())
	}
}
