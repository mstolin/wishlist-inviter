# Requirements Builder
FROM python:3.10 as requirements-stage
LABEL stage=builder
WORKDIR /tmp
RUN pip install poetry
COPY ./pyproject.toml ./poetry.lock* /tmp/
RUN poetry export -f requirements.txt --output requirements.txt --without-hashes

# Runner
FROM python:3.10
WORKDIR /code
COPY --from=requirements-stage /tmp/requirements.txt /code/requirements.txt
RUN pip install --no-cache-dir --upgrade -r /code/requirements.txt
COPY amazon_adapter /code/amazon_adapter
EXPOSE 8080
ENTRYPOINT ["uvicorn", "amazon_adapter.main:app"]
