plugin-for-ide:
	cd ./plugins/contactdevice && echo $pwd && go mod tidy && go build -buildmode=plugin -gcflags="all=-N -l" -o contactdevice.so association.go && cd ../../
	
	cd ./plugins/contacttask && echo $pwd && go mod tidy && go build -buildmode=plugin -gcflags="all=-N -l" -o contacttask.so association.go && cd ../../

plugin:
	cd ./plugins/contactdevice && echo $pwd && go mod tidy && go build -buildmode=plugin -o contactdevice.so association.go && cd ../../

	cd ./plugins/contacttask && echo $pwd && go mod tidy && go build -buildmode=plugin -o contacttask.so association.go && cd ../../