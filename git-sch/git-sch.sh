Y=$(date +%Y)
M=$(date +%m)
D=$(date +%d)
Ym="$Y-$M"
Ymd=$Y-$M-$D
GitRep="algorithm"

HomeDir="/Users/bsit/go/src/study"
GitDir="$HomeDir/$GitRep"
FileDir="$HomeDir/$GitRep/programmers"
FileName="al$Ymd"

mkdir -p $FileDir

echo "# ymd 프로그래머스" >> $FileDir/$FileName

cd $GitDir
git add .
git commit -m "# $Ymd 프로그래머스 " 
git push a 75
