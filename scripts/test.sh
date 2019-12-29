txgen clear
#CREATE POLL
rm ./*.json
txgen register-poll stest1 100 2577571765 4 > poll_unsigned.json
printf '12345678\n' |melcli tx sign poll_unsigned.json --from master > poll_signed.json
melcli tx broadcast poll_signed.json --broadcast-mode block
#JACK POLL-SUBMISSION
txgen submit-poll stest1 200 $(melcli keys show jack -a) 1> jack_poll_unsigned.json
printf '12345678\n' |melcli tx sign jack_poll_unsigned.json --from jack > jack_poll_signed.json
melcli tx broadcast jack_poll_signed.json --broadcast-mode block
# CREATE DISTRIBUTE-REWARDS TX WITH AD_BYTES
txgen distribute-rewards test > master_unsigned.json
#JACK AD SESSION
txgen register-session $(melcli keys show jack -a) 127 1027 melodia test > jack_unsigned.json
printf '12345678\n' |melcli tx sign jack_unsigned.json --from jack > jack_signed.json
melcli tx broadcast jack_signed.json
#EMPTY SWAP
printf '12345678\n'| melcli tx melodia distribute-rewards --from master -y --broadcast-mode block
#ALICE SESSION (assuming bug is here)
txgen register-session $(melcli keys show alice -a) 127 1027 melodia test > alice_unsigned.json
printf '12345678\n' |melcli tx sign alice_unsigned.json --from alice > alice_signed.json
melcli tx broadcast alice_signed.json
#SWAP WITH TEST AD_BYTES
printf '12345678\n' |melcli tx sign master_unsigned.json --from master > master_signed.json
melcli tx broadcast master_signed.json --broadcast-mode block

ALICECOINS=`melcli q account $(melcli keys show alice -a) | jq '.value.coins[1]' | tr -d '"'`
echo Alice coins:$ALICECOINS 
JACKCOINS=`melcli q account $(melcli keys show jack -a) | jq '.value.coins[1]' | tr -d '"'`
echo Jack coins:$JACKCOINS
