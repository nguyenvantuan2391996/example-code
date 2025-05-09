import requests
import json
import pandas as pd
import os

from apscheduler.schedulers.blocking import BlockingScheduler
from datetime import datetime
from dotenv import load_dotenv
from pytz import timezone
from vnstock import Vnstock

load_dotenv()


def send_slack(code, current, minimum, avg, maximum):
    if current < avg:
        action = "buy"
    else:
        action = "sell"

    payload = {
        "text": code,
        "icon_emoji": ":bar_chart:",
        "attachments": [
            {
                "image_url": "https://example.com/status.png",
                "fields": [
                    {
                        "title": "Date",
                        "value": os.getenv('START') + " - " + str(datetime.now().date()),
                        "short": True
                    },
                    {
                        "title": "Action",
                        "value": "You may " + action,
                        "short": True
                    },
                    {
                        "title": "Current Price",
                        "value": current,
                        "short": True
                    },
                    {
                        "title": "Minimum Price",
                        "value": minimum,
                        "short": True
                    },
                    {
                        "title": "Avg Price",
                        "value": avg,
                        "short": True
                    },
                    {
                        "title": "Maximum Price",
                        "value": maximum,
                        "short": True
                    },
                ]
            }
        ]
    }

    # Send the POST request
    response = requests.post(os.getenv('WEBHOOK_URL'), data=json.dumps(payload),
                             headers={'Content-Type': 'application/json'})

    # Optional: check response
    if response.status_code != 200:
        raise ValueError(f'Request failed: {response.status_code}, {response.text}')


def job():
    for code in os.getenv('CODES').split(","):
        stock = Vnstock().stock(symbol=code, source='VCI')

        data = stock.quote.history(start=os.getenv('START'), end=str(datetime.now().date()), interval='1D')

        df = pd.DataFrame(data)

        cur_close = float(df['close'].tail(1))
        min_close = float(df['close'].min())
        max_close = float(df['close'].max())
        avg_close = float(df['close'].mean())

        send_slack(code, cur_close, min_close, avg_close, max_close)


if __name__ == '__main__':
    print(f"CODES={os.getenv('CODES')}")
    print(f"WEBHOOK_URL={os.getenv('WEBHOOK_URL')}")
    print(f"START={os.getenv('START')}")
    print(f"HOUR={os.getenv('HOUR')}")
    print(f"MINUTE={os.getenv('MINUTE')}")

    scheduler = BlockingScheduler(timezone=timezone('Asia/Ho_Chi_Minh'))
    scheduler.add_job(job, 'cron', hour=os.getenv('HOUR'), minute=os.getenv('MINUTE'), day_of_week='mon-fri')

    scheduler.start()


