cd ./hello
export GO111MODULE=on
export GOCACHE=
rm -rf ./vendor
go mod vendor
rm -rf ./vendor/github.com/kubeless
zip -r ../hello.zip ./
cd ..

ARGS="$(cat dev.env |sed -e "s/^/--env /g") --runtime go1.10 \
      --from-file hello.zip \
      --handler main.Handler"

echo $ARGS

if ! kubeless function update hello-go $ARGS; then
    kubeless function deploy hello-go $ARGS
fi
