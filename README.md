# Trace2Dia

Transform a JSON trace file from Jaeger to a graphical diagram

Current diagram implemenations
- PlantUML Sequence

# Usage

This command will print the generated PlantUML structure to stdout
```
t2d trace.json
``` 

### Example trace file

A example trace file can be found [here](https://raw.githubusercontent.com/jaegertracing/jaeger/refs/heads/main/plugin/storage/integration/fixtures/traces/example_trace.json)