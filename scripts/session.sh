txgen clear
set -e
NOW=`date +%s`
BEFORE=`expr $NOW - 1000`
AD_TIME=`expr $NOW - 100`

#Givin` dat fat cash to tha aggregata`
printf '12345678\n' | melcli tx melodia aggregate $(melcli keys show jack -a) 100mfmc --from jack -y -b block

txgen generate-bytes  test
txgen distribute-rewards  100mfmc test > master_unsigned.json
txgen register-session $(melcli keys show jack -a) $BEFORE $NOW melodia test $AD_TIME > jack_unsigned.json
txgen register-session $(melcli keys show alice -a) $BEFORE $NOW melodia test $AD_TIME > alice_unsigned.json
printf '12345678\n12345678\n' |melcli tx sign jack_unsigned.json --from jack > jack_signed.json
printf '12345678\n12345678\n' |melcli tx sign alice_unsigned.json --from alice > alice_signed.json
printf '12345678\n12345678\n' |melcli tx sign master_unsigned.json --from master > master_signed.json
melcli tx broadcast jack_signed.json -b block
melcli tx broadcast alice_signed.json -b block
melcli tx broadcast master_signed.json -b block
sleep 5
printf '12345678\n12345678\n' |melcli tx sign master_unsigned.json --from master > master_signed.json
melcli tx broadcast master_signed.json --broadcast-mode block