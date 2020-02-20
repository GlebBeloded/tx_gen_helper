txgen clear

set -e

NOW=`date +%s`
BEFORE=`expr $NOW - 1000`
AFTER=`expr $NOW + 1000`

#Get tokens
printf '12345678\n' | melcli tx melodia aggregate $(melcli keys show jack -a) 10mfmc --from jack -y -b block

# create poll
txgen register-poll test $BEFORE $AFTER 10mfmc 1  > poll_unsigned.json
printf '12345678\n12345678\n' |melcli tx sign poll_unsigned.json --from master > poll_signed.json
melcli tx broadcast poll_signed.json --broadcast-mode block

# submit poll
txgen submit-poll test $NOW $(melcli keys show jack -a) 1> jack_poll_unsigned.json
printf '12345678\n12345678\n' |melcli tx sign jack_poll_unsigned.json --from jack > jack_poll_signed.json
melcli tx broadcast jack_poll_signed.json --broadcast-mode block