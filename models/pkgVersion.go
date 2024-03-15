package models

type PkgVersion struct {
	Pkg        Pkg
	Major      int32
	Minor      int32
	Patch      int32
	Identifier *string
	Metadata   *Meta
	FileID     string
}

type PkgVersions []PkgVersion

func (v PkgVersions) Len() int {
	return len(v)
}

func (v PkgVersions) Less(i, j int) bool {
	if v[i].Major != v[j].Major {
		return v[i].Major < v[i].Major
	}
	if v[i].Minor != v[j].Minor {
		return v[i].Minor < v[i].Minor
	}
	if v[i].Patch != v[j].Patch {
		return v[i].Patch < v[i].Patch
	}
	if v[i].Identifier == nil {
		return true
	}
	if v[j].Identifier == nil {
		return false
	}
	return *v[i].Identifier < *v[j].Identifier
}

func (v PkgVersions) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
