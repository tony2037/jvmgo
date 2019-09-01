export GOPATH=$(PWD)

GO=go
jrePath=/home/ztex/jdk1.8.0_221/

jvmgo: src/jvmgo
	$(GO) install $@

search_class: bin/jvmgo
	$< -Xjre $(jrePath)/jre java.lang.Object
	$< -Xjre $(jrePath)/jre tests.HelloWorld.HelloWorld

.PHONY: clean
clean:
	$(RM) bin/*
