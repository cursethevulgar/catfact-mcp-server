.PHONY: build image

NAME := catfact-mcp-server

build:
        go build -o $(NAME) .

clean:
        powershell -Command "if (Test-Path '$(NAME)') { Remove-Item '$(NAME)' }"
        powershell -Command "if (Test-Path '$(NAME).exe') { Remove-Item '$(NAME).exe' }"
        powershell -Command "if (Test-Path '$(NAME).exe~') { Remove-Item '$(NAME).exe~' }"

image:
        docker build -t $(NAME) .
