#!/usr/bin/env python3

import requests
import time
import json


def fetch_statistics():
    try:
        url = 'http://localhost:12121/statistics'

        resp = requests.get(url)

        resp.raise_for_status()

        return resp.json()
    except requests.RequestException as e:
        print(e)
        return None


def main():
    while True:
        stats = fetch_statistics()

        if stats is not None:
            # Convert dictionary to a JSON string
            stats_json = json.dumps(stats)

            # Print the result to a file
            with open("statistics_log.txt", "a") as file:
                file.write(f"{stats_json}\n")

        time.sleep(5)


if __name__ == "__main__":
    main()