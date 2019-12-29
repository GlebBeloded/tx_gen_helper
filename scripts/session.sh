txgen clear
txgen distribute-rewards test > master_unsigned.json
txgen register-session $(melcli keys show jack -a) 127 1027 melodia test > jack_unsigned.json
txgen register-session $(melcli keys show alice -a) 127 1027 melodia test > alice_unsigned.json
printf '12345678\n12345678\n' |melcli tx sign jack_unsigned.json --from jack > jack_signed.json
printf '12345678\n12345678\n' |melcli tx sign alice_unsigned.json --from alice > alice_signed.json
printf '12345678\n12345678\n' |melcli tx sign master_unsigned.json --from master > master_signed.json
melcli tx broadcast jack_signed.json
melcli tx broadcast alice_signed.json
melcli tx broadcast master_signed.json --broadcast-mode block
sleep 5
printf '12345678\n12345678\n' |melcli tx sign master_unsigned.json --from master > master_signed.json
melcli tx broadcast master_signed.json --broadcast-mode block