FROM python:3.8 AS build-env

COPY . /app
WORKDIR /app

RUN pip install --upgrade pip && pip install pipenv
RUN pipenv lock -r | pip install -r /dev/stdin


FROM python:3.8

COPY --from=build-env /usr/local/lib/python3.8/site-packages /usr/local/lib/python3.8/site-packages
COPY --from=build-env /app/cryptarithm /app/cryptarithm
COPY setup.py /app/

WORKDIR /app
ENTRYPOINT ["python", "cryptarithm/main.py"]
