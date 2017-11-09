# Upload

Then compile and run the service:

```bash
go build -o upload
./upload
```

Images can be uploaded by sending a multipart mime POST request to `/api/images`. 
The API returns the images metadata in the response including an image IDs that 
can be used to show the image medata and download it.

Using CURL as an example files can be uploaded with:

```
curl -XPOST http://localhost:8080/api/images -F"image1=@foo.png" -F"image2=@bar.png"
```

where `foo.png` and `bar.png` are files that live in the current directory.
