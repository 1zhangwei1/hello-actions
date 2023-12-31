FROM python:3.9.5-slim

RUN pip install flask redis && \
    apt-get update && \
    apt-get install -y curl && \
    groupadd -r flask && useradd -r -g flask flask && \
    mkdir /src && \
    chown -R flask:flask /src

USER flask

COPY app.py /src/app.py

WORKDIR /src

ENV FLASK=app.py REDIS_HOST=redis

EXPOSE 5000

HEALTHCHECK --interval=30s --timeout=3s \
    CMD curl -f http://localhost:5080 || exit 1


CMD ["flask","run","-h","0.0.0.0"]