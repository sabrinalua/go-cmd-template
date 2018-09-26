REL = release/
REL_W = release/win
REL_L = release/lin

init:
	mkdir $(REL) $(REL_W) $(REL_L)
	go install web
	mv bin/web $(REL_L)
	GOOS=windows go install web
	mv bin/windows_amd64/* $(REL_W)
clean:
	rm -rf $(REL) bin/*
