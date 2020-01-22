FROM iron/go

WORKDIR /app

ADD main /app/
ADD bin/FMF-Menu-Parser-75f1f836c4b1.json /app/

ENTRYPOINT ["./main"]
