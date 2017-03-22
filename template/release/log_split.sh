#!/bin/sh

#按下面格式启动
#nohup python3 -u ./write_test.py 2>&1 | bash ./log_split.sh &

max_lines=100000	#这里设定分割行数

log_folder="./nohup_logs/"
counter=1
current_file=''

if test -e ${log_folder}
then
	echo "${log_folder} exist"
else
	echo "${log_folder} not exist"
	mkdir ${log_folder}
	echo "mkdir '${log_folder}'"
fi

cat /dev/stdin | while read LINE
	do
	if [ ${counter} -le 1 ]	
	then
		current_file=${log_folder}$(date +%Y-%m-%d_%H:%M:%S).log
		echo ${LINE} >> ${current_file}
		((counter++))
	elif [ ${counter} -ge ${max_lines} ]
	then
		echo ${LINE} >> ${current_file}
		counter=1
	else
		echo ${LINE} >> ${current_file}
		((counter++))
	fi
	done
