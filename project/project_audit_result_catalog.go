package project

type ProjectAuditResultCatalog interface {
	Add(project Project)
	Remove(project Project)
	Get(id int64) Project
}
