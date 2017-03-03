package rete

import "container/list"

type NccNode struct {
	parent   IReteNode
	children *list.List
	items    *list.List
	partner  *NccPartnerNode
}

func (node NccNode) get_node_type() string {
	return NCC_NODE
}
func (node NccNode) get_parent() IReteNode {
	return node.parent
}
func (node NccNode) get_items() *list.List {
	return node.items
}
func (node NccNode) get_children() *list.List {
	return node.children
}
func (node NccNode) left_activation(t *Token, w *WME) {
	new_token := make_token(node, t, w)
	node.items.PushBack(new_token)

	new_token.ncc_results = list.New()
	buffer := node.partner.new_result_buffer
	for e := buffer.Front(); e != nil; e = e.Next() {
		result := e.Value.(*Token)
		result.owner = new_token
		new_token.ncc_results.PushBack(result)
		buffer.Remove(e)
	}
	if new_token.ncc_results.Len() > 0 {
		return
	}
	for e := node.children.Front(); e != nil; e = e.Next() {
		e.Value.(IReteNode).left_activation(new_token, nil)
	}
}
func (node NccNode) right_activation(w *WME) {
}
