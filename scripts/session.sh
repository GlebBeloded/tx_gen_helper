txgen clear
set -e
NOW=`date +%s`
BEFORE=`expr $NOW - 1000`
AD_TIME=`expr $NOW - 100`

#Givin` dat fat cash to tha aggregata`
printf '12345678\n' | melcli tx melodia aggregate $(melcli keys show jack -a) 100mfmc --from jack -y -b block

#generate random sets of bytes 
txgen generate-bytes A B C D E F G 
# generate our valid distribute rewards
txgen distribute-rewards  A 10mfmc 2 A $(melcli keys show jack -a) A $(melcli keys show alice -a) B 15mfmc 1 B $(melcli keys show gusgus -a) > master_unsigned.json
#generate empty distribute rewards for out test
txgen distribute-rewards > dr_empty.json

# our sessions
# jack session
txgen register-session $(melcli keys show jack -a) $BEFORE $NOW melodia A $AD_TIME > jack_unsigned.json
printf '12345678\n12345678\n' | melcli tx sign jack_unsigned.json --from jack > jack_signed.json
melcli tx broadcast jack_signed.json -b block
#alice session
txgen register-session $(melcli keys show alice -a) $BEFORE $NOW melodia A $AD_TIME > alice_unsigned.json
printf '12345678\n12345678\n' | melcli tx sign alice_unsigned.json --from alice > alice_signed.json
melcli tx broadcast alice_signed.json -b block
#gusgus session
txgen register-session $(melcli keys show gusgus -a) $BEFORE $NOW melodia B $AD_TIME > gusgus_unsigned.json
printf '12345678\n12345678\n' | melcli tx sign gusgus_unsigned.json --from gusgus > gusgus_signed.json
melcli tx broadcast gusgus_signed.json -b block



printf '12345678\n12345678\n' | melcli tx sign dr_empty.json --from master > dr_empty_signed.json
melcli tx broadcast dr_empty_signed.json -b block
printf '12345678\n12345678\n' | melcli tx sign master_unsigned.json --from master > master_signed.json
melcli tx broadcast master_signed.json -b block
