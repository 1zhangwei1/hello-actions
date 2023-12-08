from flask import Flask
from redis import Redis
import os
import socket
app=Flask(__name__)
redis=Redis(host=os.environ.get("REDIS_HOST","127.0.0.1"),port=6379)

@app.route("/")
def hello():
    redis.incr("hits")
    return "Hello World! I have been seen %s times." % redis.get("hits")