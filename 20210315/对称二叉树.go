/*
递归的重点 思考递归点
思考递归的退出条件
*/

func isSymmetric(root *TreeNode) bool {
	return check(root,root)
}

func check(left,right *TreeNode) bool{
	if left==nil&&right==nil{
		return true
	}
	if left==nil||right==nil{
		return false
	}
	return (left.Val==right.Val)&&check(left.Left,right.Right)&&check(left.Right,right.Left)
}
