#!/usr/bin/env python3
import logging

import requests
import time
import json
import sys


def fetch_statistics():
    try:
        url = 'http://server-app:80/statistics'

        resp = requests.get(url)

        resp.raise_for_status()

        return resp.json()
    except Exception as err:
        print('exception: ' + str(err), file=sys.stderr)
        return None


def main():
    while True:
        stats = fetch_statistics()

        if stats is not None:
            stats_json = json.dumps(stats)
            with open("statistics_log.txt", "a") as file:
                file.write(f"{stats_json}\n")

        time.sleep(5)


if __name__ == "__main__":
    main()