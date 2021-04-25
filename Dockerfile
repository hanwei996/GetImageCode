FROM alpine
ADD GetImageCode /GetImageCode
ENTRYPOINT [ "/GetImageCode" ]
