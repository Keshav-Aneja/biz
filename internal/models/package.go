package models

type PkgInfo struct {
	ID       	string 	 						`json:"_id"`
	Name     	string	 						`json:"name"`
	DistTags	DistTags 						`json:"dist-tags"`
	Description	string							`json:"description"`
	Versions	map[string]PkgVersionInfo		`json:"versions"`
}

type DistTags struct {
	// For now let's stay on the latest version, we can include canary, beta, rc, etc version lateron
	Latest		string		`json:"latest"`
}

type PkgVersionInfo struct {
	ID			string 		`json:"_id"`
	Name		string		`json:"name"`
	Version		string		`json:"version"`
	Dist		struct {
		Shasum	string		`json:"shasum"`
		Tarball	string		`json:"tarball"`
		Integrity string	`json:"integrity"`
		// Ignoring the signatures for now, we can handle authenticity later on
	}	
	NPMVersion	string		`json:"_npmVersion"`
	NodeVersion	string		`json:"_nodeVersion"`
	Description string		`json:"description"`
}
