#/bin/bash
awk '{
    for(i=1;i<=NF;i++){
        res[$i]+=1}
    } END {
    for(k in res){ 
        print k" "res[k]
    }
}' ./testdata/192.txt | sort -nr -k2


# run: sh ./192.wordsCount.sh
