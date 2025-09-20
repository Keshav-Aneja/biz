package constants

type permissionSet struct {
	DIRECTORY int
	FILE      int
}

var Permissions = permissionSet{
	DIRECTORY: 0750,
	FILE:      0660,
}

type directorySet struct {
	TEMPORARY		string
	BIZ_MODULES		string	
	PACKAGE_FILE	string	
}

var Directories = directorySet{
	TEMPORARY: "tmp",
	BIZ_MODULES: "biz_modules",
	PACKAGE_FILE: "package.json",
}
