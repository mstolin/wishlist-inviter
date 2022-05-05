FROM python:3.10-slim-bullseye

COPY pyproject.toml poetry.lock /app/
COPY amazon_scrapper /app/amazon_scrapper/

RUN pip install -U poetry

WORKDIR /app
RUN poetry config virtualenvs.create false
RUN poetry install --no-dev

EXPOSE 8080

ENTRYPOINT uvicorn amazon_scrapper.main:app --host="${HOST:-0.0.0.0}" --port=8080
