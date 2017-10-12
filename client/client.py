#!/usr/bin/python

import requests, sys
import json
from struct import pack

HOST    = "http://localhost:9001"
HEADERS = {'content-type': 'application/json','Host':'grpc'}
USAGE   = """

envoy-python-client usage:
  ./client.py set <key> <value> - sets the <key> and <value>
  ./client.py get <key>         - gets the value for <key>
  ./client.py count             - gets count of key-value multiple times
  """

class KVClient():

    def get(self, key):
        resp = requests.get(HOST + "/kv/" + key, headers=HEADERS)
        return json.loads(resp.content)


    def set(self, key, value):
        req = {"value": value}
        data = json.dumps(req)
        return requests.put(HOST + "/kv/" + key, data=data, headers=HEADERS)

    def count(self):
        resp = requests.get(HOST + "/count", headers=HEADERS)
        return json.loads(resp.content)

def run():
  if len(sys.argv) == 1:
    print(USAGE)

    sys.exit(0)

  cmd = sys.argv[1]

  client = KVClient()

  if cmd == "get":
    # ensure a key was provided
    if len(sys.argv) != 3:
      print(USAGE)
      sys.exit(1)

    # get the key to fetch
    key = sys.argv[2]

    # send the request to the server
    response = client.get(key)

    print(response['value'])
    sys.exit(0)

  elif cmd == "set":
    # ensure a key and value were provided
    if len(sys.argv) < 4:
      print(USAGE)
      sys.exit(1)

    # get the key and the full text of value
    key = sys.argv[2]
    value = " ".join(sys.argv[3:])

    # send the request to the server
    response = client.set(key, value)

    print("setf %s to %s" % (key, value))

  elif cmd == "count":
    # send the request to the server
    response = client.count()

    print(response)
    sys.exit(0)

if __name__ == '__main__':
  run()
