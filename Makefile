export GOPATH=$(PWD)

GO=go
JAVAC=javac
jrePath=/home/ztex/jdk1.8.0_221/

jvmgo: src/jvmgo
	$(GO) install $@

search_class: bin/jvmgo
	$< -Xjre $(jrePath)/jre java.lang.Object
	$< -Xjre $(jrePath)/jre tests.HelloWorld.HelloWorld

java/Test/GaussTest.class: java/Test/GaussTest.java
	$(JAVAC) $<

Gauss_Test: bin/jvmgo java/Test/GaussTest.class
	$< -Xjre $(jrePath)/jre java.Test.GaussTest

.PHONY: clean
clean:
	$(RM) bin/*
