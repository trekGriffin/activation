#!/bin/bash
echo -e "\ninput your des"
read des
echo -e "\ninput your cmd"
read cmd
echo -e "\ninput your workdir"
read workdir

echo -e "\nyour setting is"
echo -e "\ndes:" $des 
echo -e "\ncmd:" $cmd 
echo -e "\nworkdir:" $workdir

echo -e "\nis this right?[y|n]"
read str
if [[ "$str" == y* ]]; then 
echo "confirmed"
else 
echo "exiting..."
exit
fi

echo "
[Unit]
Description=$des

After=syslog.target network-online.target

[Service]
ExecStart=$cmd

WorkingDirectory=$workdir



Restart=always


[Install]
WantedBy=multi-user.target
"  > /etc/systemd/system/$des.service
systemctl daemon-reload
systemctl start $des
systemctl status $des
