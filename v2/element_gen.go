// This file is auto-generated by internal/cmd/gennodes/main.go. DO NOT EDIT

package helium

type Element struct {
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

func (*Element) Type() ElementType {
	return ElementNode
}

func (n *Element) Content() []byte {
	return n.content
}

func (n *Element) OwnerDocument() *Document {
	return n.doc
}

func (n *Element) SetOwnerDocument(v *Document) {
	n.doc = v
}

func (n *Element) FirstChild() Node {
	return n.firstChild
}

func (n *Element) setFirstChild(v Node) {
	n.firstChild = v
}

func (n *Element) LastChild() Node {
	return n.lastChild
}

func (n *Element) setLastChild(v Node) {
	n.lastChild = v
}

func (n *Element) LocalName() string {
	return n.name
}

func (n *Element) NextSibling() Node {
	return n.next
}

func (n *Element) SetNextSibling(v Node) {
	n.next = v
}

func (n *Element) Parent() Node {
	return n.parent
}

func (n *Element) SetParent(v Node) {
	n.parent = v
}

func (n *Element) PrevSibling() Node {
	return n.prev
}

func (n *Element) SetPrevSibling(v Node) {
	n.prev = v
}

func (n *Element) AddChild(c Node) error {
	return addChild(n, c)
}

func (n *Element) AddContent(b []byte) error {
	return addContent(n, b)
}

func (n *Element) AddSibling(c Node) error {
	return addSibling(n, c)
}

func (n *Element) Replace(v Node) {
	replaceNode(n, v)
}