package main

import "fmt"

type Node struct{
	element int
	left *Node
	right *Node
}

type BinaryTree struct{
	root *Node
}

func (b *BinaryTree) MakeTree(val int){
	n := &Node{element: val,left: nil,right: nil}
	if b.root==nil{
		b.root = n
		return
	}
	b.root.insert(val)
}
func (n *Node) insert(data int){
	if n == nil {
        return
    } else if data <= n.element {
        if n.left == nil {
            n.left = &Node{element: data, left: nil, right: nil}
        } else {
            n.left.insert(data)
        }
    } else {
        if n.right == nil {
            n.right = &Node{element: data, left: nil, right: nil}
        } else {
            n.right.insert(data)
        }
    }   

}
func (B *BinaryTree) Preorder(troot *Node){
	
	if troot!=nil{
		fmt.Printf("%d ", troot.element)
		B.Preorder(troot.left)
		B.Preorder(troot.right)
	}
}
func (b *BinaryTree) Inorder(troot *Node){
	if troot!=nil{
		b.Preorder(troot.left)
		fmt.Printf("%d ", troot.element)
		b.Preorder(troot.right)
	}
}
func (b *BinaryTree) Postorder(troot *Node){
	if troot!=nil{
		b.Preorder(troot.left)
		b.Preorder(troot.right)
		fmt.Printf("%d ", troot.element)
	}
}
func (b *BinaryTree) Count(troot *Node) int{
	if troot!=nil{
		x := b.Count(troot.left)
		y := b.Count(troot.right)
		return x+y+1
	}
	return 0
}
func (b *BinaryTree) Height(troot *Node) int{
	if troot!=nil{
		x := b.Height(troot.left)
		y := b.Height(troot.right)
		if x>y{
			return x+1
		}else{
			return y+1
		}
	}
	return -1
}
func (b *BinaryTree) Search(troot *Node, key int){
	if troot!=nil{
		if key==troot.element{
			fmt.Println("The key is there..!")
		}else if(key<troot.element){
			b.Search(troot.left,key)
		}else if(key>troot.element){
			b.Search(troot.right, key)
		}
	}else{
		fmt.Println("The key is not there..!")
	}
}

func main(){
	a := BinaryTree{}
	a.MakeTree(100)	
	a.MakeTree(240)
	a.MakeTree(30)
	a.MakeTree(40)
	a.MakeTree(550)
	a.MakeTree(60)
	a.MakeTree(710)
	fmt.Println("Preorder ---->")
	a.Preorder(a.root)
	fmt.Println()
	fmt.Println("Preorder ---->")
	a.Inorder(a.root)
	fmt.Println()
	fmt.Println("Postorder ---->")
	a.Postorder(a.root)
	fmt.Println()
	fmt.Println()
	fmt.Println("Height of the Tree Is --> ", a.Height(a.root))
	fmt.Println("Count of the Tree Is --> ", a.Count(a.root))
	a.Search(a.root,87)
}
/*
x.maketree(40,a,a)
y.maketree(60,a,a)
z.maketree(20,x,a)
r.maketree(50,a,a)
s.maketree(30,r,y)
t.maketree(10,z,s)
*/