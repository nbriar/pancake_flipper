# Empty the total.cov file if it exists
[ -e 'total.cov' ] && >| 'total.cov'
# Iterate through each package and run tests with coverage enabled
echo "mode: count" >> total.cov
threshold=${COVERAGE_THRESHOLD:-70}

test_errors=0

for pkg in $(go list ./... | grep -v '/vendor/');
do
    dir="$GOPATH/src/$pkg"
    len="${#PWD}"
    dir_relative=".${dir:$len}"

    # Run test coverage on each package
    go test -short -covermode=count -coverprofile="$dir_relative/total.tmp" "$dir_relative"
		if [ $? -gt 0 ]
		then
			test_errors=1
		fi

    if [ -f "$dir_relative/total.tmp" ]
    then
        # Add coverage results to a temp file and cat them into our total.cov
        # for later analysis
        cat "$dir_relative/total.tmp" | tail -n +2 >> total.cov
        rm "$dir_relative/total.tmp"
    fi
done


if [ $test_errors -gt 0 ]
then
	echo "Tests failed.  Skipping coverage."
	exit 1
fi


# Checking the coverage
echo "Analyzing test coverage"
coverage=`go tool cover -func total.cov | grep total | awk '{print $3}'`
echo "Done analyzing"
coverage_int=`echo $coverage | awk -F\. '{print $1}'`
echo "Total Test Coverage: $coverage"
echo "Target Coverage Threshold: $threshold%"

if (( $coverage_int < $threshold ))
then
    echo "Error: Test coverage of $coverage is below threshold of $threshold%!"
    exit 1
else
    echo "Success: Test coverage of $coverage is above threshold of $threshold%!"
    exit 0
fi

