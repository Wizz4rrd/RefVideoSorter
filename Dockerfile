FROM python:3.10.7

WORKDIR /usr/src/app

COPY . .

CMD [ "python3", "./app.py" ]