all:
	go build -tags=jsoniter .

clean:
	rm -f authentication gin.log
