# Serve videos from directory in browser

How to use:

```shell
docker build -t webplayer .

cd /videos/directory
docker run --rm --name webplayer -d -p 777:3000 -v "$(pwd):/app/videos:ro" webplayer
```
