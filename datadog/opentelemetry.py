import time

from flask import Flask
from opentelemetry import trace
from opentelemetry.instrumentation.flask import FlaskInstrumentor
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.resources import Resource
from opentelemetry.sdk.trace.export import BatchSpanProcessor
from opentelemetry.exporter.otlp.proto.http.trace_exporter import OTLPSpanExporter

resource = Resource.create({"service.name": "tuan-service"})
trace.set_tracer_provider(TracerProvider(resource=resource))
tracer = trace.get_tracer(__name__)

exporter = OTLPSpanExporter(endpoint="http://0.0.0.0:4318/v1/traces")
trace.get_tracer_provider().add_span_processor(BatchSpanProcessor(exporter))

app = Flask(__name__)
FlaskInstrumentor().instrument_app(app)


@app.route("/")
def hello():
    with tracer.start_as_current_span("sleep-1s"):
        time.sleep(1)
        with tracer.start_as_current_span("sleep-0.5s"):
            time.sleep(0.5)

    with tracer.start_as_current_span("return-response"):
        return "Hello OpenTelemetry → Datadog!"
