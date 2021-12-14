FROM golang:1.16-alpine

WORKDIR /app

COPY . ./

RUN go build -o /gameoflife

ENV SIZE 20
ENV PORT=8080
EXPOSE $PORT
#EXPOSE 8080
# Run the app.  CMD is required to run on Heroku
# $PORT is set by Heroku
# CMD gunicorn --bind 0.0.0.0:$PORT wsgi
CMD ["sh", "-c", "/gameoflife -addr ${PORT} -size ${SIZE}"]