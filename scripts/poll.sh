txgen clear

NOW=`date +%s`
BEFORE=`expr $NOW - 1000`
AFTER=`expr $NOW + 1000`

txgen register-poll stest $BEFORE $AFTER 6 > poll_unsigned.json
printf '12345678\n12345678\n' |melcli tx sign poll_unsigned.json --from master > poll_signed.json
melcli tx broadcast poll_signed.json --broadcast-mode block
txgen submit-poll stest $NOW $(melcli keys show jack -a) 1> jack_poll_unsigned.json
printf '12345678\n12345678\n' |melcli tx sign jack_poll_unsigned.json --from jack > jack_poll_signed.json
melcli tx broadcast jack_poll_signed.json --broadcast-mode block
txgen distribute-rewards > master_unsigned.json
printf '12345678\n12345678\n' |melcli tx sign master_unsigned.json --from master > master_signed.json
melcli tx broadcast master_signed.json --broadcast-mode block
sleep 5
printf '12345678\n12345678\n' |melcli tx sign master_unsigned.json --from master > master_signed.json
melcli tx broadcast master_signed.json --broadcast-mode block