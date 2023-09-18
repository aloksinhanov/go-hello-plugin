plugin-for-ide: 
	go build -buildmode=plugin -gcflags="all=-N -l" -o plugins/contactdevice/contactdevice.so plugins/contactdevice/association.go
	go build -buildmode=plugin -gcflags="all=-N -l" -o plugins/contacttask/contacttask.so plugins/contacttask/association.go

plugin:
	go build -buildmode=plugin -o plugins/contactdevice/contactdevice.so plugins/contactdevice/association.go
	go build -buildmode=plugin -o plugins/contacttask/contacttask.so plugins/contacttask/association.go
