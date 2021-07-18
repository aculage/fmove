NAME = fmove
SRCS = main.go 

$(NAME): $(SRCS) ./internal/get/get.go ./internal/send/send.go
	go build -o $(NAME) $(SRCS)