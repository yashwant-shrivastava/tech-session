# tech-onboarding

#### Build the image
```
docker build --no-cache --tag hello-app .
```

#### Search for the image
```
docker images
```

#### Run the container
```
docker run -p "10000:10000" -it hello-app
```

#### Test
```
curl 0.0.0.0:10000/hello
```
