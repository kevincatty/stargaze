set -ex
set -o pipefail
DENOM=ustars
CHAINID=stargaze
RLYKEY=stars12g0xe2ld0k5ws3h7lmxc39d4rpl3fyxp5qys69
starsd version --long
apk add -U --no-cache jq tree curl wget
STARGAZE_HOME=/stargaze/starsd
curl -s -v http://stargaze:8090/kill || echo "done"
sleep 10

cat $STARGAZE_HOME/config/app.toml | grep -A 10  grpc
cat $STARGAZE_HOME/config/app.toml | grep -A 10  api
starsd start --pruning nothing --home $STARGAZE_HOME --grpc.address 0.0.0.0:9090 --rpc.laddr tcp://0.0.0.0:26657
