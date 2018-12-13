function checkBinary {
	local search=$1
	local version=$2

	binary=$(which $search 2>/dev/null)
	if [ "$binary" == "" ]; then
		echo "Missing $search, please install it (minimum version: $version)"

		exit 1
	fi
}
