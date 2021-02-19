// This file is auto-generated by internal/cmd/gennodes/main.go. DO NOT EDIT

package helium

type Comment struct {
	content    []byte
	doc        *Document
	firstChild Node
	lastChild  Node
	name       string
	next       Node
	ns         *Namespace
	nsDefs     []*Namespace
	parent     Node
	prev       Node
	private    interface{}
	properties *Attribute
}

func (*Comment) Type() ElementType {
	return CommentNode
}

func (n *Comment) Content() []byte {
	return n.content
}

func (n *Comment) OwnerDocument() *Document {
	return n.doc
}

func (n *Comment) SetOwnerDocument(v *Document) {
	n.doc = v
}

func (n *Comment) FirstChild() Node {
	return n.firstChild
}

func (n *Comment) setFirstChild(v Node) {
	n.firstChild = v
}

func (n *Comment) LastChild() Node {
	return n.lastChild
}

func (n *Comment) setLastChild(v Node) {
	n.lastChild = v
}

func (n *Comment) LocalName() string {
	return n.name
}

func (n *Comment) NextSibling() Node {
	return n.next
}

func (n *Comment) SetNextSibling(v Node) {
	n.next = v
}

func (n *Comment) Parent() Node {
	return n.parent
}

func (n *Comment) SetParent(v Node) {
	n.parent = v
}

func (n *Comment) PrevSibling() Node {
	return n.prev
}

func (n *Comment) SetPrevSibling(v Node) {
	n.prev = v
}

func (n *Comment) AddChild(c Node) error {
	return addChild(n, c)
}

func (n *Comment) AddContent(b []byte) error {
	return addContent(n, b)
}

func (n *Comment) AddSibling(c Node) error {
	return addSibling(n, c)
}

func (n *Comment) Replace(v Node) {
	replaceNode(n, v)
}
