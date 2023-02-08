package main 

import (
    "fmt"
)

type RBTNode struct {
    Val int
    Left, Right, P *RBTNode
    Color bool    //true 表示red， false 表示black
}

type RBTree struct {
    Root *RBTNode
    Nil *RBTNode
    Num int
}

func RBPrint(T *RBTree) {  //层次遍历红黑树
    root := T.Root
    var queue = [](*RBTNode){root}
    n_q := 1
    for len(queue) > 0 {
        now := queue[0]
        queue = queue[1:]
        n_q--
        if now != T.Nil {
            fmt.Printf(" d:%d c:%t p:%d", now.Val, now.Color, now.P.Val)
            if now.Left != T.Nil {
                queue = append(queue, now.Left)
            }
            if now.Right != T.Nil {
                queue = append(queue, now.Right)
            }
        }
        if n_q == 0 {
            n_q = len(queue)
            fmt.Println()
        }
    }
}

func RBLeftRotate(T *RBTree, z *RBTNode) {
    zr := z.Right
    z.Right = zr.Left 
    if zr.Left != T.Nil {
        zr.Left.P = z
    }
    zr.P = z.P
    if z == z.P.Left {
        z.P.Left = zr
    } else {
        z.P.Right = zr
    }
    zr.Left = z
    z.P = zr
    if z == T.Root {      //可能需要修改根结点
        T.Root = zr
    }
}

func RBRightRotate(T *RBTree, z *RBTNode) {
    zl := z.Left 
    z.Left = zl.Right 
    if zl.Right != T.Nil {
        zl.Right.P = z
    }
    zl.P = z.P
    if z == z.P.Left {
        z.P.Left = zl
    } else {
        z.P.Right = zl
    }
    zl.Right = z 
    z.P = zl
    if z == T.Root {      //可能需要修改根结点
        T.Root = zl
    }
}

func RBInsert(T *RBTree, z *RBTNode) {  
    T.Num++
    y := T.Nil          //用来设置z.P
    x := T.Root
    for x != T.Nil {    
        y = x
        if z.Val < x.Val {
            x = x.Left
        } else {
            x = x.Right
        }
    }
    z.P = y
    if T.Root == T.Nil {
        T.Root = z
    } else if z.Val < y.Val {  //这里不能使用x == y.Left判断，因为x为Nil时，y的左右子树都指向x
        y.Left = z
    } else {
        y.Right = z
    }
    z.Color = true
    z.Left = T.Nil
    z.Right = T.Nil
    RBInsertFixup(T, z)
}

func RBInsertFixup(T *RBTree, z *RBTNode) {  //z作为红色结点插入
    for z.P.Color == true {
        if z.P == z.P.P.Left{
            y := z.P.P.Right 
            if y.Color == true {            //case1
                y.Color = false
                z.P.Color = false
                z.P.P.Color = true
                z = z.P.P
            } else {
                if z == z.P.Right {         //case2
                    z = z.P
                    RBLeftRotate(T, z)
                }
                z.P.Color = false           //case3
                z.P.P.Color = true
                RBRightRotate(T, z.P.P)
            } 
        } else {
            y := z.P.P.Left 
            if y.Color == true {
                y.Color = false
                z.P.Color = false
                z.P.P.Color = true
                z = z.P.P
            } else {
                if z == z.P.Left {
                    z = z.P
                    RBRightRotate(T, z)
                }
                z.P.Color = false
                z.P.P.Color = true
                RBLeftRotate(T, z.P.P)
            }
        }
    }
    z = T.Root                               //最后找到根结点直接涂黑
    z.Color = false
}

func RBReplace(T *RBTree, u *RBTNode, v *RBTNode) {
    if u == T.Root {
        T.Root = v
    } else if u == u.P.Left {
        u.P.Left = v
    } else {
        u.P.Right = v
    }
    v.P = u.P
}

func RBDelete(T *RBTree, z *RBTNode) {
    T.Num--
    var x *RBTNode
    y := z                       //y总是指向被删除或者移动到内部的结点
    y_original_color := y.Color
    if z.Left == T.Nil {        //case1
        x = z.Right             //x指向y原来的位置
        RBReplace(T, z, z.Right)
    } else if z.Right == T.Nil {//case2
        x = z.Left 
        RBReplace(T, z, z.Left)
    } else { 
        y = z.Right 
        for y.Left != T.Nil {
            y = y.Left
        }
        y_original_color = y.Color 
        x = y.Right 
        if y.P == z {           //case3
            x.P = y
        } else {                //case4
            RBReplace(T, y, y.Right)
            y.Right = z.Right 
            y.Right.P = y
        }
        RBReplace(T, z, y)      
        y.Left = z.Left 
        y.Left.P = y 
        y.Color = z.Color
    }
    if y_original_color == false {
        RBDeleteFixup(T, x)
    }
}

func RBDeleteFixup(T *RBTree, x *RBTNode) {
    for x != T.Root && x.Color == false {
        if x == x.P.Left {
            w := x.P.Right 
            if w.Color == true {                                  //case1
                w.Color = false 
                x.P.Color = true
                RBLeftRotate(T, x.P)
                w = x.P.Right
            }
            if w.Left.Color == false && w.Right.Color == false {  //case2
                w.Color = true
                x = x.P
            } else if w.Right.Color == false {                    //case3
                w.Left.Color = false
                w.Color = true 
                RBRightRotate(T, w)
                w = x.P.Right
            }
            w.Color = x.P.Color                                   //case4
            x.P.Color = false
            w.Right.Color = false
            RBLeftRotate(T, x.P)
            x = T.Root
        } else {
            w := x.P.Left 
            if w.Color == true {                                  //case1
                w.Color = false
                x.P.Color = true
                RBRightRotate(T, x.P)
                w = x.P.Left
            } 
            if w.Left.Color == false && w.Right.Color == false {  //case2
                w.Color = true
                x = x.P
            } else if w.Left.Color == false {                     //case3
                w.Right.Color = false
                w.Color = true
                RBLeftRotate(T, w)
                w = x.P.Left
            }
            w.Color = x.P.Color                                   //case4
            x.P.Color = false
            w.Left.Color = false
            RBRightRotate(T, x.P)
            x = T.Root
        }
    }
    x.Color = false
}

func CreateRB(nums []int) (*RBTree, [](*RBTNode)){
    var queue [](*RBTNode)
    T := &RBTree{nil, &RBTNode{}, 0}
    T.Root = T.Nil
    for i:=0; i<len(nums); i++ {
        node := &RBTNode{nums[i], T.Nil, T.Nil, T.Nil, false}
        RBInsert(T, node)
        queue = append(queue, node)
    }
    return T, queue
}

func main() {
    T, queue := CreateRB([]int{1, 2, 3, 4, 5})
    RBPrint(T)
    for _, node := range queue {
        RBDelete(T, node)
        fmt.Println("----------------")
        RBPrint(T)
    }
}
