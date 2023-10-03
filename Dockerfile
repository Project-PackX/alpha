# Get Go image from DockerHub.
FROM golang AS api

# Set working directory.
WORKDIR /compiler

# Copy dependency locks so we can cache.
COPY go.mod go.sum ./

# Get all of our dependencies.
RUN go mod download

# Copy all of our remaining application.
COPY . .

# Build our application.
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./main.go

# Use 'scratch' image for super-mini build.
FROM scratch AS prod

LABEL authors="Dominik Szilágyi <dominik.szilagyi@gmail.com>,Zsombor Töreky <toreky.zsombor@gmail.com>"
LABEL org.opencontainers.image.authors="Károly Szakály <karoly.szakaly2000@gmail.com>"

# Set working directory for this stage.
WORKDIR /production

# Copy our compiled executable from the last stage.
COPY --from=api /compiler/app .

# Run application and expose port 8080.
EXPOSE 8080
CMD ["./app"]