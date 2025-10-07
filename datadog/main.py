import logging
import sys
import time

from ddtrace import tracer, config, patch_all
from flask import Flask

app = Flask(__name__)
patch_all(logging=True)

logger = logging.getLogger("tuan-service")
logger.setLevel(logging.INFO)

handler = logging.StreamHandler(sys.stdout)
formatter = logging.Formatter(
    '%(asctime)s | %(levelname)s | %(message)s | dd.trace_id=%(dd.trace_id)s dd.span_id=%(dd.span_id)s'
)
handler.setFormatter(formatter)
logger.addHandler(handler)

config.service = "tuan-service-ddtrace"
config.env = "dev"
config.version = "1.0.0"


@app.route("/")
def hello():
    logger.info("start hello")
    with tracer.trace("sleep-1s"):
        logger.info("sleep 1s")
        time.sleep(1)
        with tracer.trace("sleep-0.5s"):
            logger.info("sleep-0.5s")
            time.sleep(0.5)

    with tracer.trace("return-response"):
        logger.info("resturn response")
        return "Hello OpenTelemetry → Datadog!"


if __name__ == '__main__':
    app.run("0.0.0.0", 3000)
