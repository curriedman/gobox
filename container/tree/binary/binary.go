package binary

import "github.com/recursivecurry/gobox/typeclass/ord"

type Interface interface {
	ord.Interface
	Left() Interface
	SetLeft(node Interface)
	Right() Interface
	SetRight(node Interface)
}

func Add(root *Interface, node Interface) {
	var parent Interface
	current := *root

	for {
		if current.Eq(node) {
			if parent == nil {
				*root = current
			} else {
				if parent.Less(node) {
					parent.SetRight(node)
				} else {
					parent.SetLeft(node)
				}
			}
			node.SetLeft(current.Left())
			node.SetRight(current.Right())
			return
		} else if current.Less(node) {
			if current.Right() == nil {
				current.SetRight(node)
				return
			}
			parent = current
			current = current.Right()
		} else {
			if current.Left() == nil {
				current.SetLeft(node)
				return
			}
			parent = current
			current = current.Left()
		}
	}
}

func Remove(root *Interface, node ord.Interface) (Interface, bool) {
	return find(root, node, true)
}

func Get(root *Interface, node ord.Interface) (Interface, bool) {
	return find(root, node, false)
}

func find(root *Interface, node ord.Interface, remove bool) (Interface, bool) {
	var parent Interface
	var found Interface
	current := *root

	for {
		if current == nil {
			return nil, false
		}

		if current.Eq(node) {
			found = current
			if remove {
				if parent.Less(current) {
					if current.Left() == nil {
						parent.SetRight(current.Right())
					} else if current.Right() == nil {
						parent.SetRight(current.Left())
					} else {
						minimum := current.Right()
						for minimum.Left() != nil {
							minimum = minimum.Left()
						}
						minimum, _ = find(&current, minimum, true)
						minimum.SetRight(current.Right())
						minimum.SetLeft(current.Left())
						parent.SetRight(minimum)
					}
				} else {
					if current.Left() == nil {
						parent.SetLeft(current.Right())
					} else if current.Right() == nil {
						parent.SetLeft(current.Left())
					} else {
						minimum := current.Right()
						for minimum.Left() != nil {
							minimum = minimum.Left()
						}
						minimum, _ = find(&current, minimum, true)
						minimum.SetRight(current.Right())
						minimum.SetLeft(current.Left())
						parent.SetLeft(minimum)
					}
				}
			}
			break
		} else if current.Less(node) {
			parent = current
			current = current.Right()
		} else {
			parent = current
			current = current.Left()
		}
	}

	return found, true
}
func removeNode(parent Interface, right bool) Interface {
	var found Interface
	var current Interface
	if right {
		current = parent.Right()
	} else {
		current = parent.Left()
	}
	for {
		if current.Right() == nil {
			found = current.Left()
			break
		} else if current.Left() == nil {
			found = current.Right()
			break
		} else {
			minimum := current.Right()
			for minimum.Left() != nil {
				minimum = minimum.Left()
			}
			found = minimum
			find(current, remove)
		}
	}

	if parent.Less(found) {
		parent.SetRight(found)
	} else {
		parent.SetLeft(found)
	}
	return current
}
