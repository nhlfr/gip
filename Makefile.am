all:
	go build -o $(builddir)/gip ./cmd/gip.go

clean-local:
	rm -f $(builddir)/gip

install-exec-local:
	install $(builddir)/gip $(bindir)/gip
