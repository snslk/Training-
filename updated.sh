#!/bin/bash

echo -e "\n----List Of Running Processes----\n"
ps r o pid,user,group,%mem,%cpu,stat # list of running process
echo -e "\n--------------------------------\n"
A=1
while [ $A != 0 ]
do
echo -e "\n1 : Display detail of a process" 
echo 2 : Terminate a process 
echo 3 : Exit
A=1
echo -e "\n enter choice :"
read choice 
err='err'
case $choice in

	1 ) echo enter PID
		read pid
		echo ----------------------------
		ps o pid,user,group,%mem,%cpu $pid 
		echo ----------------------------
		err=$(echo $?)
		if [ "$err" != 0 ]
		then
		echo *********** Invalid PID ************
		fi
		;;
	2 ) echo enter PID
		read pid	
		echo -----------------------------
		kill $pid
		err=$(echo $?)
		if [ "$err" != 0 ]
		then
		echo *********** Invalid PID / Permission ************
		fi
		;;
	3 ) echo exit 
		A=0 
		;;
	* ) echo enter correct option 
esac
done
