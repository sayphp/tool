#!/bin/bash
#set -x
#操作选项
act(){
	if [ "$1" == "start" ];then
		echo 1
		return 1
	elif [ "$1" == "stop" ];then
		echo 2
		return 2
	elif [ "$1" == "switch" ];then
		echo 3
		return 3
	elif [ "$1" == "help" ];then
		echo 4
		return 4
	else
		echo 0
		return 0
	fi
}

#参数检查
check(){
	code=`act $1`
	if [ $code == 0 ];then
		echo "请输入正确的操作选项，输入查看帮助"
		return 1
	fi
	if [[ "$2" =~ "" ]] && [[ -f $2 ]];then
		echo $2
	else
		file="${filepath}/ss.conf"
		if [ -f $file ];then
			echo "$file"
		else
			echo "请输入有效的配置文件地址,输入查看帮助"
			return 2
		fi
	fi
	return 0
}

#获取ss代理服务配置并启动
conf(){
	source $1
	f=""
	for file in ${ss_dir}/*
	do
		n=$[ $RANDOM%2 ]
		if [ $n -gt 0 ];then
			f=${file}
			break
		fi
	done
	if [ "$f" = "" ];then
		f=$file
	fi
	echo "$f"
	return 1
}

#启动
t(){
	pid=`c`
	if [ "$pid" == "" ];then
		`ss-local -c $1 &> /dev/null &`
		echo "ss成功开启"
		return 1;
	else
		echo "ss已启动"
		return 0;
	fi
}

#检查
c(){
	echo `ps aux | grep 'ss-local' | grep -v 'grep' | awk '{print $2}'`
	return 1;
}
#关闭
p(){
	pid=`c`
	if [ "$pid" != "" ];then
		`kill -9 $pid`
		echo 'ss成功关闭'
	else
		echo 'ss未启动'
	fi
	return 1
}

#切换
#暂缓开发
s(){
	echo '暂时不支持自动切换'
	#`p`
	#`t`
}

#帮助
h(){
	echo 'ss帮助文档'
	echo '启动ss服务: ./ss start'
	echo '停止ss服务：./ss stop'
	echo '====================='
	echo 'Say 20180111'
}
