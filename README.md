# wave-to-mqtt
- `cp variables.example.sh variables.sh`
- Edit the variables in `variables.sh` accordingly 
- `export GO111MODULE=on && go build`
- `sudo service bluetooth stop && sudo hciconfig hci0 down`
- `sudo setcap 'cap_net_raw,cap_net_admin=eip' wave-to-mqtt`
- `./wave-to-mqtt`
- Press the middle button on your Wave to connect

## MQTT Topics
| Topic         | Purpose                                 | Values                                         |
| --------------|:----------------------------------------|:----------------------------------------------:|
|wave/buttons/A | New state of the Top button             | Down, Long, ExtraLong, Up, LongUp, ExtraLongUp |
|wave/buttons/B | New state of the Middle button          | Up                                             |
|wave/buttons/C | New state of the Bottom button          | Down, Long, ExtraLong, Up, LongUp, ExtraLongUp |
|wave/euler/x   | X-axis Euler rotation detected while the middle button was held down | (Float)           |
|wave/euler/y   | Y-axis Euler rotation detected while the middle button was held down | (Float)           |
