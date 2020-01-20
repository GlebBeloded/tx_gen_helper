txgen clear
rm *.json
NOW=`date +%s`
printf '12345678\n'| melcli tx melodia distribute-rewards --from master -y --broadcast-mode block
BEFORE=`expr $NOW - 1000`
AFTER=`expr $NOW + 10`
#we need to sleep so our tx can happen after distribute rewards, and not at the same time
sleep 30
txgen generate-bytes ad_one ad_two
txgen register-session $(melcli keys show jack -a) `expr $BEFORE - 1` `expr $AFTER + 1` melodia ad_one $BEFORE ad_two $AFTER > jack_unsigned.json
printf '12345678\n' | melcli tx sign jack_unsigned.json --from jack > jack_signed.json
melcli tx broadcast jack_signed.json -b block