CLI client for EdgeX services

```bash
$ edgectl device list
Name                    AdminState  OperatingState  Lables
GS1-AC-Drive01              UNLOCKED    ENABLED     [modbus industrial]
KMC.BAC-121036CE01          UNLOCKED    ENABLED     [thermostat industrial]
JC.RR5.NAE9.ConfRoom.Padre.Island01 UNLOCKED    ENABLED     [thermostat industrial]
```

```bash
$ edgectl addressable get edgex-support-scheduler
Name            Protocol    HTTPMethod  Address         Port
edgex-support-scheduler HTTP        POST        edgex-support-scheduler 48085
```
