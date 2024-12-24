FROM golang:1.21.1-alpine as builder

ARG HEADER_FILE=$HEADER_FILE
ARG ENV_FILE=$ENV_FILE

RUN echo "File swagger: $HEADER_FILE"
RUN echo "File env: $ENV_FILE"

RUN apk add bash ca-certificates git gcc g++ libc-dev
# Here we copy the rest of the source code
RUN mkdir -p /projects/financial_report/financial_report_be
WORKDIR /projects/financial_report/financial_report_be

# We want to populate the module cache based on the go.{mod,sum} files. 
COPY go.mod .
COPY go.sum .
# RUN ls -la /projects/phenikaa_intern/phenikaa_intern_be

RUN go mod download

# COPY $HEADER_FILE /projects/dbcl-pdt-backend/$HEADER_FILE
COPY . /projects/phenikaa_intern/financial_report_be
COPY .env /projects/phenikaa_intern/financial_report_be/.env

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o financial_report_be .
CMD ["./financial_report_be"]