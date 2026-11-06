thrift=thrift
idl_path = idl/stream-center

gen_service:
	kitex -module github.com/kouleen/stream-center -service stream.service  $(idl_path)/stream.thrift
