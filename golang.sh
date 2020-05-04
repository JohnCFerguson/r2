#Set workign directory
WORKINGDIR=$1
echo "$WORKINGDIR"
#Set package name
PKGNAME=$2

cd "$WORKDINGDIR"
go mod init $PKGNAME
