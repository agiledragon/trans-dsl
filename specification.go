package transdsl

type Specification interface {
	Ok(transInfo *TransInfo) bool
}
