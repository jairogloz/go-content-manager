# Step 1: Usar una imagen oficial de Go runtime como base
FROM golang:1.21-alpine

# Step 2: Establecer el directorio de trabajo en /app
WORKDIR /app

# Step 3: Copiar el contenido del directorio actual al directorio de trabajo (que es /app)
COPY . .

# Step 4: Build the Go app
# Cambiamos el directorio de trabajo a /app/cmd/http/gin
WORKDIR /app/cmd/http/gin
# Descargamos las dependencias
RUN go mod tidy
# Compilamos el binario y el resultado lo guardamos en /app/myapp
RUN go build -o /app/myapp

# Step 5: Establecer el comando predeterminado para ejecutar el binario
# Cambiamos el directorio de trabajo a /app nuevamente
WORKDIR /app
# Ejecutamos el binario
CMD ["./myapp"]
