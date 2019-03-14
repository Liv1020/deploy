# deploy
A tool of deploy your code via git web hook

#

## Usage

### git
```
git clone https://github.com/Liv1020/deploy.git
cd deploy
go build

./deploy

curl http://127.0.0.1:4321/?path=/your-deploy-path
```

### build
your-deploy-path add build.sh

### deploy
your-deploy-path add deploy.sh

### view log
```
curl http://127.0.0.1:4321/log
```