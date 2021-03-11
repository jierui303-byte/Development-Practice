#!/bin/bash

#执行ｃ file
function actionCFile()
{
	fileLists=`cd $1 && ls *.c`

	arrSize=()
	arrFileName=()
	
	maxSize=($(
		
		for var in $fileLists
		do
			fileSize=`wc -c $var | cut -d " " -f 1`
			fileName=`wc -c $var | cut -d " " -f 2`
			#fileSize=`wc -c $var`
			arrSize=(${arrSize[@]} $fileSize)
			#arrFileName=(${arrFileName[@]} $fileName)
			echo $fileSize
			#echo $fileName
		done | sort -n
	))

	echo ${maxSize[@]}
	echo ${maxSize[ ${#maxSize[@]}-1 ]}
	#最大文件
	maxFileSize=${maxSize[ ${#maxSize[@]}-1 ]}
	
	for var in $fileLists
	do
		#fileSize=`wc -c $var | cut -d " " -f 1`
		fileName=`wc -c $var | cut -d " " -f 2`
		fileSize=`wc -c $var`
		arrSize=(${arrSize[@]} $fileSize)
		#arrFileName=(${arrFileName[@]} $fileName)
		#echo $fileSize
		#echo $fileName
	done


	#echo ${arrSize[@]}
	echo ${arrFileName[@]}
	
	#文件大小排序
	for ((i=0; i<${#arrSize[@]}; i++))
	do
		if [ ${arrSize[i]} == $maxFileSize ]
		then
			#奇数
			echo "${arrSize[i]} ---$maxFileSize ==${arrSize[i+1]}"
		fi

		#echo "---${arrSize[i]}"

	done
}



function actionShFile()
{
	fileLists=`cd $1 && ls *.sh`

	arrSize=()
	arrFileName=()
	
	maxSize=($(
		
		for var in $fileLists
		do
			fileSize=`wc -c $var | cut -d " " -f 1`
			fileName=`wc -c $var | cut -d " " -f 2`
			#fileSize=`wc -c $var`
			arrSize=(${arrSize[@]} $fileSize)
			#arrFileName=(${arrFileName[@]} $fileName)
			echo $fileSize
			#echo $fileName
		done | sort -n
	))

	echo ${maxSize[@]}
	echo ${maxSize[ ${#maxSize[@]}-1 ]}
	#最大文件
	maxFileSize=${maxSize[ ${#maxSize[@]}-1 ]}
	
	for var in $fileLists
	do
		#fileSize=`wc -c $var | cut -d " " -f 1`
		fileName=`wc -c $var | cut -d " " -f 2`
		fileSize=`wc -c $var`
		arrSize=(${arrSize[@]} $fileSize)
		#arrFileName=(${arrFileName[@]} $fileName)
		#echo $fileSize
		#echo $fileName
	done


	#echo ${arrSize[@]}
	echo ${arrFileName[@]}
	
	#文件大小排序
	for ((i=0; i<${#arrSize[@]}; i++))
	do
		if [ ${arrSize[i]} == $maxFileSize ]
		then
			#奇数
			echo "${arrSize[i]} ---$maxFileSize ==${arrSize[i+1]}"
		fi

		#echo "---${arrSize[i]}"

	done
}




function actionCFileNew()
{
	fileLists=`cd $1 && ls *.$2`
	echo ${fileLists[@]}

	arrSize=()
	arrFileName=()
	
	maxSize=($(
		
		for var in $fileLists
		do
			fileSize=`wc -c $var | cut -d " " -f 1`
			fileName=`wc -c $var | cut -d " " -f 2`
			#fileSize=`wc -c $var`
			arrSize=(${arrSize[@]} $fileSize)
			#arrFileName=(${arrFileName[@]} $fileName)
			echo $fileSize
			#echo $fileName
		done | sort -n
	))

	echo ${maxSize[@]}
	echo ${maxSize[ ${#maxSize[@]}-1 ]}
	#最大文件
	maxFileSize=${maxSize[ ${#maxSize[@]}-1 ]}
	
	for var in $fileLists
	do
		#fileSize=`wc -c $var | cut -d " " -f 1`
		fileName=`wc -c $var | cut -d " " -f 2`
		fileSize=`wc -c $var`
		arrSize=(${arrSize[@]} $fileSize)
		#arrFileName=(${arrFileName[@]} $fileName)
		#echo $fileSize
		#echo $fileName
	done


	#echo ${arrSize[@]}
	echo ${arrFileName[@]}
	
	#文件大小排序
	for ((i=0; i<${#arrSize[@]}; i++))
	do
		if [ ${arrSize[i]} == $maxFileSize ]
		then
			#奇数
			echo "${arrSize[i]} ---$maxFileSize ==${arrSize[i+1]}"

			#执行文件
			if []
			then
			
			fi

		fi

		#echo "---${arrSize[i]}"

	done
}


read -p "输入一个文件路径：" filePath
echo $filePath

fileList=`cd $filePath && ls *.c *.sh`
#echo "搜索结果如下："
#echo "$fileList"

#调用ｃ函数
#actionCFile $filePath
#actionShFile $filePath

#另外一种方式
actionCFileNew $filePath c
actionCFileNew $filePath sh
