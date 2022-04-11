### To launch a local privatenet with 4 validators ###

First follow the steps [here](https://docs.ckm8token.org/docs/setup) to compile the latest code of the `privatenet` branch. Next, run the commands below to launch the privatenet with 4 validators:

```
cd $ckm8_HOME
cp -r ./integration/testnet ../testnet
mkdir ../testnet/logs

# In terminal 1
ckm8 start --config=../testnet/node1 |& tee ../testnet/logs/node1.log

# In terminal 2
ckm8 start --config=../testnet/node2 |& tee ../testnet/logs/node2.log

# In terminal 3
ckm8 start --config=../testnet/node3 |& tee ../testnet/logs/node3.log

# In terminal 4
ckm8 start --config=../testnet/node4 |& tee ../testnet/logs/node4.log
```
