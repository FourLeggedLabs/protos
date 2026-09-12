# Censor Protos

Protobuf schemas for the Censor agent and API.

## Packages

- `censor.agent.v1` — `AgentPolicy`, `AgentEvent`, `AgentLogUpload` / `JobContext`

Wire format over HTTP is **protojson** (JSON). Audit logs are NDJSON lines of `AgentEvent`; uploads gzip the log file.

## Develop

```bash
task          # lint + generate + test
task lint
task generate
task test
```

Requires [buf](https://buf.build) and [Task](https://taskfile.dev).
